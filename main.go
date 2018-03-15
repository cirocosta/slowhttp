package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	duration = flag.Duration("duration", 5*time.Second, "total time to respond")
	writes   = flag.Int("writes", 10, "number of writes to perform")
	port     = flag.Int("port", 8080, "port to listen to")

	writeInterval = *duration / time.Duration(*writes)
	bytesPerWrite = len(response) / *writes
)

func slowHandler(w http.ResponseWriter, r *http.Request) {
	var (
		start int = 0
		end   int = 0
		n     int
		err   error
	)

	for i := 0; i < *writes; i++ {
		start = bytesPerWrite * i
		end = bytesPerWrite * (i + 1)

		n, err = w.Write(response[start:end])
		if err != nil {
			log.Fatal(err)
		}

		if n != bytesPerWrite {
			log.Fatal(errors.New("wrote less than expected"))
		}

		w.(http.Flusher).Flush()

		time.Sleep(writeInterval)
	}

	totalWritten := bytesPerWrite * (*writes)

	if (len(response) - totalWritten) > 0 {
		start = totalWritten
		end = len(response)

		n, err = w.Write(response[start:end])
		if err != nil {
			log.Fatal(err)
		}

		if n != (end - start) {
			log.Fatal(errors.New("wrote less than expected"))
		}
	}

	return
}

func main() {
	http.HandleFunc("/", slowHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

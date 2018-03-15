install:
	go install -v

fmt:
	go fmt ./...

image:
	docker build -t cirocosta/slowhttp .

.PHONY: fmt build image

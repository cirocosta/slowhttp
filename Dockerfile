FROM golang:alpine as builder

ADD ./ /go/src/github.com/cirocosta/slowhttp

RUN set -ex && \
  cd /go/src/github.com/cirocosta/slowhttp && \
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./slowhttp /usr/bin/slowhttp

FROM busybox

COPY --from=builder /usr/bin/slowhttp /usr/local/bin/slowhttp

ENTRYPOINT [ "slowhttp" ]

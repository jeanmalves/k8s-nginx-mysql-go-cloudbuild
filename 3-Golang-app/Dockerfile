FROM golang:1.14

COPY src/ /go/src/

WORKDIR /go/src/github.com/jeanmalves/web/app

RUN go install
EXPOSE 8000
ENTRYPOINT ["/go/bin/app"]
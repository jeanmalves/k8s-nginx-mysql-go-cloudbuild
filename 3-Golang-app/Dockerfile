FROM golang:1.14

COPY src/ /go/src/

WORKDIR /go/src/github.com/jeanmalves/web/app

EXPOSE 8000
CMD ["go", "run", "/go/src/github.com/jeanmalves/web/app/main.go"]

# Unit test command:
# docker run  --name <name-container>  -v $(pwd):/go/src -w /go/src golang-webserver  go test github.com/jeanmalves/web/helper
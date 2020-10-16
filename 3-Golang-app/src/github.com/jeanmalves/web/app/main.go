package main

import (
    "fmt"
    "net/http"
    "github.com/jeanmalves/web/helper"
)

func handler(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/html; charset=utf-8")
    res.WriteHeader(http.StatusOK)
    fmt.Fprintf(res, "%s", helper.Greeting("Code.education Rocks!"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8000", nil)
}

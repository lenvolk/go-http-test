package main

import (
    "fmt"
    "time"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func gettime(w http.ResponseWriter, req *http.Request) {
    dt := time.Now()
    fmt.Fprintf(w, "Current date and time is: %v", dt.String())
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/time", gettime)

    http.ListenAndServe(":8090", nil)
}
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "[v2] Hello, Kubernetes!11111")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}

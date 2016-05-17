package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", http.Request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3030", nil)
}
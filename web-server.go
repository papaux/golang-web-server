package main

import (
	"io"
	"net/http"
	"fmt"
)

func main() {
	fmt.Printf("Starting web server...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello Fribourg!</h1>")
}

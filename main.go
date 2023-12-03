// very simple Golang server for testing
// to run: go run main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I'm here!"))
}

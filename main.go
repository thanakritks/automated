package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, DevSecOps with Golang!")
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

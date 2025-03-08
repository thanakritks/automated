package main

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, DevSecOps with Golang!")
}

func main() {
	http.HandleFunc("/", handleHomePage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, "error listening on port 8080:", err)
		os.Exit(1)
	}
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hello, DevSecOps with Golang!\n")
}

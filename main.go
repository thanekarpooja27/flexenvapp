package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			fmt.Fprintf(w, "middleware logged in")

	})

	http.ListenAndServe(":8080", nil)
}

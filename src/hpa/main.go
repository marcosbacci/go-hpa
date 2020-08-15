package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", greeting)
	http.ListenAndServe(":9000", nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Code.education Rocks!")
}

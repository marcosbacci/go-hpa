package main

import (
	"fmt"
	"net/http"
	"math"
)

func main() {
	http.HandleFunc("/", loop)
	http.ListenAndServe(":8000", nil)
}

func loop(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 0; i < 10; i++ {
		x += math.Sqrt(x)
	}

	fmt.Fprintf(w, "Code.education Rocks!")
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// RandomFloat returns a random Float
func RandomFloat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Random float: %f\n", rand.Float64())
}

// RandomInt returns a random Int
func RandomInt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Random int: %v\n", rand.Int())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randfloat", RandomFloat)
	mux.HandleFunc("/randint", RandomInt)
	http.ListenAndServe(":8000", mux)
}

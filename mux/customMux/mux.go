package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// CustomServeMux is our own mux
type CustomServeMux struct{}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Random number: %f\n", rand.Float64())
}

func main() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}

package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/*filepath", http.Dir("./"))

	log.Fatal(http.ListenAndServe(":8000", router))
}

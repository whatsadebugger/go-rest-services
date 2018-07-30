package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// ArticleHandler will return the category and id of an article
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category : %v\n", vars["category"])
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func no(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")

	url, err := rtr.Get("articleRoute").URL("category", "comedy", "id", "1")
	no(err)
	fmt.Println(url)

	server := &http.Server{
		Handler:      rtr,
		Addr:         "localhost:8000",
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
	}

	log.Fatal(server.ListenAndServe())
}

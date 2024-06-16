package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start() {
	router := chi.NewRouter()

	router.Get("/", home)

	log.Fatalln(http.ListenAndServe(":8080", router))
}

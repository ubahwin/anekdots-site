package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	anekLogic "github.com/ubahwin/anekdots-site/internal/handlers"
	"github.com/ubahwin/anekdots-site/internal/middleware"
	"log"
	"net/http"
)

func handleRequests() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.HeadersMiddleware)
	router.HandleFunc("/api/anekdot", anekLogic.GetAnekdot)

	http.Handle("/", router)

	return cors.AllowAll().Handler(router)
}

func main() {
	err := http.ListenAndServe(":9999", handleRequests())
	if err != nil {
		log.Fatal("Internal error!")
	}
}

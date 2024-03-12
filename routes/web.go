package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {
	r.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	r.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
}

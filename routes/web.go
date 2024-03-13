package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controlers"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {

	pc := new(controlers.PagesController)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
}

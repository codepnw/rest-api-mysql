package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	sr := r.PathPrefix("/api/v1").Subrouter()

	log.Println("starting the API server at", s.port)

	return sr
}

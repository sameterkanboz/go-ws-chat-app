package main

import (
	"github.com/bmizerany/pat"
	"go-ws/internal/handlers"
	"net/http"
)

func routesBUG() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}

package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lukaszczajkowski/golang-bookings/pkg/config"
	"github.com/lukaszczajkowski/golang-bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()

	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	// middleware to avoid panic and allow graceful error recovery
	mux.Use(middleware.Recoverer)
	// just an example of using a custom middleware
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// file server - a place to get the static content from
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

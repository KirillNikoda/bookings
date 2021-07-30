package main

import (
	config2 "github.com/KirillNikoda/bookings/pkg/config"
	handlers2 "github.com/KirillNikoda/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config2.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers2.Repo.Home)
	mux.Get("/about", handlers2.Repo.About)

	return mux
}

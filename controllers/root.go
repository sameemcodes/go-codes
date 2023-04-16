package controllers

import (
	"net/http"
	"go-codes/config"
	"go-codes/durable"
	"go-codes/controllers"
	"go-codes/controllers/routes"

	"github.com/gorilla/mux"
)


func RegisterRoutes(r *mux.Router, logger *durable.Logger) {
	r.HandleFunc("/health", routes.ServeHealthCheck)
	r.HandleFunc("/isAlive", routes.ServeLivelinessCheck)
}
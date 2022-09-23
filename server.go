package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/volume/flight-path-microservice/handlers"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/health", handlers.Healthcheck)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
		os.Exit(1)
	}
}

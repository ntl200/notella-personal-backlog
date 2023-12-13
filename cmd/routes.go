package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntl200/notella-personal-backlog/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	// listing backlog
	app.Get("/backlog", handlers.ListBacklogs)
	// create backlog
	app.Post("/backlog", handlers.CreateBacklog)
	// update backlog
	app.Put("/backlog/:id", handlers.UpdateBacklog)
	// delete backlog
	app.Delete("/backlog/:id", handlers.DeleteBacklog)
}

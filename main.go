package main

import (
	"github.com/shreyachandil/go-fiber-api/middleware"
	"github.com/shreyachandil/go-fiber-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Middleware to track analytics
	app.Use(middleware.Analytics())

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Go REST API 🚀")
	})

	// Post routes
	routes.SetupPostRoutes(app)

	app.Listen(":3000")
}

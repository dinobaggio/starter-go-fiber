package routes

import (
	"starter-go-fiber/libs"

	"github.com/gofiber/fiber/v2"
)

func SetUp(router *fiber.App) {
	api := router.Group("/api")
	v1 := api.Group("/v1")

	libs.Use(v1)

	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "It's Work!",
		})
	})
}

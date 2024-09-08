package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := fiber.New()

	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from /api/v1",
		})
	})

	port := ":8080"
	log.Info().Msgf("Starting server on port %s...", port)
	if err := app.Listen(port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	} else {
		log.Info().Msg("Server started successfully")
	}
}

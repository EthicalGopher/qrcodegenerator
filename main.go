package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func main() {
	app := fiber.New()

	// Serve the main HTML page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./Frontend/index.html")
	})

	// QR code generation endpoint
	app.Post("/generate", func(c *fiber.Ctx) error {
		text := c.FormValue("text")

		// Generate QR code
		png, err := qrcode.Encode(text, qrcode.Medium, 256)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate QR code")
		}

		// Send the generated QR code as an image
		c.Set("Content-Type", "image/png")
		return c.Send(png)
	})

	fmt.Println("Server is running at http://localhost:8000")
	app.Listen(":8000")
}

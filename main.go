package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func main() {
	app := fiber.New()

	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./Frontend/index.html")
	})


	app.Post("/generate", func(c *fiber.Ctx) error {
		text := c.FormValue("text")

		
		png, err := qrcode.Encode(text, qrcode.Medium, 256)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate QR code")
		}

		
		c.Set("Content-Type", "image/png")
		return c.Send(png)
	})

	fmt.Println("Server is running at http://localhost:8000")
	app.Listen(":8000")
}

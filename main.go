package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app:=fiber.New()
	app.Static("/", "./Frontend")
	fmt.Println("http://localhost:8000")
	app.Listen(":8000")

}
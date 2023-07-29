package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Initialize() {
	app := fiber.New()

	api := app.Group("v1")

	users := api.Group("/users")

	users.Get("list", handleGetUsers)
	users.Post("/save", handlePostUser)
	users.Patch("update", handleUpdateUser)
	users.Delete("delete", handleDeleteUser)





	log.Fatal(app.Listen(":3000"))
}
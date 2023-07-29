package router

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
	Email string
}

func handleGetUsers(context *fiber.Ctx) error {
	users := []User{
		{"Wesley", "sdasd@gmail.com"},
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
	}

	return context.JSON(users)
}

func handlePostUser(context *fiber.Ctx) error {
	return context.SendString("Save a user")
}

func handleUpdateUser(context *fiber.Ctx) error {
	return context.SendString("Update a user")
}

func handleDeleteUser(context *fiber.Ctx) error {
	return context.SendString("Delete a user")
}
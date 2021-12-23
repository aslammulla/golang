package main

import (
	"github.com/aslammulla/golang/rest_api/employee"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/employee", employee.GetEmployees)
	app.Get("api/employee/:id", employee.GetEmployee)
	app.Post("api/employee", employee.CreateEmployee)
	app.Delete("api/employee", employee.DeleteEmployee)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}

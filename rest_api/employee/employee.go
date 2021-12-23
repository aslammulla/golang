package employee

import (
	"github.com/gofiber/fiber"
)

func GetEmployees(e *fiber.Ctx) {
	e.Send("All Employee")
}

func GetEmployee(e *fiber.Ctx) {
	e.Send("A single employee")
}

func CreateEmployee(e *fiber.Ctx) {
	e.Send("Store Employee")
}

func DeleteEmployee(e *fiber.Ctx) {
	e.Send("Delete Employee")
}

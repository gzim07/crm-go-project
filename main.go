package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gzim07/go-fiber-crm-basic/database"
	"github.com/gzim07/go-fiber-crm-basic/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}
func init() {
	var err error

	database.DB, err = gorm.Open(sqlite.Open("lead.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection open to database")

}

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":8080")

}

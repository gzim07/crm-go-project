package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gzim07/go-fiber-crm-basic/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DB
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
	return nil
}

func NewLead(c *fiber.Ctx) error {
	db := database.DB
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(lead)
	c.JSON(lead)
	return nil
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No Lead found with Id"))
	}
	db.Delete(&lead)
	c.Send([]byte("Lead succesfullu deleted"))
	return nil
}

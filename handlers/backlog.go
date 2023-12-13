package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ntl200/notella-personal-backlog/database"
	"github.com/ntl200/notella-personal-backlog/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Notella Personal Backlog App!")
}

func ListBacklogs(c *fiber.Ctx) error {
	backlogs := []models.Backlog{}
	database.DB.Db.Find(&backlogs)

	return c.Status(200).JSON(backlogs)
}

func CreateBacklog(c *fiber.Ctx) error {
	backlog := new(models.Backlog)
	if err := c.BodyParser(backlog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&backlog)

	return c.Status(200).JSON(backlog)
}

func UpdateBacklog(c *fiber.Ctx) error {
	backlog := new(models.Backlog)
	if err := c.BodyParser(backlog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var result models.Backlog
	database.DB.Db.First(&result, c.Params("id"))

	backlog.ID = StringToUint(c.Params("id"))
	backlog.CreatedAt = result.CreatedAt
	database.DB.Db.Save(&backlog)

	return c.Status(200).JSON(backlog)
}

func DeleteBacklog(c *fiber.Ctx) error {
	backlog := new(models.Backlog)
	database.DB.Db.Unscoped().Where("id = ?", c.Params("id")).Delete(&backlog)

	return c.Status(200).JSON(fiber.Map{
		"id": c.Params("id"),
		"message": "Backlog successfully deleted!",
	})
}

func StringToUint(s string) uint {
    i, _ := strconv.Atoi(s)
    return uint(i)
}
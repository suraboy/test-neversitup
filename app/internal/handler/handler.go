package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/test-neversitup/app/internal/service"
)

type assignTestHandler struct {
	service service.AssignTestService
}

type PermutationModel struct {
	Input string `json:"input"`
}

type OddModel struct {
	Input []int `json:"input"`
}

type SmileyFaceModel struct {
	Input []string `json:"input"`
}

func (hdl assignTestHandler) GeneratePermutations(c *fiber.Ctx) error {
	var requestData PermutationModel
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	result := hdl.service.GeneratePermutations(requestData.Input)
	return c.JSON(result)
}

func (hdl assignTestHandler) FindOdd(c *fiber.Ctx) error {
	var requestData OddModel
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	result, err := hdl.service.FindOdd(requestData.Input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(result)
}

func (hdl assignTestHandler) CountSmileyFaces(c *fiber.Ctx) error {
	var requestData SmileyFaceModel
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	result := hdl.service.CheckCountSmileys(requestData.Input)

	return c.JSON(result)
}

func SetupProcess(service service.AssignTestService) AssignTestHandler {
	return &assignTestHandler{
		service: service,
	}
}

type AssignTestHandler interface {
	GeneratePermutations(c *fiber.Ctx) error
	FindOdd(c *fiber.Ctx) error
	CountSmileyFaces(c *fiber.Ctx) error
}

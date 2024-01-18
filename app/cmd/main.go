package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/suraboy/test-neversitup/app/internal/handler"
	"github.com/suraboy/test-neversitup/app/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	f := fiber.New(fiber.Config{
		DisableKeepalive: false,
	})

	f.Use(recovery.New())
	f.Use(cors.New())

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	svc := service.NewService()
	hdl := handler.SetupProcess(svc)

	f.Post("/assignments/permutations", hdl.GeneratePermutations)
	f.Post("/assignments/odd", hdl.FindOdd)
	f.Post("/assignments/count-smiley", hdl.CountSmileyFaces)

	go func() {
		address := fmt.Sprintf(":8081")
		err := f.Listen(address)
		if err != nil {
			panic("start server error")
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)
	// Graceful shutdown
	select {
	case <-ctx.Done():
		log.Info("terminating: context cancelled")
	case <-signalChan:
		log.Info("terminating: via signal")
	}
	cancel()
}

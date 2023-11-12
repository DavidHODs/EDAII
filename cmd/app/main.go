package main

import (
	"log"

	db "github.com/DavidHODs/EDAII/internal/database"
	utils "github.com/DavidHODs/EDAII/internal/utils"
	"github.com/DavidHODs/EDAII/pkg"
	"github.com/gofiber/fiber/v2"
)

const (
	natsLog = "internal/logs/nats.log"
	dbLog   = "internal/logs/db.log"
)

func main() {
	// initializes nats log file connection
	natsLogF, err := utils.Logger(natsLog)
	if err != nil {
		log.Fatalf("error: could not create nats ops log file: %s", err)
	}
	defer natsLogF.Close()

	// initializes db log file connection
	dbLogF, err := utils.Logger(dbLog)
	if err != nil {
		log.Fatalf("error: could not create db log file: %s", err)
	}
	defer dbLogF.Close()

	// initializes nats server connection
	nc := pkg.NatServerConn(natsLogF)

	// initializes database connection
	db := db.InitDB(dbLogF)

	// loads ConnectionManager struct with required services and file connections
	cm := &pkg.ConnectionManager{
		NC:      nc,
		DB:      db,
		NatsLog: natsLogF,
	}

	app := fiber.New()

	app.Post("/publish", cm.NatsOps)

	app.Listen(":3000")
}

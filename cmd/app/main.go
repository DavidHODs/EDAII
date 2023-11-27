package main

import (
	"fmt"
	"log"

	db "github.com/DavidHODs/EDAII/internal/database"
	utils "github.com/DavidHODs/EDAII/internal/utils"
	"github.com/DavidHODs/EDAII/pkg"
	"github.com/gofiber/fiber/v2"
)

const (
	natsLog       = "internal/logs/nats.log"
	dbLog         = "internal/logs/db.log"
	interruptsLog = "internal/logs/interrupts.log"
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

	// initializes interrupt log file connection
	interruptLogLogF, err := utils.Logger(interruptsLog)
	if err != nil {
		log.Fatalf("error: could not create db log file: %s", err)
	}
	defer interruptLogLogF.Close()

	// initializes nats server connection
	nc := pkg.NatServerConn(natsLogF)

	// initializes database connection
	db := db.InitDB(dbLogF)

	isInterruptLogEmpty, err := utils.IsFileEmpty(interruptsLog)
	if err != nil {
		log.Fatalf("could not determine if interrupt log file is empty: %v", err)
	}

	// loads ConnectionManager struct with required services and file connections
	cm := &pkg.ConnectionManager{
		NC:           nc,
		DB:           db,
		NatsLog:      natsLogF,
		InterruptLog: interruptLogLogF,
	}

	if !isInterruptLogEmpty {
		fmt.Println("file not empty")
		cm.NatsRecovery("trial")
	}

	app := fiber.New()

	app.Post("/publish", cm.NatsOps)

	app.Listen(":3000")
}

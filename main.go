package main

import (
	"github.com/RichardPMachado/Gopportunities/config"
	"github.com/RichardPMachado/Gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}

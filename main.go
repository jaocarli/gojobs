package main

import (
	"github.com/jaocarli/gojobs/config"
	"github.com/jaocarli/gojobs/router"
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

	// Initialize Router
	router.Initialize()
}

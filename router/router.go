package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Listening at :8080
	router.Run(":8080")
}

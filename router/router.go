package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Iniciliaza o Router utilizando as configurações Default do Gin
	router := gin.Default()
	// Definindo uma rota
	router.GET("./ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a API, por padrão na rota 0.0.0.0:8000
	router.Run(":8080")
}

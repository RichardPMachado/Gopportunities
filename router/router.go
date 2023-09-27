package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Defina uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		},
		)
	})
	// Estamos rodando a api
	router.Run(":8080") // listen and Serve on 0.0.0.0:8080
}

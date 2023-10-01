package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	router := gin.Default()
	// Initialize Routes
	initializeRoutes(router)
	// Run the Server
	router.Run(":8080") // listen and Serve on 0.0.0.0:8080
}

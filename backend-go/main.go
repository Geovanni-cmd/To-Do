package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ruta de prueba
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080") // Ejecuta el servidor en el puerto 8080
}

package main

import (
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Text string `json:"text"`
}

type ResponseBody struct {
	Output []string `json:"output"`
}

func main() {
	// Crea una instancia de Gin
	router := gin.Default()

	// Definir el manejador para la ruta /gemini/chat/
	router.POST("/gemini/chat/", handler)

	// Iniciar el servidor en el puerto 8080
	router.Run(":8081")
}

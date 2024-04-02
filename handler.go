package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nelbermora/go-gemini/gemini"
)

func handler(c *gin.Context) {
	apiKey := c.GetHeader("api-key")

	if apiKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invaid api key"})
		return
	}

	// Decodificar el JSON del cuerpo de la solicitud
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid field text"})
		return
	}

	output, err := gemini.AskForData(requestBody.Text, apiKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseBody := ResponseBody{
		Output: output,
	}

	c.JSON(http.StatusOK, responseBody)
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", HelloWorld)
	router.GET("/hi", func(c *gin.Context) {
		log.Print("Works")
	})
	port := "4000"
	router.Run("localhost:" + port)
}

type Response struct {
	message string
}

func HelloWorld(c *gin.Context) {
	resp := Response{message: "okay"}
	c.IndentedJSON(http.StatusOK, resp)
}

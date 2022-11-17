package main

import (
	"example/todo-go/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default() //new gin router initialization

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World from Patrik Duch :)"})
	}) // first endpoint returns Hello World

	router.Run() //running application, Default port is 8080
}

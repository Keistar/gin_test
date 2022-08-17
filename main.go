package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// Road HTML
	engine.LoadHTMLGlob("templates/*")

	// Routing
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"message": "hello gin",
		})
	})
	engine.Run(":3000")
}

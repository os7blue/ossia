package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	router.Static("/static", "web/static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

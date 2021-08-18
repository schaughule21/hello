package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func regRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("D:/Training/Goworkspace/src/Day7/HelloWorld//templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	sample := r.Group("/user")
	sample.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user-form.html", nil)
	})

	sample.POST("/add", func(c *gin.Context) {
		var user User

		c.Bind(&user)

		user.ID = 8
		users["8"] = user
		c.IndentedJSON(http.StatusOK, users)
	})

	r.Static("/public", "D:/Training/Goworkspace/src/Day7/HelloWorld/public")

	return r
}

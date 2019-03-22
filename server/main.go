package main

import (
	"net/http"

	// Gin packages
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../client", true)))

	// API Route Groups
	api := router.Group("/api/")
	{

		// Blog Subgroup
		blog := api.Group("/blog/")
		{
			blog.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "GET all blog posts here",
				})
			})
			blog.GET("/:id", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "GET blog post by id",
				})
			})
			blog.PUT("/:id", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "modify an existing blog post",
				})
			})
			blog.POST("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "POST to a new blogid",
					// if there is an exsisting id don't post send an error
					// could include some meta data
				})
			})
		}
		// Begin api base routes
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "base route reached",
			})
		})
		// more api routes go here
		// example:
		// api.GET("/specific/path/to/route/:id", someHandler)
	}
	// Start and run the server
	router.Run(":5000")
}

// // A standard handler function
// func someHandler(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")
// 	c.JSON(http.StatusOK, JSON_content)
// }

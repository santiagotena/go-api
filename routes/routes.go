package routes

import (
	"example/go-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.BookById)
	router.PATCH("/checkout", handlers.CheckoutBook)
	router.PATCH("/return", handlers.ReturnBook)
	router.POST("/books", handlers.CreateBook)
}

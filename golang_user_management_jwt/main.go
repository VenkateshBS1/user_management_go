package main

import (
	"log"
	"os"

	controller "github.com/venkateshBS1/golang_user_management_jwt/controllers"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	controller.NewUserController().AuthRoutes(router)
	controller.NewUserController().UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access granted for api-2"})
	})
	router.Run(":" + port)
}

package routes

import (
	controller "github.com/venkateshBS1/golang_user_management_jwt/controllers"
	"github.com/venkateshBS1/golang_user_management_jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

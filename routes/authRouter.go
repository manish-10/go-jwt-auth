package routes

import (
	controller "github.com/manish-10/go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/singup", controller.Signup())
	incomingRoutes.POST("/user/login", controller.Login())
}

package routes

import (
	controller "github.com/manish-10/go-jwt/controllers"
	middleware "github.com/manish-10/go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users/signup", controller.Signup())
}

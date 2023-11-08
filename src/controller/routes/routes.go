package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jorgebarcelos/first-go-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createtUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
package routes

import(
	controller "parkyee_backend/controllers"

	"girhub.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users" , controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}

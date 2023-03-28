package routes

import(
	controller "github.com/ArinCodes/parkyee_backend/controllers"
	"github.com/ArinCodes/parkyee_backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users" , controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}

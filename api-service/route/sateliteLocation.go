package route

import (
	"github.com/Abeldlp/fullinfo/api-service/controller"
	"github.com/gin-gonic/gin"
)

func AppendSateliteLocationRoute(r *gin.Engine) {
	satelite := r.Group("/")

	{
		satelite.GET("/", controller.GetAllSateliteLocations)
	}
}
package routes

import (
	"valorant-roulette/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.gohtml")

	r.Static("/agents", "wwwroot/agents/")
	r.Static("/css", "wwwroot/css/")
	r.Static("/js", "wwwroot/js/")
	r.Static("/lib", "wwwroot/lib/")
	r.Static("/maps", "wwwroot/maps/")
	r.Static("/weapons", "wwwroot/weapons/")
	r.Static("/favicon.ico", "wwwroot/favicon.ico/")

	r.GET("/", controllers.Index)
	r.GET("/agent", controllers.Agent)
	r.GET("/team", controllers.Team)
	r.GET("/map", controllers.Map)

	r.Run(":8080")
}

package routes

import (
	"github.com/jamesguru/dev-api/controllers"
	"github.com/labstack/echo/v4"
)

func Setup() {
	e := echo.New()
	e.POST("/ui", controllers.AddUI)
	e.GET("/ui/:id", controllers.GetUI)
	e.GET("/uis", controllers.GetUIs)
	e.Start(":3000")
}

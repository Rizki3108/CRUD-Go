package routes

import (
	"crud-api/controllers"

	"github.com/labstack/echo/v4"
)

func Routes() (e *echo.Echo) {
	e = echo.New()

	productRoutes := e.Group("/product")
	productRoutes.GET("", controllers.ReadAllProducts)
	productRoutes.POST("/create", controllers.CreateProducts)
	productRoutes.GET("/:id", controllers.ReadDetailProducts)
	productRoutes.DELETE("/:id", controllers.DeleteProduct)
	productRoutes.PUT("/update", controllers.UpdateProduct)

	categoryRoutes := e.Group("/category")
	categoryRoutes.GET("", controllers.ReadAllCategories)
	categoryRoutes.POST("/create", controllers.CreateCategories)
	categoryRoutes.GET("/:id", controllers.ReadDetailCategories)
	categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
	categoryRoutes.PUT("/update", controllers.UpdateCategory)

	return
}

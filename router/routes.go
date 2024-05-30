package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Guisegtrab/gooportunidades/handler"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})


		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})

		v1.DELETE("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Opening",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})


	}


	// Initialize Handler
	// handler.InitializeHandler()
	// basePath := "/api/v1"
	// docs.SwaggerInfo.BasePath = basePath
	// v1 := router.Group(basePath)
	// {
	// 	v1.GET("/opening", handler.ShowOpeningHandler)
	// 	v1.POST("/opening", handler.CreateOpeningHandler)
	// 	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	// 	v1.PUT("/opening", handler.UpdateOpeningHandler)
	// 	v1.GET("/openings", handler.ListOpeningsHandler)

	// }
	// // Initialize Swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
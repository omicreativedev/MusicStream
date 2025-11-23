package routes

import (
	controller "github.com/omicreativedev/MusicStream/Server/MusicStreamServer/controllers"
	"github.com/omicreativedev/MusicStream/Server/MusicStreamServer/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/music/:imdb_id", controller.GetMusic(client))
	router.POST("/addmusic", controller.AddMusic(client))
	router.GET("/recommendedmusics", controller.GetRecommendedMusics(client))
	router.PATCH("/updatereview/:music_id", controller.AdminReviewUpdate(client))
}

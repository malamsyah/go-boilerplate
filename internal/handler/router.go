package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/malamsyah/go-boilerplate/internal/middleware"
	"gorm.io/gorm"
)

func SetupRouter(_ *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.JSONLoggerMiddleware())
	r.GET("/health", Health)

	return r
}

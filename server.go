package main

import (
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const logpath = "./logs/go.log"

var logger *zap.Logger

func setupLog() {
	os.OpenFile(logpath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logpath}
	logger, _ = c.Build()
}

func main() {
	setupLog()

	r := gin.Default()
	// Settung gin as Zap Logger
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}

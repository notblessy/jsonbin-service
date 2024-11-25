package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/handler"
	"github.com/notblessy/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file")
	}
}

func main() {
	mysql := connectMysql()

	err := mysql.AutoMigrate(&model.PublicJSON{})
	if err != nil {
		logrus.Fatal("Failed to migrate database")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://jsonbin.sepiksel.com", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	h := handler.New(mysql)

	e.POST("/api", h.SaveJSON)
	e.GET("/api/:id", h.FindByID)

	e.Logger.Fatal(e.Start(":8080"))
}

func connectMysql() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to database")
	}

	return db
}

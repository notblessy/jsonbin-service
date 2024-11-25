package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/model"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *handler {
	return &handler{db}
}

func (h *handler) SaveJSON(c echo.Context) error {
	var jsonReq model.PublicJSON

	if err := c.Bind(&jsonReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	jsonReq.CreatedAt = time.Now()

	if err := h.db.Save(&jsonReq).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": jsonReq})
}

func (h *handler) FindByID(c echo.Context) error {
	id := c.Param("id")

	var json model.PublicJSON
	if err := h.db.Where("id = ?", id).First(&json).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": json})
}

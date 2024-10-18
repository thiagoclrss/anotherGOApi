package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagoclrss/anotherGOApi/model"
	"github.com/thiagoclrss/anotherGOApi/service"
	"net/http"
	"strconv"
)

type StockController struct {
	service *service.StockService
}

func NewStockController(service *service.StockService) *StockController {
	return &StockController{
		service: service,
	}
}

func (c *StockController) InitRoutes() {
	app := gin.Default()
	api := app.Group("/api/stock")

	api.GET("/:id", c.findByID)
}

func (c *StockController) findByID(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	stock, err := c.service.FindByID(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}
	context.JSON(http.StatusOK, stock)
}

func (c *StockController) saveStock(context *gin.Context) {
	stock := new(model.Stock)
	if err := context.ShouldBindJSON(&stock); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	id, err := c.service.SaveStock(*stock)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}
	context.JSON(
		http.StatusCreated,
		gin.H{"id": id},
	)
}

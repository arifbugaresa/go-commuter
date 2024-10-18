package station

import (
	"github.com/arifbugaresa/go-commuter/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Initiate(router *gin.RouterGroup) {
	var (
		stationService = NewService()
	)

	station := router.Group("/stations")
	{
		station.GET("", func(c *gin.Context) {
			GetAllStation(c, stationService)
		})
	}

	schedule := router.Group("/schedules")
	{
		schedule.GET("/:id", func(c *gin.Context) {
			CheckScheduleByStation(c, stationService)
		})
	}

}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.APIResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		common.APIResponse{
			Success: true,
			Message: "successfully get all station",
			Data:    datas,
		},
	)
}

func CheckScheduleByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckScheduleByStation(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.APIResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		common.APIResponse{
			Success: true,
			Message: "successfully get schedules",
			Data:    datas,
		},
	)
}

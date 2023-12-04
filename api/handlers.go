package api

import (
	"errors"
	"strconv"

	"github.com/BerryTracer/gps-data-service/service"

	"github.com/BerryTracer/gps-data-service/model"
)

type GPSHandler struct {
	Service service.GPSService
}

func NewGPSHandler(service service.GPSService) *GPSHandler {
	return &GPSHandler{Service: service}
}

func (g *GPSHandler) SaveGPSData(c HttpContext) error {
	var gpsData model.GPSData
	err := c.Bind(&gpsData)
	if err != nil {
		c.JSON(400, err)
		return err
	}

	return g.Service.Save(c.Context(), &gpsData)
}

func (g *GPSHandler) GetGPSDataByDeviceId(c HttpContext) error {
	deviceId := c.Params("device_id")
	if deviceId == "" {
		c.SendStatus(400)
		return errors.New("device_id is required")
	}

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}

	gpsDataArray, err := g.Service.FindByDeviceID(c.Context(), deviceId, limit, offset)
	if err != nil {
		c.JSON(400, err)
		return err
	}

	c.JSON(200, gpsDataArray)
	return nil
}

func (g *GPSHandler) GetGPSDataByUserId(c HttpContext) error {
	userId := c.Params("user_id")
	if userId == "" {
		c.SendStatus(400)
		return errors.New("user_id is required")
	}

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}

	gpsDataArray, err := g.Service.FindByUserID(c.Context(), userId, limit, offset)
	if err != nil {
		c.JSON(400, err)
		return err
	}

	c.JSON(200, gpsDataArray)
	return nil
}

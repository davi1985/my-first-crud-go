package controller

import (
	"github.com/davi1985/my-first-crud-go/src/config/rest_exceptions"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_exceptions.NewBadRequestException("Invalid route")

	c.JSON(err.Code, err)
}

package controller

import (
	"fmt"

	"github.com/davi1985/my-first-crud-go/src/config/exceptions"
	"github.com/davi1985/my-first-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		exception := exceptions.NewBadRequestException(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()),
		)

		c.JSON(exception.Code, exception)
		return
	}

	fmt.Println(userRequest)
}

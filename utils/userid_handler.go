package utils

import (
	// "net/http"
	"github.com/gin-gonic/gin"
)

func getUserIdFromContext(ctx *gin.Context) string {
	userIdInteface, exitsts := ctx.Get("user_id")
	if !exitsts {
		return "invalid user_id"
	}

	userId, ok := userIdInteface.(string)
	if !ok {
		return "User ID is not of type string"
	}

	return userId
}

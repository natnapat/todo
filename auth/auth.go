package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AccessToken(signature string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		})
		ss, err := token.SignedString([]byte(signature))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": ss,
		})
	}
}

package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/norrico31/rest-api-ecom/config"
)

func GenerateToken(userId int) (string, error) {
	exp := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(exp).Unix(),
	})
	secret := []byte(config.Envs.JWTSecret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

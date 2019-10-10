package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	KEY                    []byte = []byte("jwt-key")
	DEFAULT_EXPIRE_SECONDS int    = 600
)

// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	userName string
	jwt.StandardClaims
}

func GenerateToken(expiredSeconds int) (tokenString string) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	// Create the Claims
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	fmt.Println("token will be expired at ", time.Unix(expireAt, 0))
	// pass parameter to this func or not
	userName := "test"
	claims := MyCustomClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    userName,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate json web token failed !! error :", err)
	}
	return tokenStr

}

// 校验token是否有效
func ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.userName, claims.StandardClaims.ExpiresAt)
		fmt.Println("token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))
	} else {
		fmt.Println("validate tokenString failed !!!", err)
		return err
	}
	return nil
}

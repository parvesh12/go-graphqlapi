package api

import (
	"context"
	"fmt"

	"github.com/parvesh12/go-graphqlapi/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type API struct {
	DB *gorm.DB
}

func (api API) LoginHandler(ctx context.Context, user model.Login) (*bool, error) {

	var duser model.TblUser

	if err := api.DB.Model(model.TblUser{}).Where("user_name=?", user.Username).First(&duser).Error; err != nil {

		fmt.Printf("Error Getting user: %s\n", err)

		// resp := map[string]interface{}{
		// 	"status":  404,
		// 	"message": err.Error(),
		// }

		return false, err
	}

	perr := bcrypt.CompareHashAndPassword([]byte(duser.Password), []byte(user.Password))

	if perr != nil {

	}

	// resp := map[string]interface{}{
	// 	"status":  200,
	// 	"message": "logged in",
	// }

	return nil, nil
}

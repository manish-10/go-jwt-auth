package modules

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validation:"required, min = 2, max = 16"`
	LastName     *string            `json:"last_name " validation:"required, min = 2, max = 16"`
	Password     *string            `json:"password" validation:"required, min = 8, max = 16"`
	Phone        *string            `json:"phone" validation:"required"`
	Email        *string            `json:"email" validation:"required, email"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type" validation:"required, eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id"`
}

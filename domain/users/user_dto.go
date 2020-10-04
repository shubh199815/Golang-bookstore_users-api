package users

import (
	"strings"
	"github.com/shubh199815/bookstore_users-api/utils/errors"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

// func Validate(user *User) *errors.RestErr{
// 	user.Email = strings.Trimspace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}
// 	return nil
// }

// The above is a 'function' defined in the users package

// This is a 'method' defined for the User type so that a user can validate itself
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
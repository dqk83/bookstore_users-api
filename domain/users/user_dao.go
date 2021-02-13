package users

import (
	"fmt"

	"github.com/dqk83/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not fount", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		fmt.Errorf("user already exist %v\n", user.Id)
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %v already exists", current.Id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", current.Id))
	}
	fmt.Printf("creating user %v\n", user.Id)
	usersDB[user.Id] = user
	return nil
}

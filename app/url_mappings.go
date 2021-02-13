package app

import (
	"github.com/dqk83/bookstore_users-api/controllers/ping"
	"github.com/dqk83/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/userssearch", users.SearchUser)
}

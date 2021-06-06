package mysql_utils

import (
	"strings"

	"github.com/dqk83/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNoRows = "no rows"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated data")
	}
	return errors.NewInternalServerError("error processing request")
}

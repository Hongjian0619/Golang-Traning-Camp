package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func dao() error {
	return errors.Wrap(sql.ErrNoRows, "dao failed")
}

func biz() error {
	return errors.WithMessage(dao(), "biz failed")
}

func main() {
	err := biz()
	if err != nil {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}

	//do stuff
}

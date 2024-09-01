package main

import (
	"fmt"

	"github.com/malamsyah/go-boilerplate/internal/handler"
	"github.com/malamsyah/go-boilerplate/pkg/config"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting server...")

	// conf := config.Instance()

	var dbConn *gorm.DB
	var err error

	// fmt.Println("Connecting to database...")
	// dbConn, err = db.ConnectPostgres(conf)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Migrating database...")
	// err = db.Migrate(dbConn)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Setting up router...")
	r := handler.SetupRouter(dbConn)

	err = r.Run(":" + config.Instance().AppPort)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/firmafadi/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.connectB()
)

func main() {
	defer config.disconnectDB()

	// run all routes
	routes.Routes()
}

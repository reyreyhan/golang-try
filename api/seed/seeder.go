package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/reyreyhan/golang-try/api/models"
)

var users = []models.User{
	models.User{
		Username: "Reyhan Alphard",
		Email:    "rey@gmail.com",
		Password: "password",
		Address:  "Surabaya",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}

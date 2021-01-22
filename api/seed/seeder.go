package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/reyreyhan/golang-try/api/models"
)

var users = []models.User{
	models.User{
		Username: "reyhanalphard",
		Email:    "rey@gmail.com",
		Password: "password",
		Address:  "Surabaya",
	},

	models.User{
		Username: "alphardsavero",
		Email:    "alphard@gmail.com",
		Password: "password",
		Address:  "Sidoarjo",
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

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

}

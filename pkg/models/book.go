package models

import (
	"github.com/DhruvinShiroya/go-book-management-system/pkg/config"
	"github.com/DhruvinShiroya/go-toolrentalstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{

  gorm.Model
  Name string `json:"name"`
  Author string `json:"author"`
  Publication string `json:"Publication"`
}


func init(){
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}

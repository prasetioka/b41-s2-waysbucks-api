package database

import (
  "waysbucks-api/models"
  "waysbucks-api/pkg/mysql"
  "fmt"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(
    &models.User{},
    &models.Profile{},
    &models.Product{},
    &models.Toping{},
    &models.Transaction{},
  )

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}
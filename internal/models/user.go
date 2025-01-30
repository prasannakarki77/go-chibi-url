package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	Password     string `gorm:"null"`
	Provider     string `gorm:"not null; default:'email'"`
	ProviderID   string `gorm:"null"`
	RefreshToken string `gorm:"null"`
	LastLogin    time.Time
}

func (User) TableName() string {
	return "users"
}

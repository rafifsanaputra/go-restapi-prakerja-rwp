package models

type Guitar struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Brand       string `gorm:"type:varchar(300)" json:"brand"`
	Model       string `gorm:"type:varchar(300)" json:"model"`
	Description string `gorm:"type:text" json:"description"`
}

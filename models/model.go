package models

import "gorm.io/gorm"

type Backlog struct {
	gorm.Model
	Type	string 	`gorm:"type:varchar(64)" json:type`
	Title 	string 	`gorm:"type:varchar(300)" json:title`
	Creator	string 	`gorm:"type:varchar(300)" json:creator`
	Status	string 	`gorm:"type:varchar(64)" json:status`
}

package authtypes

type UserType struct {
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Firstname string `json:"firstname" gorm:"not null"`
	Lastname  string `json:"lastname" gorm:"not null"`
}

type GroupType struct {
	Name string `json:"name" gorm:"unique;not null"`
}

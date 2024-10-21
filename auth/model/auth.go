// Package users - User specific functions
package authmodel

import (
	"log"
	"net/http"

	authtypes "github.com/rbaylon/trixie_mods/auth/types"
	"github.com/rbaylon/trixie_mods/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	authtypes.UserType
}

// MigrateDB - Create the table if not exist in DB
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	var user User
	result := db.Where("Username = ?", "admin").First(&user)
	if result.Error != nil {
		log.Println("Starting app for the first time.")
		var (
			uname = database.GetEnvVariable("APP_ADMIN")
			upass = database.GetEnvVariable("APP_ADMIN_PW")
		)
		safepassword, _ := bcrypt.GenerateFromPassword([]byte(upass), 14)
		user.Username = uname
		user.Password = string(safepassword)
		user.Firstname = "Admin"
		user.Lastname = "istrator"
		res := db.Create(&user)
		if res == nil {
			log.Fatal("Failed to create admin user")
		}
	}
}

// Bind interface as required by go-chi/render
func (a *User) Bind(r *http.Request) error {
	return nil
}

type Crud interface {
	GetAll() ([]User, error)
	GetById(uid uint) (*User, error)
	Add(user *User) error
	Update(user *User) error
	Delete(user *User) error
	GetByUsername(username string) (*User, error)
}

type Storage struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) Add(user *User) error {
	safepassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(safepassword)
	result := s.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) GetAll() ([]User, error) {
	var users []User
	result := s.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *Storage) GetById(id uint) (*User, error) {
	var user User
	result := s.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *Storage) GetByUsername(username string) (*User, error) {
	var user User
	result := s.DB.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *Storage) Update(user *User) error {
	cost, _ := bcrypt.Cost([]byte(user.Password))
	if cost == 0 {
		safepassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			return err
		}
		user.Password = string(safepassword)
	}
	result := s.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) Delete(user *User) error {
	result := s.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

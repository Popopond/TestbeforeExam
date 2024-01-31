package entity
import (

)

type Member struct{
	ID uint
	Name string `valid:"required~Need Name, stringlength(2|10)"`
	PhoneNumber string `valid:"required~Need Phone, matches(^[0]\\d{9}$)~PhoneNumber is invalid"`
	Email string `valid:"required~Need Email, email~Email"`
}
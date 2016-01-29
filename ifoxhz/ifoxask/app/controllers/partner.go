package controllers

import (
//	"golang.org/x/crypto/bcrypt"
	//"fmt"
	"github.com/revel/revel"
	//"github.com/ifoxhz/ifoxask/app/models"
	//"github.com/ifoxhz/ifoxask/app/routes"
	//"strings"
)

type Partner struct {
		*revel.Controller
}

func (c Partner)  Showpangtuo()  revel.Result {
	return c.Render()
}

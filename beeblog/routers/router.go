package routers

import (
	"beeblog/controllers"
	"beeblog/models"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
}

package routers

import (
	"beeblog/controllers"
	"beeblog/models"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
	beego.Router("/", &controllers.MainController{})
}

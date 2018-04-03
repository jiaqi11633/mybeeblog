package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Ctx.WriteString("appname: " + beego.AppConfig.String("appname") +
		"\nhttpport: " + beego.AppConfig.String("httpport") +
		"\nrunmode: " + beego.AppConfig.String("runmode"))
	beego.Trace("trace test1")
	beego.Info("Info test1")
	beego.SetLevel(beego.LevelNotice)
	beego.Trace("trace test2")
	beego.Info("Info test2")

}

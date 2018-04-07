package main

import (
	_ "beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

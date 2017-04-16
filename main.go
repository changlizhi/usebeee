package main

import (
	"github.com/astaxie/beego"
	_ "usebeee/routers"
	"usebeee/utils"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
}

type DataJsonController struct {
	beego.Controller
}

func (this *DataJsonController) Get() {
	jsonstr := utils.ReadAll("static/data.json")
	this.Ctx.WriteString(jsonstr)
}

func main() {
	beego.Router("/main", &MainController{})
	beego.Router("/api/data", &DataJsonController{})
	beego.Run()
}

package routers

import (
	"PrometheusAlert/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//page
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	//template
	beego.Router("/template", &controllers.MainController{}, "get:Template")
	beego.Router("/template/add", &controllers.MainController{}, "get:TemplateAdd")
	beego.Router("/template/addtpl", &controllers.MainController{}, "post:AddTpl")
	beego.Router("/template/edit", &controllers.MainController{}, "get,post:TemplateEdit")
	beego.Router("/template/del", &controllers.MainController{}, "get:TemplateDel")
	beego.Router("/template/import", &controllers.MainController{}, "post:ImportTpl")
	//test
	beego.Router("/alerttest", &controllers.MainController{}, "post:AlertTest")
	beego.Router("/markdowntest", &controllers.MainController{}, "get,post:MarkdownTest")
	beego.Router("/test", &controllers.MainController{}, "get:Test")
	//record
	beego.Router("/record", &controllers.MainController{}, "get:Record")
	beego.Router("/record/clean", &controllers.MainController{}, "get:RecordClean")
	//alertrouter
	beego.Router("/alertrouter", &controllers.MainController{}, "get:AlertRouter")
	beego.Router("/alertrouter/add", &controllers.MainController{}, "get:RouterAdd")
	beego.Router("/alertrouter/edit", &controllers.MainController{}, "get:RouterEdit")
	beego.Router("/alertrouter/addrouter", &controllers.MainController{}, "post:AddRouter")
	beego.Router("/alertrouter/del", &controllers.MainController{}, "get:RouterDel")
	//setup
	beego.Router("/setup/weixin", &controllers.MainController{}, "get,post:SetupWeixin")

	// health
	beego.Router("/health", &controllers.MainController{}, "get:Health")

	//prometheus
	beego.Router("/prometheus/alert", &controllers.PrometheusController{}, "post:PrometheusAlert")
	beego.Router("/prometheus/router", &controllers.PrometheusController{}, "post:PrometheusRouter")

	beego.Router("/tengxun/status", &controllers.TengXunStatusController{}, "post:TengXunStatus")
	//zabbix
	beego.Router("/zabbix/alert", &controllers.ZabbixController{}, "post:ZabbixAlert")

	//webhook
	beego.Router("/prometheusalert", &controllers.PrometheusAlertController{}, "get,post:PrometheusAlert")

	// gitlab
	beego.Router("/gitlab/weixin", &controllers.GitlabController{}, "post:GitlabWeixin")
	beego.Router("/gitlab/dingding", &controllers.GitlabController{}, "post:GitlabDingding")
	beego.Router("/gitlab/feishu", &controllers.GitlabController{}, "post:GitlabFeishu")
	// Todo
	// Email
	// Feishu
}

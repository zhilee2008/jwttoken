// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"jwttoken/controllers"
	"jwttoken/utils"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"DELETE", "PUT", "POST", "GET", "OPTIONS", "*"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "Accept", "X-Requested-With"},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	// 	AllowCredentials: true,
	// }))

	beego.InsertFilter("/v1/object/*", beego.BeforeRouter, func(ctx *context.Context) {
		token := ctx.Input.Header("Authorization")
		if ctx.Request.RequestURI != "/login" && nil != utils.ValidateToken(token) {
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/v1/user/login", http.StatusMovedPermanently)
		}
	})
	utils.Init()
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

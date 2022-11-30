package controllers

import (
	"test_api/models"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// @Title TestEndpoint
// @Description Tests the API
// @Success 200 {object} models.Test
// @Failure 403 body is empty
// @router /hello [get]
func (t *TestController) Test() {
	Response := models.TestFunction()
	t.Data["json"] = Response
	t.ServeJSON()
}

// func (t *TestController) Get() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		user, err := models.GetUser(uid)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = user
// 		}
// 	}
// 	u.ServeJSON()
// }

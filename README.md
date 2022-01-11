> example

```
package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	valid "github.com/lshsuper/gin-validator/validator"
	"net/http"
)

type Request struct {
	valid.BaseResult
	Phone   string  `json:"phone" binding:"phone"`
	UserID  int     `json:"userId" binding:"gt=0"`
	Content string  `json:"content" binding:"must"`
	Age     int     `json:"age" binding:"gt=20"`
	Money   float64 `json:"money" binding:"gt=35.6"`
	Type  int       `json:"type" binding:"in=66/33/3"`

}

//Check 重写check,做一些定制操作
func (r *Request) Check(validTag, tag, param string) error {

	//这里做重写，可以得到定制化的提示信息
	if tag == "money" {
		return errors.New("这数字不对路子")
	}

	return r.BaseResult.Check(validTag, tag, param)

}

func main() {

	r := gin.Default()

	//全局注册验证器
	valid.Register()

	r.POST("/post", func(context *gin.Context) {
		var req Request
		if err := valid.Form(context).CheckJSON(&req); err != nil {
			context.JSON(http.StatusOK, err.Error())
			return
		}

		context.JSON(http.StatusOK, "ok")

	})

	r.Run(":10086")

}

```



> example

```
package main

import (
	"github.com/gin-gonic/gin"
	"go-test/valid"
	"net/http"
)

type Request struct {
	Phone   string  `json:"phone" binding:"phone"`
	UserID  int     `json:"userId" binding:"gt=0"`
	Content string  `json:"content" binding:"must"`
	Age     int     `json:"age" binding:"gt=20"`
	Money   float64 `json:"money" binding:"gt=35.6"`
}

func (r Request) Check() error {
	return nil
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



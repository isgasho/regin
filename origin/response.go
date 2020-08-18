package origin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-touch/regin/base"
)

type ResponseHandler struct {
	page404 string
	page500 string
}

// 定义ResponseHandler
var Response *ResponseHandler

func init() {
	Response = &ResponseHandler{}
}

// 输出
func (rh *ResponseHandler) Output(c *gin.Context, result *base.Result) error {
	// 响应
	switch result.Type {
	case "String":
		return rh.OutputString(c, result)
	case "Json":
		return rh.OutputJson(c, result)
	case "Html":
		return rh.OutputHtml(c, result)
	case "Stream":
		return rh.OutputStream(c, result)
	}
	return nil
}

// 输出json
func (rh *ResponseHandler) OutputJson(c *gin.Context, result *base.Result) error {
	// 响应
	switch result.Status {
	case 200:
		if result.GetData("code") == 0 {
			result.SetData("msg", "请求成功")
		}
	case 404:
		result.SetData("code", 404)
		result.SetData("msg", result.Msg)
	case 500:
		result.SetData("code", 500)
		result.SetData("msg", result.Msg)
	default:
		result.SetData("code", result.Status)
		result.SetData("msg", result.Msg)
	}

	// 输出
	c.JSON(result.Status, result.Data)
	return nil
}

// 输出String
func (rh *ResponseHandler) OutputString(c *gin.Context, result *base.Result) error {
	// 响应
	switch result.Status {
	case 404:
		result.SetData("code", 404)
		result.SetData("msg", result.Msg)
	case 500:
		result.SetData("code", 500)
		result.SetData("msg", result.Msg)
	default:
		result.SetData("msg", "请求成功")
	}

	// 输出
	c.JSON(result.Status, result.Data)
	return nil
}

// 输出Html
func (rh *ResponseHandler) OutputHtml(c *gin.Context, result *base.Result) error {
	// 响应
	switch result.Status {
	case 404:
		result.SetData("code", 404)
		result.SetData("msg", result.Msg)
	case 500:
		result.SetData("code", 500)
		result.SetData("msg", result.Msg)
	default:
		result.SetData("msg", "请求成功")
	}

	// 输出
	c.JSON(result.Status, result.Data)
	return nil
}

// 输出Html
func (rh *ResponseHandler) OutputStream(c *gin.Context, result *base.Result) error {
	// 响应
	switch result.Status {
	case 404:
		result.SetData("code", 404)
		result.SetData("msg", result.Msg)
	case 500:
		result.SetData("code", 500)
		result.SetData("msg", result.Msg)
	default:
		result.SetData("msg", "请求成功")
	}

	// 头信息
	if header := result.GetHeader(); len(header) > 0 {
		for key, value := range header {
			c.Header(key, value)
		}
	}

	// 用户自定义输出
	if userFunc := result.GetOutput(); userFunc != nil {
		userFunc(c.Writer)
	} else { // 默认输出
		c.JSON(result.Status, result.Data)
	}
	return nil
}

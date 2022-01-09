package request

import (
	"net/http"
	"time"
)


// Request  基础请求参数
type Http struct {
	BaseUri     string            // 基础uri
	BaseHeaders map[string]string //基础请求头
}

// Response 响应
type Response struct {
	BodyStr string //body字符串
	*http.Response
}


//http请求客户端
var httpClient = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Timeout: time.Second*30,
}

// NewRequestClient 创建请求客户端
func NewRequestClient()  {
	
}

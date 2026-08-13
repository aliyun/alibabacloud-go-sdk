// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendChatMessageResponseBody
	GetCode() *string
	SetData(v interface{}) *SendChatMessageResponseBody
	GetData() interface{}
	SetMessage(v string) *SendChatMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendChatMessageResponseBody
	GetRequestId() *string
}

type SendChatMessageResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应数据负载
	//
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendChatMessageResponseBody) GetData() interface{} {
	return s.Data
}

func (s *SendChatMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendChatMessageResponseBody) SetCode(v string) *SendChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatMessageResponseBody) SetData(v interface{}) *SendChatMessageResponseBody {
	s.Data = v
	return s
}

func (s *SendChatMessageResponseBody) SetMessage(v string) *SendChatMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendChatMessageResponseBody) SetRequestId(v string) *SendChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

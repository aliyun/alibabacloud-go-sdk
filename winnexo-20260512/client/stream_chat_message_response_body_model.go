// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StreamChatMessageResponseBody
	GetCode() *string
	SetData(v interface{}) *StreamChatMessageResponseBody
	GetData() interface{}
	SetMessage(v string) *StreamChatMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *StreamChatMessageResponseBody
	GetRequestId() *string
}

type StreamChatMessageResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// SSE 事件流负载；成功时响应为 text/event-stream 原始帧，需按流式方式逐帧消费
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

func (s StreamChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StreamChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *StreamChatMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *StreamChatMessageResponseBody) GetData() interface{} {
	return s.Data
}

func (s *StreamChatMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StreamChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StreamChatMessageResponseBody) SetCode(v string) *StreamChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *StreamChatMessageResponseBody) SetData(v interface{}) *StreamChatMessageResponseBody {
	s.Data = v
	return s
}

func (s *StreamChatMessageResponseBody) SetMessage(v string) *StreamChatMessageResponseBody {
	s.Message = &v
	return s
}

func (s *StreamChatMessageResponseBody) SetRequestId(v string) *StreamChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *StreamChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

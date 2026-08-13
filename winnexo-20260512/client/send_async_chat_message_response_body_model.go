// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendAsyncChatMessageResponseBody
	GetCode() *string
	SetMessage(v string) *SendAsyncChatMessageResponseBody
	GetMessage() *string
	SetMessageId(v string) *SendAsyncChatMessageResponseBody
	GetMessageId() *string
	SetRequestId(v string) *SendAsyncChatMessageResponseBody
	GetRequestId() *string
	SetSessionCreated(v bool) *SendAsyncChatMessageResponseBody
	GetSessionCreated() *bool
	SetSessionId(v string) *SendAsyncChatMessageResponseBody
	GetSessionId() *string
}

type SendAsyncChatMessageResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 助手消息ID；用于随后调用 streamChatMessage 订阅生成结果
	//
	// example:
	//
	// 3cf84d92-f273-4bb7-ab3c-52646d25ec30
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 本次调用是否新建了会话
	//
	// example:
	//
	// true
	SessionCreated *bool `json:"sessionCreated,omitempty" xml:"sessionCreated,omitempty"`
	// 会话ID；续写会话时与入参一致，新建会话时为服务端生成值
	//
	// example:
	//
	// bd772dcc-afab-44ad-9fb8-bca716726201
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s SendAsyncChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendAsyncChatMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendAsyncChatMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SendAsyncChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendAsyncChatMessageResponseBody) GetSessionCreated() *bool {
	return s.SessionCreated
}

func (s *SendAsyncChatMessageResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *SendAsyncChatMessageResponseBody) SetCode(v string) *SendAsyncChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) SetMessage(v string) *SendAsyncChatMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) SetMessageId(v string) *SendAsyncChatMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) SetRequestId(v string) *SendAsyncChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) SetSessionCreated(v bool) *SendAsyncChatMessageResponseBody {
	s.SessionCreated = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) SetSessionId(v string) *SendAsyncChatMessageResponseBody {
	s.SessionId = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendWorkItemMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendWorkItemMessageResponseBody
	GetCode() *string
	SetMessage(v string) *SendWorkItemMessageResponseBody
	GetMessage() *string
	SetMessageId(v string) *SendWorkItemMessageResponseBody
	GetMessageId() *string
	SetRequestId(v string) *SendWorkItemMessageResponseBody
	GetRequestId() *string
	SetSessionId(v string) *SendWorkItemMessageResponseBody
	GetSessionId() *string
	SetUserMessageId(v string) *SendWorkItemMessageResponseBody
	GetUserMessageId() *string
	SetWorkItemId(v string) *SendWorkItemMessageResponseBody
	GetWorkItemId() *string
}

type SendWorkItemMessageResponseBody struct {
	// 公共业务状态码；成功为200，失败为后端错误码。
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 公共状态说明；成功时为success，失败时描述具体原因。
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 本轮助手消息ID；回复在后台异步生成，取得该ID不表示AI已开始或已完成回复。
	//
	// example:
	//
	// 23456789-abcd-4ef0-8123-456789abcdef
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// 请求追踪ID，可用于定位本次调用及排查失败原因。
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 本次消息发送到的已有话题会话ID，不是新建会话ID。
	//
	// example:
	//
	// 01234567-89ab-4def-8123-456789abcdef
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 本次提交的用户消息ID。
	//
	// example:
	//
	// 12345678-9abc-4def-8123-456789abcdef
	UserMessageId *string `json:"userMessageId,omitempty" xml:"userMessageId,omitempty"`
	// 请求传入的工作项ID；仅传conversationId时返回legacy-topic:<conversationId>形式的兼容ID。
	//
	// example:
	//
	// wi_0123456789abcdef0123456789abcdef
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s SendWorkItemMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendWorkItemMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendWorkItemMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendWorkItemMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendWorkItemMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SendWorkItemMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendWorkItemMessageResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *SendWorkItemMessageResponseBody) GetUserMessageId() *string {
	return s.UserMessageId
}

func (s *SendWorkItemMessageResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *SendWorkItemMessageResponseBody) SetCode(v string) *SendWorkItemMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) SetMessage(v string) *SendWorkItemMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) SetMessageId(v string) *SendWorkItemMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) SetRequestId(v string) *SendWorkItemMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) SetSessionId(v string) *SendWorkItemMessageResponseBody {
	s.SessionId = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) SetUserMessageId(v string) *SendWorkItemMessageResponseBody {
	s.UserMessageId = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) SetWorkItemId(v string) *SendWorkItemMessageResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *SendWorkItemMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

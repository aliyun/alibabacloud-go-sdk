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
	SetContent(v string) *SendChatMessageResponseBody
	GetContent() *string
	SetData(v interface{}) *SendChatMessageResponseBody
	GetData() interface{}
	SetMessage(v string) *SendChatMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendChatMessageResponseBody
	GetRequestId() *string
	SetType(v string) *SendChatMessageResponseBody
	GetType() *string
	SetWorkMode(v string) *SendChatMessageResponseBody
	GetWorkMode() *string
}

type SendChatMessageResponseBody struct {
	// Deprecated
	//
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The streaming response content.
	//
	// example:
	//
	// yes，i\\"m ready
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The Server-Sent Events (SSE) event stream payload. If the request is successful, the response is the raw frame of text/event-stream, which must be consumed frame by frame in a streaming manner.
	//
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Deprecated
	//
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Deprecated
	//
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The event type.
	//
	// example:
	//
	// think
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The session work mode. Valid values: ask (quick Q&A, which trims tools, skills, and connectors for single-turn direct answers), work (deep work, which is the default value), and direct (direct connection mode at the request level, which does not start the sandbox and prevents context pollution, equivalent to directChat=true). The ask and work modes are at the session level. The mode is fixed when a new session is created. Subsequent chats inherit the session mode by default. If the explicitly passed value is inconsistent with the session mode, a parameter error is returned. To switch the mode, create a new session or fork the session. If ask is passed in the multi-digital employee or task execution scenario, it takes effect as work. If directChat is set to true, this parameter is ignored.
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
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

func (s *SendChatMessageResponseBody) GetContent() *string {
	return s.Content
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

func (s *SendChatMessageResponseBody) GetType() *string {
	return s.Type
}

func (s *SendChatMessageResponseBody) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendChatMessageResponseBody) SetCode(v string) *SendChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatMessageResponseBody) SetContent(v string) *SendChatMessageResponseBody {
	s.Content = &v
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

func (s *SendChatMessageResponseBody) SetType(v string) *SendChatMessageResponseBody {
	s.Type = &v
	return s
}

func (s *SendChatMessageResponseBody) SetWorkMode(v string) *SendChatMessageResponseBody {
	s.WorkMode = &v
	return s
}

func (s *SendChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

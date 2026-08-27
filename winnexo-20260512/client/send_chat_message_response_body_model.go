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
	// The SSE event stream payload. On success, the response is a text/event-stream raw frame that must be consumed frame by frame in streaming mode.
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

func (s *SendChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

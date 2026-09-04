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
	SetContent(v string) *StreamChatMessageResponseBody
	GetContent() *string
	SetData(v interface{}) *StreamChatMessageResponseBody
	GetData() interface{}
	SetMessage(v string) *StreamChatMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *StreamChatMessageResponseBody
	GetRequestId() *string
	SetType(v string) *StreamChatMessageResponseBody
	GetType() *string
}

type StreamChatMessageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The incremental content of the current SSE frame.
	//
	// example:
	//
	// Hello
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The SSE event stream payload. On success, the response is returned as raw text/event-stream frames that must be consumed frame by frame in streaming mode.
	//
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The SSE event type, such as text, think, heartbeat, done, or error.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *StreamChatMessageResponseBody) GetContent() *string {
	return s.Content
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

func (s *StreamChatMessageResponseBody) GetType() *string {
	return s.Type
}

func (s *StreamChatMessageResponseBody) SetCode(v string) *StreamChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *StreamChatMessageResponseBody) SetContent(v string) *StreamChatMessageResponseBody {
	s.Content = &v
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

func (s *StreamChatMessageResponseBody) SetType(v string) *StreamChatMessageResponseBody {
	s.Type = &v
	return s
}

func (s *StreamChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

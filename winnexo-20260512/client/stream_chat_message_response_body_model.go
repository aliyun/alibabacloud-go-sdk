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
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The SSE event stream payload. On success, the response is in text/event-stream raw frames and must be consumed frame by frame in a streaming manner.
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

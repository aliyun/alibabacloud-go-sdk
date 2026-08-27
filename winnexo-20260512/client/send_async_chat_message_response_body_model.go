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
	SetUserMessageId(v string) *SendAsyncChatMessageResponseBody
	GetUserMessageId() *string
}

type SendAsyncChatMessageResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- or InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This is empty when the request succeeds.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The assistant message ID. Use this ID to call streamChatMessage to subscribe to the generation results.
	//
	// example:
	//
	// 3cf84d92-f273-4bb7-ab3c-52646d25ec30
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether a new session was created by this call.
	//
	// example:
	//
	// true
	SessionCreated *bool `json:"sessionCreated,omitempty" xml:"sessionCreated,omitempty"`
	// The session ID. For continued sessions, this matches the input value. For new sessions, this is a server-generated value.
	//
	// example:
	//
	// bd772dcc-afab-44ad-9fb8-bca716726201
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user message ID. Use this ID to establish a pairing relationship with the assistant message in this turn.
	//
	// example:
	//
	// 60756cc6-8c53-4d1f-8db8-b8c09b81a5cb
	UserMessageId *string `json:"userMessageId,omitempty" xml:"userMessageId,omitempty"`
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

func (s *SendAsyncChatMessageResponseBody) GetUserMessageId() *string {
	return s.UserMessageId
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

func (s *SendAsyncChatMessageResponseBody) SetUserMessageId(v string) *SendAsyncChatMessageResponseBody {
	s.UserMessageId = &v
	return s
}

func (s *SendAsyncChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

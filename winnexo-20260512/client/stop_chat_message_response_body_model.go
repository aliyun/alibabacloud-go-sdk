// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopChatMessageResponseBody
	GetCode() *string
	SetFinishReason(v string) *StopChatMessageResponseBody
	GetFinishReason() *string
	SetMessage(v string) *StopChatMessageResponseBody
	GetMessage() *string
	SetMessageId(v string) *StopChatMessageResponseBody
	GetMessageId() *string
	SetPartialContent(v string) *StopChatMessageResponseBody
	GetPartialContent() *string
	SetRequestId(v string) *StopChatMessageResponseBody
	GetRequestId() *string
	SetSessionId(v string) *StopChatMessageResponseBody
	GetSessionId() *string
	SetStatus(v string) *StopChatMessageResponseBody
	GetStatus() *string
}

type StopChatMessageResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The reason for stopping.
	//
	// example:
	//
	// string_value
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The message ID.
	//
	// example:
	//
	// exampleMessageId
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The partially generated content.
	//
	// example:
	//
	// string_value
	PartialContent *string `json:"partialContent,omitempty" xml:"partialContent,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The final status of the message.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StopChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *StopChatMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopChatMessageResponseBody) GetFinishReason() *string {
	return s.FinishReason
}

func (s *StopChatMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopChatMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *StopChatMessageResponseBody) GetPartialContent() *string {
	return s.PartialContent
}

func (s *StopChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopChatMessageResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *StopChatMessageResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StopChatMessageResponseBody) SetCode(v string) *StopChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *StopChatMessageResponseBody) SetFinishReason(v string) *StopChatMessageResponseBody {
	s.FinishReason = &v
	return s
}

func (s *StopChatMessageResponseBody) SetMessage(v string) *StopChatMessageResponseBody {
	s.Message = &v
	return s
}

func (s *StopChatMessageResponseBody) SetMessageId(v string) *StopChatMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *StopChatMessageResponseBody) SetPartialContent(v string) *StopChatMessageResponseBody {
	s.PartialContent = &v
	return s
}

func (s *StopChatMessageResponseBody) SetRequestId(v string) *StopChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopChatMessageResponseBody) SetSessionId(v string) *StopChatMessageResponseBody {
	s.SessionId = &v
	return s
}

func (s *StopChatMessageResponseBody) SetStatus(v string) *StopChatMessageResponseBody {
	s.Status = &v
	return s
}

func (s *StopChatMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

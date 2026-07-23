// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *ChatConversationResponseBody
	GetAnswer() *string
	SetCallId(v string) *ChatConversationResponseBody
	GetCallId() *string
	SetConversationId(v string) *ChatConversationResponseBody
	GetConversationId() *string
	SetErrorCode(v string) *ChatConversationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ChatConversationResponseBody
	GetErrorMessage() *string
	SetEvent(v string) *ChatConversationResponseBody
	GetEvent() *string
	SetFunc(v string) *ChatConversationResponseBody
	GetFunc() *string
	SetGmtCreateTime(v string) *ChatConversationResponseBody
	GetGmtCreateTime() *string
	SetMessageId(v string) *ChatConversationResponseBody
	GetMessageId() *string
	SetRequestId(v string) *ChatConversationResponseBody
	GetRequestId() *string
	SetSkillName(v string) *ChatConversationResponseBody
	GetSkillName() *string
}

type ChatConversationResponseBody struct {
	// The reply content.
	//
	// example:
	//
	// FeatureDB近7天读写量为xxx
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The call ID.
	//
	// example:
	//
	// ""
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// e47cfae9-c0cc-42e1-91e2-e67cdb0e7b96
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SERVER_ERROR
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// connection failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The event.
	//
	// example:
	//
	// message
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The function name.
	//
	// example:
	//
	// ""
	Func *string `json:"Func,omitempty" xml:"Func,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The message ID.
	//
	// example:
	//
	// chat-abcdefg
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill name.
	//
	// example:
	//
	// ""
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s ChatConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ChatConversationResponseBody) GetAnswer() *string {
	return s.Answer
}

func (s *ChatConversationResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *ChatConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *ChatConversationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChatConversationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChatConversationResponseBody) GetEvent() *string {
	return s.Event
}

func (s *ChatConversationResponseBody) GetFunc() *string {
	return s.Func
}

func (s *ChatConversationResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ChatConversationResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ChatConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatConversationResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *ChatConversationResponseBody) SetAnswer(v string) *ChatConversationResponseBody {
	s.Answer = &v
	return s
}

func (s *ChatConversationResponseBody) SetCallId(v string) *ChatConversationResponseBody {
	s.CallId = &v
	return s
}

func (s *ChatConversationResponseBody) SetConversationId(v string) *ChatConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *ChatConversationResponseBody) SetErrorCode(v string) *ChatConversationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChatConversationResponseBody) SetErrorMessage(v string) *ChatConversationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChatConversationResponseBody) SetEvent(v string) *ChatConversationResponseBody {
	s.Event = &v
	return s
}

func (s *ChatConversationResponseBody) SetFunc(v string) *ChatConversationResponseBody {
	s.Func = &v
	return s
}

func (s *ChatConversationResponseBody) SetGmtCreateTime(v string) *ChatConversationResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *ChatConversationResponseBody) SetMessageId(v string) *ChatConversationResponseBody {
	s.MessageId = &v
	return s
}

func (s *ChatConversationResponseBody) SetRequestId(v string) *ChatConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatConversationResponseBody) SetSkillName(v string) *ChatConversationResponseBody {
	s.SkillName = &v
	return s
}

func (s *ChatConversationResponseBody) Validate() error {
	return dara.Validate(s)
}

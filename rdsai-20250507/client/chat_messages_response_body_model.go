// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *ChatMessagesResponseBody
	GetAnswer() *string
	SetConversationId(v string) *ChatMessagesResponseBody
	GetConversationId() *string
	SetCreatedAt(v int64) *ChatMessagesResponseBody
	GetCreatedAt() *int64
	SetEvent(v string) *ChatMessagesResponseBody
	GetEvent() *string
	SetId(v string) *ChatMessagesResponseBody
	GetId() *string
	SetMessageId(v string) *ChatMessagesResponseBody
	GetMessageId() *string
	SetMode(v string) *ChatMessagesResponseBody
	GetMode() *string
	SetRequestId(v string) *ChatMessagesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ChatMessagesResponseBody
	GetTaskId() *string
}

type ChatMessagesResponseBody struct {
	// The answer.
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The ID of the conversation.
	//
	// example:
	//
	// 9cbbe885-b240-4803-9d15-6781a3fd****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The creation time of the conversation.
	//
	// example:
	//
	// 1763986004
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The event.
	//
	// example:
	//
	// MysqlIOException
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 60b335ca-124d-4ee1-864b-de554987****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The message ID.
	//
	// example:
	//
	// oas8pwy2-slxw-sf98-bx83-cb2hkktl****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The query mode.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 01c3d43d-9466-4bd5-8196-4cbbce08****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatMessagesResponseBody) GetAnswer() *string {
	return s.Answer
}

func (s *ChatMessagesResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *ChatMessagesResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ChatMessagesResponseBody) GetEvent() *string {
	return s.Event
}

func (s *ChatMessagesResponseBody) GetId() *string {
	return s.Id
}

func (s *ChatMessagesResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ChatMessagesResponseBody) GetMode() *string {
	return s.Mode
}

func (s *ChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatMessagesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ChatMessagesResponseBody) SetAnswer(v string) *ChatMessagesResponseBody {
	s.Answer = &v
	return s
}

func (s *ChatMessagesResponseBody) SetConversationId(v string) *ChatMessagesResponseBody {
	s.ConversationId = &v
	return s
}

func (s *ChatMessagesResponseBody) SetCreatedAt(v int64) *ChatMessagesResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ChatMessagesResponseBody) SetEvent(v string) *ChatMessagesResponseBody {
	s.Event = &v
	return s
}

func (s *ChatMessagesResponseBody) SetId(v string) *ChatMessagesResponseBody {
	s.Id = &v
	return s
}

func (s *ChatMessagesResponseBody) SetMessageId(v string) *ChatMessagesResponseBody {
	s.MessageId = &v
	return s
}

func (s *ChatMessagesResponseBody) SetMode(v string) *ChatMessagesResponseBody {
	s.Mode = &v
	return s
}

func (s *ChatMessagesResponseBody) SetRequestId(v string) *ChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatMessagesResponseBody) SetTaskId(v string) *ChatMessagesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ChatMessagesResponseBody) Validate() error {
	return dara.Validate(s)
}

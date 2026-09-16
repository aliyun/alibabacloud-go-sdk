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
	SetApprovalStatus(v string) *ChatMessagesResponseBody
	GetApprovalStatus() *string
	SetCallId(v string) *ChatMessagesResponseBody
	GetCallId() *string
	SetConversationId(v string) *ChatMessagesResponseBody
	GetConversationId() *string
	SetCreatedAt(v int64) *ChatMessagesResponseBody
	GetCreatedAt() *int64
	SetDescription(v string) *ChatMessagesResponseBody
	GetDescription() *string
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
	SetRoundId(v string) *ChatMessagesResponseBody
	GetRoundId() *string
	SetTaskId(v string) *ChatMessagesResponseBody
	GetTaskId() *string
	SetToolArguments(v map[string]interface{}) *ChatMessagesResponseBody
	GetToolArguments() map[string]interface{}
	SetToolName(v string) *ChatMessagesResponseBody
	GetToolName() *string
}

type ChatMessagesResponseBody struct {
	// The answer content.
	//
	// example:
	//
	// The disk usage of instance rm-bp14as9914vd3***	- is 23%, and storage expansion is not needed at this time. If you need to view the detailed configuration, performance monitoring, or perform other operations for an instance, let me know your specific requirements!
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The tool invocation approval status.
	//
	// example:
	//
	// pending
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The tool invocation ID.
	//
	// example:
	//
	// call-example
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The conversation ID.
	//
	// example:
	//
	// 9cbbe885-b240-4803-9d15-6781a3fd****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1763986004
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The tool invocation description.
	//
	// example:
	//
	// Search ContextDB records
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	//
	// example:
	//
	// This field will be deprecated in the future. Ignore it
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tool approval round ID.
	//
	// example:
	//
	// round-example
	RoundId *string `json:"RoundId,omitempty" xml:"RoundId,omitempty"`
	// The asynchronous task ID.
	//
	// example:
	//
	// 01c3d43d-9466-4bd5-8196-4cbbce08****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The tool invocation parameters.
	ToolArguments map[string]interface{} `json:"ToolArguments,omitempty" xml:"ToolArguments,omitempty"`
	// The tool name.
	//
	// example:
	//
	// contextdb.search
	ToolName *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
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

func (s *ChatMessagesResponseBody) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *ChatMessagesResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *ChatMessagesResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *ChatMessagesResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ChatMessagesResponseBody) GetDescription() *string {
	return s.Description
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

func (s *ChatMessagesResponseBody) GetRoundId() *string {
	return s.RoundId
}

func (s *ChatMessagesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ChatMessagesResponseBody) GetToolArguments() map[string]interface{} {
	return s.ToolArguments
}

func (s *ChatMessagesResponseBody) GetToolName() *string {
	return s.ToolName
}

func (s *ChatMessagesResponseBody) SetAnswer(v string) *ChatMessagesResponseBody {
	s.Answer = &v
	return s
}

func (s *ChatMessagesResponseBody) SetApprovalStatus(v string) *ChatMessagesResponseBody {
	s.ApprovalStatus = &v
	return s
}

func (s *ChatMessagesResponseBody) SetCallId(v string) *ChatMessagesResponseBody {
	s.CallId = &v
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

func (s *ChatMessagesResponseBody) SetDescription(v string) *ChatMessagesResponseBody {
	s.Description = &v
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

func (s *ChatMessagesResponseBody) SetRoundId(v string) *ChatMessagesResponseBody {
	s.RoundId = &v
	return s
}

func (s *ChatMessagesResponseBody) SetTaskId(v string) *ChatMessagesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ChatMessagesResponseBody) SetToolArguments(v map[string]interface{}) *ChatMessagesResponseBody {
	s.ToolArguments = v
	return s
}

func (s *ChatMessagesResponseBody) SetToolName(v string) *ChatMessagesResponseBody {
	s.ToolName = &v
	return s
}

func (s *ChatMessagesResponseBody) Validate() error {
	return dara.Validate(s)
}

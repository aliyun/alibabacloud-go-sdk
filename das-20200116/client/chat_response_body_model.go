// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivityType(v string) *ChatResponseBody
	GetActivityType() *string
	SetContent(v string) *ChatResponseBody
	GetContent() *string
	SetDelta(v string) *ChatResponseBody
	GetDelta() *string
	SetMessageId(v string) *ChatResponseBody
	GetMessageId() *string
	SetName(v string) *ChatResponseBody
	GetName() *string
	SetParentMessageId(v string) *ChatResponseBody
	GetParentMessageId() *string
	SetRole(v string) *ChatResponseBody
	GetRole() *string
	SetRunId(v string) *ChatResponseBody
	GetRunId() *string
	SetStepName(v string) *ChatResponseBody
	GetStepName() *string
	SetTaskTrackerId(v string) *ChatResponseBody
	GetTaskTrackerId() *string
	SetThreadId(v string) *ChatResponseBody
	GetThreadId() *string
	SetToolCallId(v string) *ChatResponseBody
	GetToolCallId() *string
	SetToolCallName(v string) *ChatResponseBody
	GetToolCallName() *string
	SetType(v string) *ChatResponseBody
	GetType() *string
	SetValue(v interface{}) *ChatResponseBody
	GetValue() interface{}
}

type ChatResponseBody struct {
	// example:
	//
	// waiting_for_agent_thinking
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// example:
	//
	// I see you have several PolarDB instances, and I will query them for you shortly
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// hello
	Delta *string `json:"Delta,omitempty" xml:"Delta,omitempty"`
	// example:
	//
	// 61820b594664275c4429****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// summary
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 76bee207-31ee-4707-8851-6b9d4da033aa
	ParentMessageId *string `json:"ParentMessageId,omitempty" xml:"ParentMessageId,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// ed7cb7b1-ddc8-45d7-9ff3-b315726cb5f7
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// sub_agent_performance_diagnose_mysql
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// example:
	//
	// das_api
	TaskTrackerId *string `json:"TaskTrackerId,omitempty" xml:"TaskTrackerId,omitempty"`
	// example:
	//
	// 8e481be1-21d5-4a92-a2fb-fb54be0ab4f6
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// example:
	//
	// call_edf9cdb69e0e4c9796a6a5a6
	ToolCallId *string `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
	// example:
	//
	// das_api
	ToolCallName *string `json:"ToolCallName,omitempty" xml:"ToolCallName,omitempty"`
	// example:
	//
	// TEXT_MESSAGE_CONTENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// {"CharCount":393,"End":1777428785996,"RequestId":"BE59AED5-D831-5811-BBAD-590B917B2089","SessionId":"123e4567-e89b-12d3-a456-xxxxxxxxxxxx","Start":1777428707927}
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetActivityType() *string {
	return s.ActivityType
}

func (s *ChatResponseBody) GetContent() *string {
	return s.Content
}

func (s *ChatResponseBody) GetDelta() *string {
	return s.Delta
}

func (s *ChatResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ChatResponseBody) GetName() *string {
	return s.Name
}

func (s *ChatResponseBody) GetParentMessageId() *string {
	return s.ParentMessageId
}

func (s *ChatResponseBody) GetRole() *string {
	return s.Role
}

func (s *ChatResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *ChatResponseBody) GetStepName() *string {
	return s.StepName
}

func (s *ChatResponseBody) GetTaskTrackerId() *string {
	return s.TaskTrackerId
}

func (s *ChatResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *ChatResponseBody) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *ChatResponseBody) GetToolCallName() *string {
	return s.ToolCallName
}

func (s *ChatResponseBody) GetType() *string {
	return s.Type
}

func (s *ChatResponseBody) GetValue() interface{} {
	return s.Value
}

func (s *ChatResponseBody) SetActivityType(v string) *ChatResponseBody {
	s.ActivityType = &v
	return s
}

func (s *ChatResponseBody) SetContent(v string) *ChatResponseBody {
	s.Content = &v
	return s
}

func (s *ChatResponseBody) SetDelta(v string) *ChatResponseBody {
	s.Delta = &v
	return s
}

func (s *ChatResponseBody) SetMessageId(v string) *ChatResponseBody {
	s.MessageId = &v
	return s
}

func (s *ChatResponseBody) SetName(v string) *ChatResponseBody {
	s.Name = &v
	return s
}

func (s *ChatResponseBody) SetParentMessageId(v string) *ChatResponseBody {
	s.ParentMessageId = &v
	return s
}

func (s *ChatResponseBody) SetRole(v string) *ChatResponseBody {
	s.Role = &v
	return s
}

func (s *ChatResponseBody) SetRunId(v string) *ChatResponseBody {
	s.RunId = &v
	return s
}

func (s *ChatResponseBody) SetStepName(v string) *ChatResponseBody {
	s.StepName = &v
	return s
}

func (s *ChatResponseBody) SetTaskTrackerId(v string) *ChatResponseBody {
	s.TaskTrackerId = &v
	return s
}

func (s *ChatResponseBody) SetThreadId(v string) *ChatResponseBody {
	s.ThreadId = &v
	return s
}

func (s *ChatResponseBody) SetToolCallId(v string) *ChatResponseBody {
	s.ToolCallId = &v
	return s
}

func (s *ChatResponseBody) SetToolCallName(v string) *ChatResponseBody {
	s.ToolCallName = &v
	return s
}

func (s *ChatResponseBody) SetType(v string) *ChatResponseBody {
	s.Type = &v
	return s
}

func (s *ChatResponseBody) SetValue(v interface{}) *ChatResponseBody {
	s.Value = v
	return s
}

func (s *ChatResponseBody) Validate() error {
	return dara.Validate(s)
}

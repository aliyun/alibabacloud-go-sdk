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
	SetAgentId(v string) *ChatResponseBody
	GetAgentId() *string
	SetContent(v string) *ChatResponseBody
	GetContent() *string
	SetDelta(v string) *ChatResponseBody
	GetDelta() *string
	SetKind(v string) *ChatResponseBody
	GetKind() *string
	SetLabel(v string) *ChatResponseBody
	GetLabel() *string
	SetMessageId(v string) *ChatResponseBody
	GetMessageId() *string
	SetName(v string) *ChatResponseBody
	GetName() *string
	SetOriginatingToolCallId(v string) *ChatResponseBody
	GetOriginatingToolCallId() *string
	SetParentAgentId(v string) *ChatResponseBody
	GetParentAgentId() *string
	SetParentMessageId(v string) *ChatResponseBody
	GetParentMessageId() *string
	SetRole(v string) *ChatResponseBody
	GetRole() *string
	SetRunId(v string) *ChatResponseBody
	GetRunId() *string
	SetStepName(v string) *ChatResponseBody
	GetStepName() *string
	SetStepStatus(v string) *ChatResponseBody
	GetStepStatus() *string
	SetTaskTrackerId(v string) *ChatResponseBody
	GetTaskTrackerId() *string
	SetThreadId(v string) *ChatResponseBody
	GetThreadId() *string
	SetTimestamp(v int64) *ChatResponseBody
	GetTimestamp() *int64
	SetToolCallError(v string) *ChatResponseBody
	GetToolCallError() *string
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
	// The heartbeat.
	//
	// example:
	//
	// waiting_for_agent_thinking
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	AgentId      *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The response content.
	//
	// example:
	//
	// I see you have several PolarDB instances, and I will query them for you shortly
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the content is incremental.
	//
	// example:
	//
	// hello
	Delta *string `json:"Delta,omitempty" xml:"Delta,omitempty"`
	Kind  *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 61820b594664275c4429****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The extension key.
	//
	// example:
	//
	// summary
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginatingToolCallId *string `json:"OriginatingToolCallId,omitempty" xml:"OriginatingToolCallId,omitempty"`
	ParentAgentId         *string `json:"ParentAgentId,omitempty" xml:"ParentAgentId,omitempty"`
	// The parent message ID.
	//
	// example:
	//
	// 76bee207-31ee-4707-8851-6b9d4da033aa
	ParentMessageId *string `json:"ParentMessageId,omitempty" xml:"ParentMessageId,omitempty"`
	// The conversation role ID.
	//
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The run ID.
	//
	// example:
	//
	// ed7cb7b1-ddc8-45d7-9ff3-b315726cb5f7
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The execution step name.
	//
	// example:
	//
	// sub_agent_performance_diagnose_mysql
	StepName   *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepStatus *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
	// The callback tool class.
	//
	// example:
	//
	// das_api
	TaskTrackerId *string `json:"TaskTrackerId,omitempty" xml:"TaskTrackerId,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// 8e481be1-21d5-4a92-a2fb-fb54be0ab4f6
	ThreadId      *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ToolCallError *string `json:"ToolCallError,omitempty" xml:"ToolCallError,omitempty"`
	// The tool calling invoke ID.
	//
	// example:
	//
	// call_edf9cdb69e0e4c9796a6a5a6
	ToolCallId *string `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
	// The tool name.
	//
	// example:
	//
	// das_api
	ToolCallName *string `json:"ToolCallName,omitempty" xml:"ToolCallName,omitempty"`
	// The event type.
	//
	// example:
	//
	// TEXT_MESSAGE_CONTENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The extension value.
	//
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

func (s *ChatResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *ChatResponseBody) GetContent() *string {
	return s.Content
}

func (s *ChatResponseBody) GetDelta() *string {
	return s.Delta
}

func (s *ChatResponseBody) GetKind() *string {
	return s.Kind
}

func (s *ChatResponseBody) GetLabel() *string {
	return s.Label
}

func (s *ChatResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ChatResponseBody) GetName() *string {
	return s.Name
}

func (s *ChatResponseBody) GetOriginatingToolCallId() *string {
	return s.OriginatingToolCallId
}

func (s *ChatResponseBody) GetParentAgentId() *string {
	return s.ParentAgentId
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

func (s *ChatResponseBody) GetStepStatus() *string {
	return s.StepStatus
}

func (s *ChatResponseBody) GetTaskTrackerId() *string {
	return s.TaskTrackerId
}

func (s *ChatResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *ChatResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ChatResponseBody) GetToolCallError() *string {
	return s.ToolCallError
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

func (s *ChatResponseBody) SetAgentId(v string) *ChatResponseBody {
	s.AgentId = &v
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

func (s *ChatResponseBody) SetKind(v string) *ChatResponseBody {
	s.Kind = &v
	return s
}

func (s *ChatResponseBody) SetLabel(v string) *ChatResponseBody {
	s.Label = &v
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

func (s *ChatResponseBody) SetOriginatingToolCallId(v string) *ChatResponseBody {
	s.OriginatingToolCallId = &v
	return s
}

func (s *ChatResponseBody) SetParentAgentId(v string) *ChatResponseBody {
	s.ParentAgentId = &v
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

func (s *ChatResponseBody) SetStepStatus(v string) *ChatResponseBody {
	s.StepStatus = &v
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

func (s *ChatResponseBody) SetTimestamp(v int64) *ChatResponseBody {
	s.Timestamp = &v
	return s
}

func (s *ChatResponseBody) SetToolCallError(v string) *ChatResponseBody {
	s.ToolCallError = &v
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

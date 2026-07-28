// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatAiAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompleted(v bool) *ChatAiAgentResponseBody
	GetCompleted() *bool
	SetDisplayName(v string) *ChatAiAgentResponseBody
	GetDisplayName() *string
	SetErrorMessage(v string) *ChatAiAgentResponseBody
	GetErrorMessage() *string
	SetErrorType(v string) *ChatAiAgentResponseBody
	GetErrorType() *string
	SetEvent(v string) *ChatAiAgentResponseBody
	GetEvent() *string
	SetInput(v interface{}) *ChatAiAgentResponseBody
	GetInput() interface{}
	SetItems(v []*ChatAiAgentResponseBodyItems) *ChatAiAgentResponseBody
	GetItems() []*ChatAiAgentResponseBodyItems
	SetMessage(v string) *ChatAiAgentResponseBody
	GetMessage() *string
	SetSessionId(v string) *ChatAiAgentResponseBody
	GetSessionId() *string
	SetSuccess(v string) *ChatAiAgentResponseBody
	GetSuccess() *string
	SetText(v string) *ChatAiAgentResponseBody
	GetText() *string
	SetToolCallId(v string) *ChatAiAgentResponseBody
	GetToolCallId() *string
	SetToolName(v string) *ChatAiAgentResponseBody
	GetToolName() *string
	SetUsage(v *ChatAiAgentResponseBodyUsage) *ChatAiAgentResponseBody
	GetUsage() *ChatAiAgentResponseBodyUsage
}

type ChatAiAgentResponseBody struct {
	// Indicates whether this text segment is complete (the last segment of the message it belongs to).
	//
	// example:
	//
	// true
	Completed *bool `json:"completed,omitempty" xml:"completed,omitempty"`
	// The localized display name of the tool.
	//
	// example:
	//
	// ""
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The error message when the tool call fails (only when success is false).
	//
	// example:
	//
	// TIMEOUT
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The error type when the tool call fails (only when success is false).
	//
	// example:
	//
	// TIMEOUT
	ErrorType *string `json:"errorType,omitempty" xml:"errorType,omitempty"`
	// The event type.
	//
	// example:
	//
	// -
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// The tool input key-value pairs. The structure varies depending on the toolName.
	//
	// example:
	//
	// { "namespace": "vvp-dev-team", "jobId": "aa91ec66-...", "deploymentId": "b78aae4c-..." }
	Input interface{} `json:"input,omitempty" xml:"input,omitempty"`
	// The list of items pending approval.
	Items []*ChatAiAgentResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error message (for error events).
	//
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The session ID for this conversation.
	//
	// example:
	//
	// 462E2707-590E-51B6-9940-0AB33044828B-deliverData-202603020950-WCSN4MEC8T
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Indicates whether the tool calling invoke is successful.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
	// The text output from the assistant.
	//
	// example:
	//
	// -
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The tool calling ID, used to pair the invoke call and result.
	//
	// example:
	//
	// "c1"
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
	// The tool function name.
	//
	// example:
	//
	// "get_job_events"
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
	// The token usage.
	Usage *ChatAiAgentResponseBodyUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s ChatAiAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ChatAiAgentResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *ChatAiAgentResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ChatAiAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChatAiAgentResponseBody) GetErrorType() *string {
	return s.ErrorType
}

func (s *ChatAiAgentResponseBody) GetEvent() *string {
	return s.Event
}

func (s *ChatAiAgentResponseBody) GetInput() interface{} {
	return s.Input
}

func (s *ChatAiAgentResponseBody) GetItems() []*ChatAiAgentResponseBodyItems {
	return s.Items
}

func (s *ChatAiAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatAiAgentResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatAiAgentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ChatAiAgentResponseBody) GetText() *string {
	return s.Text
}

func (s *ChatAiAgentResponseBody) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *ChatAiAgentResponseBody) GetToolName() *string {
	return s.ToolName
}

func (s *ChatAiAgentResponseBody) GetUsage() *ChatAiAgentResponseBodyUsage {
	return s.Usage
}

func (s *ChatAiAgentResponseBody) SetCompleted(v bool) *ChatAiAgentResponseBody {
	s.Completed = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetDisplayName(v string) *ChatAiAgentResponseBody {
	s.DisplayName = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetErrorMessage(v string) *ChatAiAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetErrorType(v string) *ChatAiAgentResponseBody {
	s.ErrorType = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetEvent(v string) *ChatAiAgentResponseBody {
	s.Event = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetInput(v interface{}) *ChatAiAgentResponseBody {
	s.Input = v
	return s
}

func (s *ChatAiAgentResponseBody) SetItems(v []*ChatAiAgentResponseBodyItems) *ChatAiAgentResponseBody {
	s.Items = v
	return s
}

func (s *ChatAiAgentResponseBody) SetMessage(v string) *ChatAiAgentResponseBody {
	s.Message = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetSessionId(v string) *ChatAiAgentResponseBody {
	s.SessionId = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetSuccess(v string) *ChatAiAgentResponseBody {
	s.Success = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetText(v string) *ChatAiAgentResponseBody {
	s.Text = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetToolCallId(v string) *ChatAiAgentResponseBody {
	s.ToolCallId = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetToolName(v string) *ChatAiAgentResponseBody {
	s.ToolName = &v
	return s
}

func (s *ChatAiAgentResponseBody) SetUsage(v *ChatAiAgentResponseBodyUsage) *ChatAiAgentResponseBody {
	s.Usage = v
	return s
}

func (s *ChatAiAgentResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatAiAgentResponseBodyItems struct {
	// The original tool parameter key-value pairs.
	//
	// example:
	//
	// []
	Args interface{} `json:"args,omitempty" xml:"args,omitempty"`
	// The display name of the tool.
	//
	// example:
	//
	// “”
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The approval item ID, used when returning hitlDecisions.
	//
	// example:
	//
	// ""
	HitlId *string `json:"hitlId,omitempty" xml:"hitlId,omitempty"`
	// The name of the intercepted tool.
	//
	// example:
	//
	// “”
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s ChatAiAgentResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ChatAiAgentResponseBodyItems) GetArgs() interface{} {
	return s.Args
}

func (s *ChatAiAgentResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ChatAiAgentResponseBodyItems) GetHitlId() *string {
	return s.HitlId
}

func (s *ChatAiAgentResponseBodyItems) GetToolName() *string {
	return s.ToolName
}

func (s *ChatAiAgentResponseBodyItems) SetArgs(v interface{}) *ChatAiAgentResponseBodyItems {
	s.Args = v
	return s
}

func (s *ChatAiAgentResponseBodyItems) SetDisplayName(v string) *ChatAiAgentResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ChatAiAgentResponseBodyItems) SetHitlId(v string) *ChatAiAgentResponseBodyItems {
	s.HitlId = &v
	return s
}

func (s *ChatAiAgentResponseBodyItems) SetToolName(v string) *ChatAiAgentResponseBodyItems {
	s.ToolName = &v
	return s
}

func (s *ChatAiAgentResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type ChatAiAgentResponseBodyUsage struct {
	// The number of input tokens.
	//
	// example:
	//
	// 10
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 5
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The total number of tokens.
	//
	// example:
	//
	// 15
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s ChatAiAgentResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *ChatAiAgentResponseBodyUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *ChatAiAgentResponseBodyUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *ChatAiAgentResponseBodyUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ChatAiAgentResponseBodyUsage) SetInputTokens(v int64) *ChatAiAgentResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *ChatAiAgentResponseBodyUsage) SetOutputTokens(v int64) *ChatAiAgentResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *ChatAiAgentResponseBodyUsage) SetTotalTokens(v int64) *ChatAiAgentResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

func (s *ChatAiAgentResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}

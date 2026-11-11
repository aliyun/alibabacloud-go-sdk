// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeAgentResponseBody
	GetCode() *string
	SetContent(v string) *InvokeAgentResponseBody
	GetContent() *string
	SetDelta(v string) *InvokeAgentResponseBody
	GetDelta() *string
	SetMessage(v string) *InvokeAgentResponseBody
	GetMessage() *string
	SetMessageId(v string) *InvokeAgentResponseBody
	GetMessageId() *string
	SetName(v string) *InvokeAgentResponseBody
	GetName() *string
	SetOutcome(v *InvokeAgentResponseBodyOutcome) *InvokeAgentResponseBody
	GetOutcome() *InvokeAgentResponseBodyOutcome
	SetParentMessageId(v string) *InvokeAgentResponseBody
	GetParentMessageId() *string
	SetRole(v string) *InvokeAgentResponseBody
	GetRole() *string
	SetRunId(v string) *InvokeAgentResponseBody
	GetRunId() *string
	SetSnapshot(v interface{}) *InvokeAgentResponseBody
	GetSnapshot() interface{}
	SetThreadId(v string) *InvokeAgentResponseBody
	GetThreadId() *string
	SetToolCallId(v string) *InvokeAgentResponseBody
	GetToolCallId() *string
	SetToolCallName(v string) *InvokeAgentResponseBody
	GetToolCallName() *string
	SetType(v string) *InvokeAgentResponseBody
	GetType() *string
	SetValue(v interface{}) *InvokeAgentResponseBody
	GetValue() interface{}
}

type InvokeAgentResponseBody struct {
	// The error code for the RUN_ERROR event.
	//
	// example:
	//
	// InternalError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The tool output text for the TOOL_CALL_RESULT event. The text may be truncated.
	//
	// example:
	//
	// done
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The incremental fragment of text, reasoning, or tool parameters. Fragments with the same ID are concatenated in order.
	//
	// example:
	//
	// Hello
	Delta *string `json:"delta,omitempty" xml:"delta,omitempty"`
	// The error description for the RUN_ERROR event.
	//
	// example:
	//
	// Request failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the message to which the text, reasoning, or tool result belongs.
	//
	// example:
	//
	// m-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The name of the CUSTOM extension event, such as analysis_result.
	//
	// example:
	//
	// analysis_result
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The completion reason for the RUN_FINISHED event. A value of success indicates normal completion, and interrupt indicates waiting for a response.
	//
	// example:
	//
	// {"type":"success"}
	Outcome *InvokeAgentResponseBodyOutcome `json:"outcome,omitempty" xml:"outcome,omitempty" type:"Struct"`
	// The optional parent message ID for the TOOL_CALL_START event.
	//
	// example:
	//
	// m-1
	ParentMessageId *string `json:"parentMessageId,omitempty" xml:"parentMessageId,omitempty"`
	// The role of the message. The role is assistant for text messages and tool for tool results.
	//
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The run ID of the current run for RUN_STARTED and RUN_FINISHED events.
	//
	// example:
	//
	// r-1
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The thread snapshot for the STATE_SNAPSHOT event, including sessionId and requestId.
	//
	// example:
	//
	// {"sessionId":"abcdef0123456789abcdef0123456789","requestId":"req-1"}
	Snapshot interface{} `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
	// The thread ID for RUN_STARTED and RUN_FINISHED events. It is consistent with the request and the thread snapshot.
	//
	// example:
	//
	// abcdef0123456789abcdef0123456789
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// The tool invoke ID. It associates TOOL_CALL_START/ARGS/END/RESULT events with approval interrupts.
	//
	// example:
	//
	// call-1
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
	// The name of the tool for the TOOL_CALL_START event.
	//
	// example:
	//
	// example_tool
	ToolCallName *string `json:"toolCallName,omitempty" xml:"toolCallName,omitempty"`
	// The event type of the SSE data object. For the required fields of each event type, see the response description.
	//
	// This parameter is required.
	//
	// example:
	//
	// RUN_STARTED
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The extension value of the CUSTOM event. It can be any JSON object or null.
	//
	// example:
	//
	// {"reference":null,"actions":[],"messageId":null}
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InvokeAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeAgentResponseBody) GetContent() *string {
	return s.Content
}

func (s *InvokeAgentResponseBody) GetDelta() *string {
	return s.Delta
}

func (s *InvokeAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeAgentResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *InvokeAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *InvokeAgentResponseBody) GetOutcome() *InvokeAgentResponseBodyOutcome {
	return s.Outcome
}

func (s *InvokeAgentResponseBody) GetParentMessageId() *string {
	return s.ParentMessageId
}

func (s *InvokeAgentResponseBody) GetRole() *string {
	return s.Role
}

func (s *InvokeAgentResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *InvokeAgentResponseBody) GetSnapshot() interface{} {
	return s.Snapshot
}

func (s *InvokeAgentResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *InvokeAgentResponseBody) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *InvokeAgentResponseBody) GetToolCallName() *string {
	return s.ToolCallName
}

func (s *InvokeAgentResponseBody) GetType() *string {
	return s.Type
}

func (s *InvokeAgentResponseBody) GetValue() interface{} {
	return s.Value
}

func (s *InvokeAgentResponseBody) SetCode(v string) *InvokeAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeAgentResponseBody) SetContent(v string) *InvokeAgentResponseBody {
	s.Content = &v
	return s
}

func (s *InvokeAgentResponseBody) SetDelta(v string) *InvokeAgentResponseBody {
	s.Delta = &v
	return s
}

func (s *InvokeAgentResponseBody) SetMessage(v string) *InvokeAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeAgentResponseBody) SetMessageId(v string) *InvokeAgentResponseBody {
	s.MessageId = &v
	return s
}

func (s *InvokeAgentResponseBody) SetName(v string) *InvokeAgentResponseBody {
	s.Name = &v
	return s
}

func (s *InvokeAgentResponseBody) SetOutcome(v *InvokeAgentResponseBodyOutcome) *InvokeAgentResponseBody {
	s.Outcome = v
	return s
}

func (s *InvokeAgentResponseBody) SetParentMessageId(v string) *InvokeAgentResponseBody {
	s.ParentMessageId = &v
	return s
}

func (s *InvokeAgentResponseBody) SetRole(v string) *InvokeAgentResponseBody {
	s.Role = &v
	return s
}

func (s *InvokeAgentResponseBody) SetRunId(v string) *InvokeAgentResponseBody {
	s.RunId = &v
	return s
}

func (s *InvokeAgentResponseBody) SetSnapshot(v interface{}) *InvokeAgentResponseBody {
	s.Snapshot = v
	return s
}

func (s *InvokeAgentResponseBody) SetThreadId(v string) *InvokeAgentResponseBody {
	s.ThreadId = &v
	return s
}

func (s *InvokeAgentResponseBody) SetToolCallId(v string) *InvokeAgentResponseBody {
	s.ToolCallId = &v
	return s
}

func (s *InvokeAgentResponseBody) SetToolCallName(v string) *InvokeAgentResponseBody {
	s.ToolCallName = &v
	return s
}

func (s *InvokeAgentResponseBody) SetType(v string) *InvokeAgentResponseBody {
	s.Type = &v
	return s
}

func (s *InvokeAgentResponseBody) SetValue(v interface{}) *InvokeAgentResponseBody {
	s.Value = v
	return s
}

func (s *InvokeAgentResponseBody) Validate() error {
	if s.Outcome != nil {
		if err := s.Outcome.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeAgentResponseBodyOutcome struct {
	// The list of interrupts waiting for a response. This field appears only when outcome.type is interrupt and contains at least one item.
	//
	// example:
	//
	// [{"id":"i-1","reason":"tool_call_approve","toolCallId":"call-1"}]
	Interrupts []*InvokeAgentResponseBodyOutcomeInterrupts `json:"interrupts,omitempty" xml:"interrupts,omitempty" type:"Repeated"`
	// The outcome type. A value of success indicates normal completion. A value of interrupt indicates waiting for a response, in which case the interrupts array contains at least one item.
	//
	// This parameter is required.
	//
	// example:
	//
	// interrupt
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvokeAgentResponseBodyOutcome) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentResponseBodyOutcome) GoString() string {
	return s.String()
}

func (s *InvokeAgentResponseBodyOutcome) GetInterrupts() []*InvokeAgentResponseBodyOutcomeInterrupts {
	return s.Interrupts
}

func (s *InvokeAgentResponseBodyOutcome) GetType() *string {
	return s.Type
}

func (s *InvokeAgentResponseBodyOutcome) SetInterrupts(v []*InvokeAgentResponseBodyOutcomeInterrupts) *InvokeAgentResponseBodyOutcome {
	s.Interrupts = v
	return s
}

func (s *InvokeAgentResponseBodyOutcome) SetType(v string) *InvokeAgentResponseBodyOutcome {
	s.Type = &v
	return s
}

func (s *InvokeAgentResponseBodyOutcome) Validate() error {
	if s.Interrupts != nil {
		for _, item := range s.Interrupts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InvokeAgentResponseBodyOutcomeInterrupts struct {
	// The optional expiration time string of the interrupt. This field is not output for current tool approvals.
	//
	// example:
	//
	// 2026-11-04T00:00:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The ID of the interrupt. It is used as resume[].interruptId when resuming the run.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The approval description or follow-up question text displayed to the user.
	//
	// example:
	//
	// Proceed with execution?
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The optional extension information of the interrupt. It cannot replace tool events.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The reason for the interrupt. The value is tool_call_approve for tool approval and ask_human for follow-up questions.
	//
	// This parameter is required.
	//
	// example:
	//
	// tool_call_approve
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The JSON Schema for the response. Tool approval requires decision=yes/no, and follow-up questions use strings.
	//
	// example:
	//
	// {"type":"object","properties":{"decision":{"type":"string","enum":["yes","no"]}},"required":["decision"]}
	ResponseSchema map[string]interface{} `json:"responseSchema,omitempty" xml:"responseSchema,omitempty"`
	// The optional associated tool invoke ID. It is the same as the toolCallId of the preceding TOOL_CALL_START event.
	//
	// example:
	//
	// call-1
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
}

func (s InvokeAgentResponseBodyOutcomeInterrupts) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentResponseBodyOutcomeInterrupts) GoString() string {
	return s.String()
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetId() *string {
	return s.Id
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetMessage() *string {
	return s.Message
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetReason() *string {
	return s.Reason
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetResponseSchema() map[string]interface{} {
	return s.ResponseSchema
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetExpiresAt(v string) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.ExpiresAt = &v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetId(v string) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.Id = &v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetMessage(v string) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.Message = &v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetMetadata(v map[string]interface{}) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.Metadata = v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetReason(v string) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.Reason = &v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetResponseSchema(v map[string]interface{}) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.ResponseSchema = v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) SetToolCallId(v string) *InvokeAgentResponseBodyOutcomeInterrupts {
	s.ToolCallId = &v
	return s
}

func (s *InvokeAgentResponseBodyOutcomeInterrupts) Validate() error {
	return dara.Validate(s)
}

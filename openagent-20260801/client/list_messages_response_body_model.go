// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMessagesResponseBodyData) *ListMessagesResponseBody
	GetData() *ListMessagesResponseBodyData
	SetRequestId(v string) *ListMessagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMessagesResponseBody
	GetSuccess() *bool
}

type ListMessagesResponseBody struct {
	// The message history and pagination information.
	Data *ListMessagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBody) GetData() *ListMessagesResponseBodyData {
	return s.Data
}

func (s *ListMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMessagesResponseBody) SetData(v *ListMessagesResponseBodyData) *ListMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListMessagesResponseBody) SetRequestId(v string) *ListMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessagesResponseBody) SetSuccess(v bool) *ListMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessagesResponseBodyData struct {
	// The pagination token for the next page. A value of null indicates that no more data is available.
	//
	// example:
	//
	// example-next-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the session to which the messages belong.
	//
	// example:
	//
	// session_example_001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The message records returned in a timeline.
	Timeline []*ListMessagesResponseBodyDataTimeline `json:"timeline,omitempty" xml:"timeline,omitempty" type:"Repeated"`
	// The number of timeline records on the current page, which is equal to the length of the timeline array.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMessagesResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *ListMessagesResponseBodyData) GetTimeline() []*ListMessagesResponseBodyDataTimeline {
	return s.Timeline
}

func (s *ListMessagesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMessagesResponseBodyData) SetNextToken(v string) *ListMessagesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetSessionId(v string) *ListMessagesResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetTimeline(v []*ListMessagesResponseBodyDataTimeline) *ListMessagesResponseBodyData {
	s.Timeline = v
	return s
}

func (s *ListMessagesResponseBodyData) SetTotal(v int32) *ListMessagesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMessagesResponseBodyData) Validate() error {
	if s.Timeline != nil {
		for _, item := range s.Timeline {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimeline struct {
	// The message object. Different message types are distinguished by role, and each type uses the corresponding fields.
	Message *ListMessagesResponseBodyDataTimelineMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
	// The ID of the run to which the message belongs.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
}

func (s ListMessagesResponseBodyDataTimeline) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimeline) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimeline) GetMessage() *ListMessagesResponseBodyDataTimelineMessage {
	return s.Message
}

func (s *ListMessagesResponseBodyDataTimeline) GetRunId() *string {
	return s.RunId
}

func (s *ListMessagesResponseBodyDataTimeline) SetMessage(v *ListMessagesResponseBodyDataTimelineMessage) *ListMessagesResponseBodyDataTimeline {
	s.Message = v
	return s
}

func (s *ListMessagesResponseBodyDataTimeline) SetRunId(v string) *ListMessagesResponseBodyDataTimeline {
	s.RunId = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimeline) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimelineMessage struct {
	// The type identifier of the activity message.
	//
	// example:
	//
	// example_activity
	ActivityType *string `json:"activityType,omitempty" xml:"activityType,omitempty"`
	// The message content. The structure is determined by the message type.
	//
	// example:
	//
	// Query the status of the sample resource
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The encrypted content. The example is an invalid placeholder value and does not contain real ciphertext.
	//
	// example:
	//
	// EXAMPLE_ENCRYPTED_VALUE
	EncryptedValue *string `json:"encryptedValue,omitempty" xml:"encryptedValue,omitempty"`
	// The error description when tool execution fails.
	//
	// example:
	//
	// Sample tool execution failure
	Error *string `json:"error,omitempty" xml:"error,omitempty"`
	// The message ID.
	//
	// example:
	//
	// message_example_001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The result information of historical interrupts.
	InterruptOutcome *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome `json:"interruptOutcome,omitempty" xml:"interruptOutcome,omitempty" type:"Struct"`
	// The extended metadata of the message.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The name of the message sender.
	//
	// example:
	//
	// example-user
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resume input corresponding to the historical interrupt.
	ResumeInput *ListMessagesResponseBodyDataTimelineMessageResumeInput `json:"resumeInput,omitempty" xml:"resumeInput,omitempty" type:"Struct"`
	// The role of the message. Valid values: `user` (user message), `assistant` (assistant message), `tool` (tool result), `activity` (activity message), `reasoning` (reasoning message), and `interrupt` (interrupt record).
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The ID of the tool invocation associated with the tool result.
	//
	// example:
	//
	// tool_call_example_001
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
	// The list of tool invocations initiated by the agent.
	ToolCalls []*ListMessagesResponseBodyDataTimelineMessageToolCalls `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s ListMessagesResponseBodyDataTimelineMessage) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessage) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetActivityType() *string {
	return s.ActivityType
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetContent() interface{} {
	return s.Content
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetEncryptedValue() *string {
	return s.EncryptedValue
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetError() *string {
	return s.Error
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetId() *string {
	return s.Id
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetInterruptOutcome() *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome {
	return s.InterruptOutcome
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetName() *string {
	return s.Name
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetResumeInput() *ListMessagesResponseBodyDataTimelineMessageResumeInput {
	return s.ResumeInput
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetRole() *string {
	return s.Role
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *ListMessagesResponseBodyDataTimelineMessage) GetToolCalls() []*ListMessagesResponseBodyDataTimelineMessageToolCalls {
	return s.ToolCalls
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetActivityType(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.ActivityType = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetContent(v interface{}) *ListMessagesResponseBodyDataTimelineMessage {
	s.Content = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetEncryptedValue(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.EncryptedValue = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetError(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.Error = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetId(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.Id = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetInterruptOutcome(v *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) *ListMessagesResponseBodyDataTimelineMessage {
	s.InterruptOutcome = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetMetadata(v map[string]interface{}) *ListMessagesResponseBodyDataTimelineMessage {
	s.Metadata = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetName(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.Name = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetResumeInput(v *ListMessagesResponseBodyDataTimelineMessageResumeInput) *ListMessagesResponseBodyDataTimelineMessage {
	s.ResumeInput = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetRole(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.Role = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetToolCallId(v string) *ListMessagesResponseBodyDataTimelineMessage {
	s.ToolCallId = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) SetToolCalls(v []*ListMessagesResponseBodyDataTimelineMessageToolCalls) *ListMessagesResponseBodyDataTimelineMessage {
	s.ToolCalls = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessage) Validate() error {
	if s.InterruptOutcome != nil {
		if err := s.InterruptOutcome.Validate(); err != nil {
			return err
		}
	}
	if s.ResumeInput != nil {
		if err := s.ResumeInput.Validate(); err != nil {
			return err
		}
	}
	if s.ToolCalls != nil {
		for _, item := range s.ToolCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimelineMessageInterruptOutcome struct {
	// The list of historical interrupts.
	Interrupts []*ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts `json:"interrupts,omitempty" xml:"interrupts,omitempty" type:"Repeated"`
	// The type of the interrupt result. The value is fixed to `interrupt`.
	//
	// example:
	//
	// interrupt
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) GetInterrupts() []*ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	return s.Interrupts
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) GetType() *string {
	return s.Type
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) SetInterrupts(v []*ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome {
	s.Interrupts = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) SetType(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome {
	s.Type = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcome) Validate() error {
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

type ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts struct {
	// The supplementary business data for the interrupt.
	Data *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The expiration time of the interrupt, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The interrupt ID.
	//
	// example:
	//
	// interrupt_example_001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The interrupt prompt displayed to the user.
	//
	// example:
	//
	// Specify the target region
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The extended metadata of the interrupt.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The reason for the interrupt. Valid values: `tool_call` (tool invocation requires processing), `input_required` (user input is required), `confirmation` (confirmation is required), and `out_of_band_approval` (waiting for out-of-band approval).
	//
	// example:
	//
	// input_required
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The JSON schema that the resume result payload must follow.
	//
	// example:
	//
	// {"type":"object","properties":{"region":{"type":"string"}},"required":["region"]}
	ResponseSchema map[string]interface{} `json:"responseSchema,omitempty" xml:"responseSchema,omitempty"`
	// The ID of the tool invocation associated with the interrupt.
	//
	// example:
	//
	// tool_call_example_001
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetData() *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	return s.Data
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetId() *string {
	return s.Id
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetMessage() *string {
	return s.Message
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetReason() *string {
	return s.Reason
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetResponseSchema() map[string]interface{} {
	return s.ResponseSchema
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetData(v *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.Data = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetExpiresAt(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.ExpiresAt = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetId(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.Id = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetMessage(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.Message = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetMetadata(v map[string]interface{}) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.Metadata = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetReason(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.Reason = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetResponseSchema(v map[string]interface{}) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.ResponseSchema = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) SetToolCallId(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts {
	s.ToolCallId = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterrupts) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData struct {
	// The list of API operations that require approval.
	ApiCalls []*ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls `json:"apiCalls,omitempty" xml:"apiCalls,omitempty" type:"Repeated"`
	// The associated approval ID.
	//
	// example:
	//
	// approval_example_001
	ApprovalId *string `json:"approvalId,omitempty" xml:"approvalId,omitempty"`
	// The URL of the approval page. The example URL cannot be used for actual approval.
	//
	// example:
	//
	// https://example.com/approvals/approval_example_001
	ApprovalUrl *string `json:"approvalUrl,omitempty" xml:"approvalUrl,omitempty"`
	// Specifies whether the approval can be canceled.
	//
	// example:
	//
	// true
	Cancellable *bool `json:"cancellable,omitempty" xml:"cancellable,omitempty"`
	// The digest identifier of the change pending approval.
	//
	// example:
	//
	// example-change-digest
	ChangeDigest *string `json:"changeDigest,omitempty" xml:"changeDigest,omitempty"`
	// The title of the approval item.
	//
	// example:
	//
	// Sample change approval
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GetApiCalls() []*ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls {
	return s.ApiCalls
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GetApprovalUrl() *string {
	return s.ApprovalUrl
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GetCancellable() *bool {
	return s.Cancellable
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GetChangeDigest() *string {
	return s.ChangeDigest
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) GetTitle() *string {
	return s.Title
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) SetApiCalls(v []*ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	s.ApiCalls = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) SetApprovalId(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	s.ApprovalId = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) SetApprovalUrl(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	s.ApprovalUrl = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) SetCancellable(v bool) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	s.Cancellable = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) SetChangeDigest(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	s.ChangeDigest = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) SetTitle(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData {
	s.Title = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsData) Validate() error {
	if s.ApiCalls != nil {
		for _, item := range s.ApiCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls struct {
	// The name of the API action.
	//
	// example:
	//
	// ExampleAction
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The description of the operation pending approval.
	//
	// example:
	//
	// Execute changes on the sample resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) GetAction() *string {
	return s.Action
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) GetDescription() *string {
	return s.Description
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) SetAction(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls {
	s.Action = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) SetDescription(v string) *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls {
	s.Description = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageInterruptOutcomeInterruptsDataApiCalls) Validate() error {
	return dara.Validate(s)
}

type ListMessagesResponseBodyDataTimelineMessageResumeInput struct {
	// The time when the interrupt processing was completed, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:01:00Z
	ResolvedAt *string `json:"resolvedAt,omitempty" xml:"resolvedAt,omitempty"`
	// The submitted processing results for each interrupt.
	Resume []*ListMessagesResponseBodyDataTimelineMessageResumeInputResume `json:"resume,omitempty" xml:"resume,omitempty" type:"Repeated"`
	// The ID of the original run that was resumed.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The source of the resume request. Valid values: `USER` (submitted by a user) and `INTERNAL` (submitted by an internal process).
	//
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageResumeInput) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageResumeInput) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) GetResolvedAt() *string {
	return s.ResolvedAt
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) GetResume() []*ListMessagesResponseBodyDataTimelineMessageResumeInputResume {
	return s.Resume
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) GetRunId() *string {
	return s.RunId
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) GetSource() *string {
	return s.Source
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) SetResolvedAt(v string) *ListMessagesResponseBodyDataTimelineMessageResumeInput {
	s.ResolvedAt = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) SetResume(v []*ListMessagesResponseBodyDataTimelineMessageResumeInputResume) *ListMessagesResponseBodyDataTimelineMessageResumeInput {
	s.Resume = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) SetRunId(v string) *ListMessagesResponseBodyDataTimelineMessageResumeInput {
	s.RunId = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) SetSource(v string) *ListMessagesResponseBodyDataTimelineMessageResumeInput {
	s.Source = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInput) Validate() error {
	if s.Resume != nil {
		for _, item := range s.Resume {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimelineMessageResumeInputResume struct {
	// The ID of the processed interrupt.
	//
	// example:
	//
	// interrupt_example_001
	InterruptId *string `json:"interruptId,omitempty" xml:"interruptId,omitempty"`
	// The processing result for the corresponding interrupt, which must comply with the responseSchema of the interrupt.
	//
	// example:
	//
	// {"region":"cn-hangzhou"}
	Payload interface{} `json:"payload,omitempty" xml:"payload,omitempty"`
	// The response method for the current interrupt. Valid values: `resolved` (the answer, confirmation, or tool result required by the interrupt has been submitted, including the `payload`) and `cancelled` (the interrupt has been canceled, excluding the `payload`).
	//
	// example:
	//
	// resolved
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageResumeInputResume) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageResumeInputResume) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) GetInterruptId() *string {
	return s.InterruptId
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) GetPayload() interface{} {
	return s.Payload
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) GetStatus() *string {
	return s.Status
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) SetInterruptId(v string) *ListMessagesResponseBodyDataTimelineMessageResumeInputResume {
	s.InterruptId = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) SetPayload(v interface{}) *ListMessagesResponseBodyDataTimelineMessageResumeInputResume {
	s.Payload = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) SetStatus(v string) *ListMessagesResponseBodyDataTimelineMessageResumeInputResume {
	s.Status = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageResumeInputResume) Validate() error {
	return dara.Validate(s)
}

type ListMessagesResponseBodyDataTimelineMessageToolCalls struct {
	// The encrypted content of the tool invocation. The example is only a placeholder value.
	//
	// example:
	//
	// EXAMPLE_ENCRYPTED_VALUE
	EncryptedValue *string `json:"encryptedValue,omitempty" xml:"encryptedValue,omitempty"`
	// The information about the function to be executed.
	Function *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// The tool invocation ID.
	//
	// example:
	//
	// tool_call_example_001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The extended metadata of the tool invocation.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The type of the tool invocation. The value is fixed to `function`.
	//
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageToolCalls) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageToolCalls) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) GetEncryptedValue() *string {
	return s.EncryptedValue
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) GetFunction() *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction {
	return s.Function
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) GetId() *string {
	return s.Id
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) GetType() *string {
	return s.Type
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) SetEncryptedValue(v string) *ListMessagesResponseBodyDataTimelineMessageToolCalls {
	s.EncryptedValue = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) SetFunction(v *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) *ListMessagesResponseBodyDataTimelineMessageToolCalls {
	s.Function = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) SetId(v string) *ListMessagesResponseBodyDataTimelineMessageToolCalls {
	s.Id = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) SetMetadata(v map[string]interface{}) *ListMessagesResponseBodyDataTimelineMessageToolCalls {
	s.Metadata = v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) SetType(v string) *ListMessagesResponseBodyDataTimelineMessageToolCalls {
	s.Type = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCalls) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessagesResponseBodyDataTimelineMessageToolCallsFunction struct {
	// The tool parameters in the form of a JSON string.
	//
	// example:
	//
	// {"keyword":"example"}
	Arguments *string `json:"arguments,omitempty" xml:"arguments,omitempty"`
	// The name of the function or tool.
	//
	// example:
	//
	// lookup_example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) GetArguments() *string {
	return s.Arguments
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) GetName() *string {
	return s.Name
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) SetArguments(v string) *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction {
	s.Arguments = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) SetName(v string) *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction {
	s.Name = &v
	return s
}

func (s *ListMessagesResponseBodyDataTimelineMessageToolCallsFunction) Validate() error {
	return dara.Validate(s)
}

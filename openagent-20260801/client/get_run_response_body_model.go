// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRunResponseBodyData) *GetRunResponseBody
	GetData() *GetRunResponseBodyData
	SetRequestId(v string) *GetRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRunResponseBody
	GetSuccess() *bool
}

type GetRunResponseBody struct {
	// The returned data.
	Data *GetRunResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunResponseBody) GetData() *GetRunResponseBodyData {
	return s.Data
}

func (s *GetRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRunResponseBody) SetData(v *GetRunResponseBodyData) *GetRunResponseBody {
	s.Data = v
	return s
}

func (s *GetRunResponseBody) SetRequestId(v string) *GetRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunResponseBody) SetSuccess(v bool) *GetRunResponseBody {
	s.Success = &v
	return s
}

func (s *GetRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunResponseBodyData struct {
	// The time when the run was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The run outcome.
	Outcome *GetRunResponseBodyDataOutcome `json:"outcome,omitempty" xml:"outcome,omitempty" type:"Struct"`
	// The run ID.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// session_example_001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The run status: `PENDING` (pending execution), `RUNNING` (running), `WAITING_USER` (waiting for interrupt processing), `COMPLETED` (completed), `FAILED` (failed), `CANCELLED` (canceled), or `REJECTED` (rejected).
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the run was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:01:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetRunResponseBodyData) GetOutcome() *GetRunResponseBodyDataOutcome {
	return s.Outcome
}

func (s *GetRunResponseBodyData) GetRunId() *string {
	return s.RunId
}

func (s *GetRunResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetRunResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRunResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetRunResponseBodyData) SetCreatedAt(v string) *GetRunResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetRunResponseBodyData) SetOutcome(v *GetRunResponseBodyDataOutcome) *GetRunResponseBodyData {
	s.Outcome = v
	return s
}

func (s *GetRunResponseBodyData) SetRunId(v string) *GetRunResponseBodyData {
	s.RunId = &v
	return s
}

func (s *GetRunResponseBodyData) SetSessionId(v string) *GetRunResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetRunResponseBodyData) SetStatus(v string) *GetRunResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRunResponseBodyData) SetUpdatedAt(v string) *GetRunResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetRunResponseBodyData) Validate() error {
	if s.Outcome != nil {
		if err := s.Outcome.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunResponseBodyDataOutcome struct {
	// The list of interrupts.
	Interrupts []*GetRunResponseBodyDataOutcomeInterrupts `json:"interrupts,omitempty" xml:"interrupts,omitempty" type:"Repeated"`
	// The outcome type: `success` (completed successfully) or `interrupt` (waiting for interrupt processing).
	//
	// example:
	//
	// interrupt
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetRunResponseBodyDataOutcome) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponseBodyDataOutcome) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyDataOutcome) GetInterrupts() []*GetRunResponseBodyDataOutcomeInterrupts {
	return s.Interrupts
}

func (s *GetRunResponseBodyDataOutcome) GetType() *string {
	return s.Type
}

func (s *GetRunResponseBodyDataOutcome) SetInterrupts(v []*GetRunResponseBodyDataOutcomeInterrupts) *GetRunResponseBodyDataOutcome {
	s.Interrupts = v
	return s
}

func (s *GetRunResponseBodyDataOutcome) SetType(v string) *GetRunResponseBodyDataOutcome {
	s.Type = &v
	return s
}

func (s *GetRunResponseBodyDataOutcome) Validate() error {
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

type GetRunResponseBodyDataOutcomeInterrupts struct {
	// The interrupt data.
	Data *GetRunResponseBodyDataOutcomeInterruptsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The expiration time of the interrupt, in ISO 8601 format.
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
	// The interrupt prompt.
	//
	// example:
	//
	// Specify the region of the target resource
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The reason for the interrupt: `tool_call` (tool calling requires processing), `input_required` (user input is required), `confirmation` (confirmation is required), or `out_of_band_approval` (waiting for out-of-band approval).
	//
	// example:
	//
	// input_required
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The JSON Schema that the payload must conform to when resuming the interrupt.
	//
	// example:
	//
	// {"type":"object","properties":{"region":{"type":"string"}},"required":["region"]}
	ResponseSchema map[string]interface{} `json:"responseSchema,omitempty" xml:"responseSchema,omitempty"`
	// The ID used to invoke the tool calling.
	//
	// example:
	//
	// tool_call_example_001
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
}

func (s GetRunResponseBodyDataOutcomeInterrupts) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponseBodyDataOutcomeInterrupts) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetData() *GetRunResponseBodyDataOutcomeInterruptsData {
	return s.Data
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetId() *string {
	return s.Id
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetMessage() *string {
	return s.Message
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetReason() *string {
	return s.Reason
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetResponseSchema() map[string]interface{} {
	return s.ResponseSchema
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetData(v *GetRunResponseBodyDataOutcomeInterruptsData) *GetRunResponseBodyDataOutcomeInterrupts {
	s.Data = v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetExpiresAt(v string) *GetRunResponseBodyDataOutcomeInterrupts {
	s.ExpiresAt = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetId(v string) *GetRunResponseBodyDataOutcomeInterrupts {
	s.Id = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetMessage(v string) *GetRunResponseBodyDataOutcomeInterrupts {
	s.Message = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetMetadata(v map[string]interface{}) *GetRunResponseBodyDataOutcomeInterrupts {
	s.Metadata = v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetReason(v string) *GetRunResponseBodyDataOutcomeInterrupts {
	s.Reason = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetResponseSchema(v map[string]interface{}) *GetRunResponseBodyDataOutcomeInterrupts {
	s.ResponseSchema = v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) SetToolCallId(v string) *GetRunResponseBodyDataOutcomeInterrupts {
	s.ToolCallId = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterrupts) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunResponseBodyDataOutcomeInterruptsData struct {
	// The list of APIs pending approval.
	ApiCalls []*GetRunResponseBodyDataOutcomeInterruptsDataApiCalls `json:"apiCalls,omitempty" xml:"apiCalls,omitempty" type:"Repeated"`
	// The approval ID.
	//
	// example:
	//
	// approval_example_001
	ApprovalId *string `json:"approvalId,omitempty" xml:"approvalId,omitempty"`
	// The approval URL.
	//
	// example:
	//
	// https://example.com/approvals/approval_example_001
	ApprovalUrl *string `json:"approvalUrl,omitempty" xml:"approvalUrl,omitempty"`
	// Specifies whether the interrupt can be canceled.
	//
	// example:
	//
	// true
	Cancellable *bool `json:"cancellable,omitempty" xml:"cancellable,omitempty"`
	// The change digest.
	//
	// example:
	//
	// example-change-digest
	ChangeDigest *string `json:"changeDigest,omitempty" xml:"changeDigest,omitempty"`
	// The approval title.
	//
	// example:
	//
	// Sample change approval
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetRunResponseBodyDataOutcomeInterruptsData) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponseBodyDataOutcomeInterruptsData) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) GetApiCalls() []*GetRunResponseBodyDataOutcomeInterruptsDataApiCalls {
	return s.ApiCalls
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) GetApprovalUrl() *string {
	return s.ApprovalUrl
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) GetCancellable() *bool {
	return s.Cancellable
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) GetChangeDigest() *string {
	return s.ChangeDigest
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) GetTitle() *string {
	return s.Title
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) SetApiCalls(v []*GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) *GetRunResponseBodyDataOutcomeInterruptsData {
	s.ApiCalls = v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) SetApprovalId(v string) *GetRunResponseBodyDataOutcomeInterruptsData {
	s.ApprovalId = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) SetApprovalUrl(v string) *GetRunResponseBodyDataOutcomeInterruptsData {
	s.ApprovalUrl = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) SetCancellable(v bool) *GetRunResponseBodyDataOutcomeInterruptsData {
	s.Cancellable = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) SetChangeDigest(v string) *GetRunResponseBodyDataOutcomeInterruptsData {
	s.ChangeDigest = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) SetTitle(v string) *GetRunResponseBodyDataOutcomeInterruptsData {
	s.Title = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsData) Validate() error {
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

type GetRunResponseBodyDataOutcomeInterruptsDataApiCalls struct {
	// The API name.
	//
	// example:
	//
	// ExampleAction
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The operation description.
	//
	// example:
	//
	// Perform changes on the sample resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) GetAction() *string {
	return s.Action
}

func (s *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) GetDescription() *string {
	return s.Description
}

func (s *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) SetAction(v string) *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls {
	s.Action = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) SetDescription(v string) *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls {
	s.Description = &v
	return s
}

func (s *GetRunResponseBodyDataOutcomeInterruptsDataApiCalls) Validate() error {
	return dara.Validate(s)
}

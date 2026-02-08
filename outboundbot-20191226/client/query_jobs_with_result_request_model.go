// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsWithResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndActualTimeFilter(v int64) *QueryJobsWithResultRequest
	GetEndActualTimeFilter() *int64
	SetHasAnsweredFilter(v bool) *QueryJobsWithResultRequest
	GetHasAnsweredFilter() *bool
	SetHasHangUpByRejectionFilter(v bool) *QueryJobsWithResultRequest
	GetHasHangUpByRejectionFilter() *bool
	SetHasReachedEndOfFlowFilter(v bool) *QueryJobsWithResultRequest
	GetHasReachedEndOfFlowFilter() *bool
	SetInstanceId(v string) *QueryJobsWithResultRequest
	GetInstanceId() *string
	SetJobFailureReasonsFilter(v string) *QueryJobsWithResultRequest
	GetJobFailureReasonsFilter() *string
	SetJobGroupId(v string) *QueryJobsWithResultRequest
	GetJobGroupId() *string
	SetJobStatusFilter(v string) *QueryJobsWithResultRequest
	GetJobStatusFilter() *string
	SetLabelsJson(v []*string) *QueryJobsWithResultRequest
	GetLabelsJson() []*string
	SetPageNumber(v int32) *QueryJobsWithResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryJobsWithResultRequest
	GetPageSize() *int32
	SetQueryText(v string) *QueryJobsWithResultRequest
	GetQueryText() *string
	SetStartActualTimeFilter(v int64) *QueryJobsWithResultRequest
	GetStartActualTimeFilter() *int64
	SetTaskStatusFilter(v string) *QueryJobsWithResultRequest
	GetTaskStatusFilter() *string
}

type QueryJobsWithResultRequest struct {
	// Filter condition for call end time
	//
	// example:
	//
	// 1579055783000
	EndActualTimeFilter *int64 `json:"EndActualTimeFilter,omitempty" xml:"EndActualTimeFilter,omitempty"`
	// Is answered
	//
	// example:
	//
	// true
	HasAnsweredFilter *bool `json:"HasAnsweredFilter,omitempty" xml:"HasAnsweredFilter,omitempty"`
	// Whether the call was hung up due to rejection
	//
	// example:
	//
	// false
	HasHangUpByRejectionFilter *bool `json:"HasHangUpByRejectionFilter,omitempty" xml:"HasHangUpByRejectionFilter,omitempty"`
	// Is completed
	//
	// example:
	//
	// true
	HasReachedEndOfFlowFilter *bool `json:"HasReachedEndOfFlowFilter,omitempty" xml:"HasReachedEndOfFlowFilter,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 9d53cd72-4050-4419-8c17-acc0bf158147
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of job failure reasons
	//
	// example:
	//
	// ["NoAnswer"]
	JobFailureReasonsFilter *string `json:"JobFailureReasonsFilter,omitempty" xml:"JobFailureReasonsFilter,omitempty"`
	// Task group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ad16fc35-d824-4102-a606-2be51c1aa6dd
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job status filter
	//
	// - Scheduling: Scheduling.
	//
	// - Executing: Executing.
	//
	// - Succeeded: Ended – Reached.
	//
	// - Paused: Paused.
	//
	// - Failed: Ended – Not reached.
	//
	// - Cancelled: Cancelled – Manual intervention.
	//
	// example:
	//
	// Succeeded
	JobStatusFilter *string `json:"JobStatusFilter,omitempty" xml:"JobStatusFilter,omitempty"`
	// Tag filtering conditions for calls
	//
	// > This condition supports filtering only by labels that have specific enumeration tag values configured—that is, labels with explicitly defined tag values in the LLM scenario.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// Page number
	//
	// 	Notice: This parameter is Required
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// 	Notice: This parameter is Required
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Search content; you can search by phone number
	//
	// example:
	//
	// 1882020****
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// Filter condition for call start time
	//
	// example:
	//
	// 1579055782000
	StartActualTimeFilter *int64 `json:"StartActualTimeFilter,omitempty" xml:"StartActualTimeFilter,omitempty"`
	// Call status. Valid values: ["Executing","Succeeded"]. Separate multiple values with commas.
	//
	// Value range:
	//
	// (Note: The **Succeeded*	- status has been subdivided by reason. The generic **Succeeded**: 1 (Answered) status will not be displayed; instead, only the specific subdivided reason types will be returned.)
	//
	// - **Executing**: 0 (Dialing in progress).
	//
	// - **Succeeded**: 1 (Answered).
	//
	// - **NoAnswer**: 2 (Not answered – No one picked up).
	//
	// - **NotExist**: 3 (Not answered – Nonexistent number).
	//
	// - **Busy**: 4 (Not answered – Line busy).
	//
	// - **Cancelled**: 5 (Not dialed – Job stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not answered – Unable to connect).
	//
	// - **PoweredOff**: 8 (Not answered – Powered off).
	//
	// - **OutOfService**: 9 (Not answered – Called party suspended).
	//
	// - **InArrears**: 10 (Not answered – Called party has overdue payment).
	//
	// - **EmptyNumber**: 11 (Not dialed – Nonexistent number, no call placed).
	//
	// - **PerDayCallCountLimit**: 12 (Not dialed – Exceeded daily call limit).
	//
	// - **ContactBlockList**: 13 (Not dialed – Blacklist).
	//
	// - **CallerNotRegistered**: 14 (Not dialed – Caller number not registered).
	//
	// - **Terminated**: 15 (Not dialed – Terminated).
	//
	// - **VerificationCancelled**: 16 (Not dialed – Cancelled due to failed pre-call authentication).
	//
	// - **OutOfServiceNoCall**: 17 (Not dialed – Called party suspended, no call placed).
	//
	// - **InArrearsNoCall**: 18 (Not dialed – Called party has overdue payment, no call placed).
	//
	// - **CallingNumberNotExist**: 19 (Not dialed – Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Answered – Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Answered – Bot hung up after rejection).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Answered – Bot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Answered – User hung up after rejection).
	//
	// - **SucceededClientHangUp**: 24 (Answered – User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Answered – Transferred to agent due to intent hit).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Answered – Transferred to agent after rejection).
	//
	// - **SucceededInoInterAction**: 27 (Answered – No interaction from user side).
	//
	// - **SucceededError**: 28 (Answered – Interrupted by system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Answered – Intercepted under special condition – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Answered – Intercepted under special condition – Extension transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Answered – Intercepted under special condition – Custom intercept).
	//
	// - **HighRiskSipCode**: 32 (Not dialed – High-risk, no call placed).
	//
	// example:
	//
	// ["Executing"]
	TaskStatusFilter *string `json:"TaskStatusFilter,omitempty" xml:"TaskStatusFilter,omitempty"`
}

func (s QueryJobsWithResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultRequest) GetEndActualTimeFilter() *int64 {
	return s.EndActualTimeFilter
}

func (s *QueryJobsWithResultRequest) GetHasAnsweredFilter() *bool {
	return s.HasAnsweredFilter
}

func (s *QueryJobsWithResultRequest) GetHasHangUpByRejectionFilter() *bool {
	return s.HasHangUpByRejectionFilter
}

func (s *QueryJobsWithResultRequest) GetHasReachedEndOfFlowFilter() *bool {
	return s.HasReachedEndOfFlowFilter
}

func (s *QueryJobsWithResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryJobsWithResultRequest) GetJobFailureReasonsFilter() *string {
	return s.JobFailureReasonsFilter
}

func (s *QueryJobsWithResultRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *QueryJobsWithResultRequest) GetJobStatusFilter() *string {
	return s.JobStatusFilter
}

func (s *QueryJobsWithResultRequest) GetLabelsJson() []*string {
	return s.LabelsJson
}

func (s *QueryJobsWithResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryJobsWithResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryJobsWithResultRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *QueryJobsWithResultRequest) GetStartActualTimeFilter() *int64 {
	return s.StartActualTimeFilter
}

func (s *QueryJobsWithResultRequest) GetTaskStatusFilter() *string {
	return s.TaskStatusFilter
}

func (s *QueryJobsWithResultRequest) SetEndActualTimeFilter(v int64) *QueryJobsWithResultRequest {
	s.EndActualTimeFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetHasAnsweredFilter(v bool) *QueryJobsWithResultRequest {
	s.HasAnsweredFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetHasHangUpByRejectionFilter(v bool) *QueryJobsWithResultRequest {
	s.HasHangUpByRejectionFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetHasReachedEndOfFlowFilter(v bool) *QueryJobsWithResultRequest {
	s.HasReachedEndOfFlowFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetInstanceId(v string) *QueryJobsWithResultRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetJobFailureReasonsFilter(v string) *QueryJobsWithResultRequest {
	s.JobFailureReasonsFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetJobGroupId(v string) *QueryJobsWithResultRequest {
	s.JobGroupId = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetJobStatusFilter(v string) *QueryJobsWithResultRequest {
	s.JobStatusFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetLabelsJson(v []*string) *QueryJobsWithResultRequest {
	s.LabelsJson = v
	return s
}

func (s *QueryJobsWithResultRequest) SetPageNumber(v int32) *QueryJobsWithResultRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetPageSize(v int32) *QueryJobsWithResultRequest {
	s.PageSize = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetQueryText(v string) *QueryJobsWithResultRequest {
	s.QueryText = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetStartActualTimeFilter(v int64) *QueryJobsWithResultRequest {
	s.StartActualTimeFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetTaskStatusFilter(v string) *QueryJobsWithResultRequest {
	s.TaskStatusFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) Validate() error {
	return dara.Validate(s)
}

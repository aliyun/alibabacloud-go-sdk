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
	// Filters for calls that ended on or before the specified time. Specify the time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1579055783000
	EndActualTimeFilter *int64 `json:"EndActualTimeFilter,omitempty" xml:"EndActualTimeFilter,omitempty"`
	// Filters jobs by whether the call was answered.
	//
	// example:
	//
	// true
	HasAnsweredFilter *bool `json:"HasAnsweredFilter,omitempty" xml:"HasAnsweredFilter,omitempty"`
	// Filters jobs by whether the call was disconnected due to a rejection.
	//
	// example:
	//
	// false
	HasHangUpByRejectionFilter *bool `json:"HasHangUpByRejectionFilter,omitempty" xml:"HasHangUpByRejectionFilter,omitempty"`
	// Filters jobs by whether the call flow was completed.
	//
	// example:
	//
	// true
	HasReachedEndOfFlowFilter *bool `json:"HasReachedEndOfFlowFilter,omitempty" xml:"HasReachedEndOfFlowFilter,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9d53cd72-4050-4419-8c17-acc0bf158147
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job failure reasons to filter by.
	//
	// example:
	//
	// ["NoAnswer"]
	JobFailureReasonsFilter *string `json:"JobFailureReasonsFilter,omitempty" xml:"JobFailureReasonsFilter,omitempty"`
	// The ID of the job group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ad16fc35-d824-4102-a606-2be51c1aa6dd
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// The job status to filter by. Valid values:
	//
	// - `Scheduling`: The job is scheduled and awaiting execution.
	//
	// - `Executing`: The job is in progress.
	//
	// - `Succeeded`: The job is completed and the contact was reached.
	//
	// - `Paused`: The job is paused.
	//
	// - `Failed`: The job completed but failed to reach the contact.
	//
	// - `Cancelled`: The job was canceled by a user.
	//
	// example:
	//
	// Succeeded
	JobStatusFilter *string `json:"JobStatusFilter,omitempty" xml:"JobStatusFilter,omitempty"`
	// The filter conditions for calls, based on their labels.
	//
	// > This filter applies only to labels that are configured with a predefined set of values (enumerated values). These labels are typically used in large model scenarios.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// The page number.
	//
	// 	Notice:
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// 	Notice: This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search query for a specific job, such as the contact\\"s phone number.
	//
	// example:
	//
	// 1882020****
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// Filters for calls that started on or after the specified time. Specify the time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1579055782000
	StartActualTimeFilter *int64 `json:"StartActualTimeFilter,omitempty" xml:"StartActualTimeFilter,omitempty"`
	// The call statuses to filter by. You can specify multiple statuses as a JSON array of strings, such as `["Executing", "Succeeded"]`.
	//
	// Valid values:
	//
	// (Note: The **Succeeded*	- status is subdivided into more specific reasons. The general **Succeeded*	- (1: Connected) status is no longer returned. Instead, one of the more specific sub-statuses is returned.)
	//
	// - **Executing*	- (0): The call is being placed.
	//
	// - **Succeeded*	- (1): The call was connected.
	//
	// - **NoAnswer*	- (2): Not connected - No answer.
	//
	// - **NotExist*	- (3): Not connected - The dialed number does not exist.
	//
	// - **Busy*	- (4): Not connected - The line was busy.
	//
	// - **Cancelled*	- (5): Not placed - The job was stopped before the call could be dialed.
	//
	// - **Failed*	- (6): The call failed.
	//
	// - **NotConnected*	- (7): Not connected - The call could not be connected.
	//
	// - **PoweredOff*	- (8): Not connected - The recipient\\"s phone was powered off.
	//
	// - **OutOfService*	- (9): Not connected - The recipient\\"s number is out of service.
	//
	// - **InArrears*	- (10): Not connected - The recipient\\"s account is in arrears.
	//
	// - **EmptyNumber*	- (11): Not placed - The number was identified as an empty number and was not dialed.
	//
	// - **PerDayCallCountLimit*	- (12): Not placed - The daily call limit was reached.
	//
	// - **ContactBlockList*	- (13): Not placed - The number is on a blocklist.
	//
	// - **CallerNotRegistered*	- (14): Not placed - The calling number is not registered.
	//
	// - **Terminated*	- (15): Not placed - The call was terminated.
	//
	// - **VerificationCancelled*	- (16): Not placed - Canceled after failing a pre-call verification.
	//
	// - **OutOfServiceNoCall*	- (17): Not placed - The number is out of service and was not dialed.
	//
	// - **InArrearsNoCall*	- (18): Not placed - The recipient is in arrears and was not dialed.
	//
	// - **CallingNumberNotExist*	- (19): Not placed - The calling number does not exist.
	//
	// - **SucceededFinish*	- (20): Connected - The call completed normally.
	//
	// - **SucceededChatbotHangUpAfterNoAnswer*	- (21): Connected - The chatbot hung up after a rejection.
	//
	// - **SucceededChatbotHangUpAfterSilence*	- (22): Connected - The chatbot hung up due to a silence timeout.
	//
	// - **SucceededClientHangUpAfterNoAnswer*	- (23): Connected - The user hung up after a rejection.
	//
	// - **SucceededClientHangUp*	- (24): Connected - The user hung up for no specific reason.
	//
	// - **SucceededTransferByIntent*	- (25): Connected - The call was transferred to an agent based on user intent.
	//
	// - **SucceededTransferAfterNoAnswer*	- (26): Connected - The call was transferred to an agent after a rejection.
	//
	// - **SucceededInoInterAction*	- (27): Connected - No interaction from the user.
	//
	// - **SucceededError*	- (28): Connected - The call was interrupted by a system error.
	//
	// - **SucceededSpecialInterceptVoiceAssistant*	- (29): Connected - Intercepted by a voice assistant.
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer*	- (30): Connected - Intercepted for an extension number transfer.
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept*	- (31): Connected - Intercepted by a custom rule.
	//
	// - **HighRiskSipCode*	- (32): Not placed - High-risk SIP code.
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

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
	// The filter condition for the call end time.
	//
	// example:
	//
	// 1579055783000
	EndActualTimeFilter *int64 `json:"EndActualTimeFilter,omitempty" xml:"EndActualTimeFilter,omitempty"`
	// Specifies whether the call is answered.
	//
	// example:
	//
	// true
	HasAnsweredFilter *bool `json:"HasAnsweredFilter,omitempty" xml:"HasAnsweredFilter,omitempty"`
	// Specifies whether the call is hung up due to rejection.
	//
	// example:
	//
	// false
	HasHangUpByRejectionFilter *bool `json:"HasHangUpByRejectionFilter,omitempty" xml:"HasHangUpByRejectionFilter,omitempty"`
	// Specifies whether the call has reached the end of the flow.
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
	// The list of job failure reasons.
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
	// The job status filter. Valid values:
	//
	// - Scheduling: Scheduling in progress.
	//
	// - Executing: Executing in progress.
	//
	// - Succeeded: Ended - Reached.
	//
	// - Paused: Paused.
	//
	// - Failed: Ended - Not reached.
	//
	// - Cancelled: Cancelled - Manual intervention.
	//
	// example:
	//
	// Succeeded
	JobStatusFilter *string `json:"JobStatusFilter,omitempty" xml:"JobStatusFilter,omitempty"`
	// The label-based filter condition for calls.
	//
	// >This condition supports filtering only by labels that have specific enumerated label values configured, that is, labels with specific label values configured in large language model scenarios.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search content. You can search by phone number.
	//
	// example:
	//
	// 1882020****
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The filter condition for the call start time.
	//
	// example:
	//
	// 1579055782000
	StartActualTimeFilter *int64 `json:"StartActualTimeFilter,omitempty" xml:"StartActualTimeFilter,omitempty"`
	// The call status. Example: ["Executing","Succeeded"]. Separate multiple values with commas.
	//
	// Valid values:
	//
	// (Note: The **Succeeded*	- status has been subdivided into specific reasons. The general **Succeeded**: 1 (Connected) status is no longer returned. Instead, specific sub-reason types are returned.)
	//
	// - **Executing**: 0 (Calling).
	//
	// - **Succeeded**: 1 (Connected).
	//
	// - **NoAnswer**: 2 (Not connected - No answer).
	//
	// - **NotExist**: 3 (Not connected - Nonexistent number).
	//
	// - **Busy**: 4 (Not connected - Busy).
	//
	// - **Cancelled**: 5 (Not dialed - Task stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not connected - Unreachable).
	//
	// - **PoweredOff**: 8 (Not connected - Powered off).
	//
	// - **OutOfService**: 9 (Not connected - Callee out of service).
	//
	// - **InArrears**: 10 (Not connected - Callee has overdue payment).
	//
	// - **EmptyNumber**: 11 (Not dialed - Nonexistent number, no outbound call).
	//
	// - **PerDayCallCountLimit**: 12 (Not dialed - Daily limit exceeded).
	//
	// - **ContactBlockList**: 13 (Not dialed - Blacklisted).
	//
	// - **CallerNotRegistered**: 14 (Not dialed - Caller number not registered).
	//
	// - **Terminated**: 15 (Not dialed - Terminated).
	//
	// - **VerificationCancelled**: 16 (Not dialed - Pre-call verification failed, cancelled).
	//
	// - **OutOfServiceNoCall**: 17 (Not dialed - Callee out of service, no outbound call).
	//
	// - **InArrearsNoCall**: 18 (Not dialed - Callee has overdue payment, no outbound call).
	//
	// - **CallingNumberNotExist**: 19 (Not dialed - Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Connected - Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Connected - Robot hung up after no recognition).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Connected - Hung up due to silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Connected - User hung up after no recognition).
	//
	// - **SucceededClientHangUp**: 24 (Connected - User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Connected - Transferred to agent by intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Connected - Transferred to agent after no recognition).
	//
	// - **SucceededInoInterAction**: 27 (Connected - No interaction from user).
	//
	// - **SucceededError**: 28 (Connected - System exception interruption).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Connected - Special intercept - Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Connected - Special intercept - Extension number transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Connected - Special intercept - Custom intercept).
	//
	// - **HighRiskSipCode**: 32 (Not dialed - High risk, no outbound call).
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

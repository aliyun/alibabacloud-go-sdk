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
	// Specifies whether the call was answered.
	//
	// example:
	//
	// true
	HasAnsweredFilter *bool `json:"HasAnsweredFilter,omitempty" xml:"HasAnsweredFilter,omitempty"`
	// Specifies whether the call was hung up due to rejection.
	//
	// example:
	//
	// false
	HasHangUpByRejectionFilter *bool `json:"HasHangUpByRejectionFilter,omitempty" xml:"HasHangUpByRejectionFilter,omitempty"`
	// Specifies whether the call reached the end of the flow.
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
	// - Scheduling: scheduling.
	//
	// - Executing: executing.
	//
	// - Succeeded: ended - reached.
	//
	// - Paused: paused.
	//
	// - Failed: ended - not reached.
	//
	// - Cancelled: cancelled - manual intervention.
	//
	// example:
	//
	// Succeeded
	JobStatusFilter *string `json:"JobStatusFilter,omitempty" xml:"JobStatusFilter,omitempty"`
	// The filter condition for labels associated with calls.
	//
	// > This condition only supports filtering by labels that have specific enumerated label values configured, that is, labels with specific label values configured in large language model scenarios.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// The page number.
	//
	// 	Notice: This parameter is required.</notice>
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// 	Notice: This parameter is required.</notice>
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
	// The call status, such as ["Executing","Succeeded"]. Separate multiple values with commas (,).
	//
	// Valid values:
	//
	// (Note: The **Succeeded*	- status has been subdivided into specific reasons. The **Succeeded**: 1 (answered) status is no longer returned. Instead, specific sub-reason types are returned.)
	//
	// - **Executing**: 0 (dialing).
	//
	// - **Succeeded**: 1 (answered).
	//
	// - **NoAnswer**: 2 (not answered - no one picked up).
	//
	// - **NotExist**: 3 (not answered - nonexistent number).
	//
	// - **Busy**: 4 (not answered - busy).
	//
	// - **Cancelled**: 5 (not dialed - task stopped).
	//
	// - **Failed**: 6 (failed).
	//
	// - **NotConnected**: 7 (not answered - unreachable).
	//
	// - **PoweredOff**: 8 (not answered - powered off).
	//
	// - **OutOfService**: 9 (not answered - callee out of service).
	//
	// - **InArrears**: 10 (not answered - callee has overdue payment).
	//
	// - **EmptyNumber**: 11 (not dialed - nonexistent number, no outbound call).
	//
	// - **PerDayCallCountLimit**: 12 (not dialed - daily limit exceeded).
	//
	// - **ContactBlockList**: 13 (not dialed - blacklisted).
	//
	// - **CallerNotRegistered**: 14 (not dialed - caller number not registered).
	//
	// - **Terminated**: 15 (not dialed - terminated).
	//
	// - **VerificationCancelled**: 16 (not dialed - cancelled due to pre-call verification failure).
	//
	// - **OutOfServiceNoCall**: 17 (not dialed - callee out of service, no outbound call).
	//
	// - **InArrearsNoCall**: 18 (not dialed - callee has overdue payment, no outbound call).
	//
	// - **CallingNumberNotExist**: 19 (not dialed - caller number does not exist).
	//
	// - **SucceededFinish**: 20 (answered - completed normally).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (answered - robot hung up after rejection).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (answered - hung up due to silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (answered - user hung up after rejection).
	//
	// - **SucceededClientHangUp**: 24 (answered - user hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (answered - transferred to agent by intent).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (answered - transferred to agent after rejection).
	//
	// - **SucceededInoInterAction**: 27 (answered - no interaction from user side).
	//
	// - **SucceededError**: 28 (answered - interrupted by system error).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (answered - special interception - voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (answered - special interception - extension number transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (answered - special interception - custom interception).
	//
	// - **HighRiskSipCode**: 32 (not dialed - high risk, no outbound call).
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

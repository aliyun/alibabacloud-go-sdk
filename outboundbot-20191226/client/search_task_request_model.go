// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActualTimeGte(v int64) *SearchTaskRequest
	GetActualTimeGte() *int64
	SetActualTimeLte(v int64) *SearchTaskRequest
	GetActualTimeLte() *int64
	SetCallDurationGte(v int64) *SearchTaskRequest
	GetCallDurationGte() *int64
	SetCallDurationLte(v int64) *SearchTaskRequest
	GetCallDurationLte() *int64
	SetCalledNumber(v string) *SearchTaskRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *SearchTaskRequest
	GetCallingNumber() *string
	SetInstanceId(v string) *SearchTaskRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *SearchTaskRequest
	GetJobGroupId() *string
	SetJobGroupNameQuery(v string) *SearchTaskRequest
	GetJobGroupNameQuery() *string
	SetJobId(v string) *SearchTaskRequest
	GetJobId() *string
	SetJobStatusStringList(v string) *SearchTaskRequest
	GetJobStatusStringList() *string
	SetLabelsJson(v []*string) *SearchTaskRequest
	GetLabelsJson() []*string
	SetOtherId(v string) *SearchTaskRequest
	GetOtherId() *string
	SetPageIndex(v int32) *SearchTaskRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *SearchTaskRequest
	GetPageSize() *int32
	SetRecordingDurationGte(v int64) *SearchTaskRequest
	GetRecordingDurationGte() *int64
	SetRecordingDurationLte(v int64) *SearchTaskRequest
	GetRecordingDurationLte() *int64
	SetScriptNameQuery(v string) *SearchTaskRequest
	GetScriptNameQuery() *string
	SetSortBy(v string) *SearchTaskRequest
	GetSortBy() *string
	SetSortOrder(v string) *SearchTaskRequest
	GetSortOrder() *string
	SetTaskCreateTimeGte(v int64) *SearchTaskRequest
	GetTaskCreateTimeGte() *int64
	SetTaskCreateTimeLte(v int64) *SearchTaskRequest
	GetTaskCreateTimeLte() *int64
	SetTaskId(v string) *SearchTaskRequest
	GetTaskId() *string
	SetTaskStatusStringList(v string) *SearchTaskRequest
	GetTaskStatusStringList() *string
	SetUserIdMatch(v string) *SearchTaskRequest
	GetUserIdMatch() *string
}

type SearchTaskRequest struct {
	// The start time of the call.
	//
	// example:
	//
	// 1646582400000
	ActualTimeGte *int64 `json:"ActualTimeGte,omitempty" xml:"ActualTimeGte,omitempty"`
	// The end time of the call.
	//
	// example:
	//
	// 1643126399000
	ActualTimeLte *int64 `json:"ActualTimeLte,omitempty" xml:"ActualTimeLte,omitempty"`
	// The minimum call duration for the search, in milliseconds.
	//
	// example:
	//
	// 12341155
	CallDurationGte *int64 `json:"CallDurationGte,omitempty" xml:"CallDurationGte,omitempty"`
	// The maximum call duration for the search, in milliseconds.
	//
	// example:
	//
	// 12341155
	CallDurationLte *int64 `json:"CallDurationLte,omitempty" xml:"CallDurationLte,omitempty"`
	// The called number.
	//
	// example:
	//
	// 15126426342
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 051085500215
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d481cebe-0bb6-4d13-9649-42ce5074fb75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job group ID.
	//
	// example:
	//
	// 3a30ae7c-27b2-4305-9444-7185ced9d51f
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// The name of the job group.
	//
	// example:
	//
	// ChinaLifePCHospitalizationInsuranceRenewal_20220301_134130
	JobGroupNameQuery *string `json:"JobGroupNameQuery,omitempty" xml:"JobGroupNameQuery,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 11994321-e6bc-47bb-8b1c-8eef8f2f768b
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job statuses. Separate multiple statuses with commas (,). If this field is not empty, it overwrites jobStatusList. Valid values:
	//
	// - Scheduling: The job is being scheduled.
	//
	// - Executing: The job is being executed.
	//
	// - Succeeded: The job is finished and the called party is reached.
	//
	// - Paused: The job is stopped.
	//
	// - Failed: The job is finished but the called party is not reached.
	//
	// - Cancelled: The job is canceled.
	//
	// example:
	//
	// Succeeded
	JobStatusStringList *string `json:"JobStatusStringList,omitempty" xml:"JobStatusStringList,omitempty"`
	// The filter conditions for call labels.
	//
	// > You can only filter by labels that are configured with specific enumerated label values, which means labels configured with specific values in large language model scenarios.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// The other IDs.
	//
	// **This includes the following IDs:**
	//
	// - sessionID
	//
	// - taskid
	//
	// - jobid
	//
	// example:
	//
	// AVD-2021-39685
	OtherId *string `json:"OtherId,omitempty" xml:"OtherId,omitempty"`
	// The page number.
	//
	// > Page numbers for this API operation start from 0 by default.
	//
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page.
	//
	// > If this parameter is left empty, the default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The minimum ringing duration for the search.
	//
	// example:
	//
	// 10
	RecordingDurationGte *int64 `json:"RecordingDurationGte,omitempty" xml:"RecordingDurationGte,omitempty"`
	// The minimum ringing duration for the search.
	//
	// example:
	//
	// 60
	RecordingDurationLte *int64 `json:"RecordingDurationLte,omitempty" xml:"RecordingDurationLte,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// China Life P&C
	ScriptNameQuery *string `json:"ScriptNameQuery,omitempty" xml:"ScriptNameQuery,omitempty"`
	// The sort field. Default value: actualTime.
	//
	// example:
	//
	// actualTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	// - asc: Ascending order.
	//
	// - desc: Descending order. This is the default value.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The start time of the task.
	//
	// > This parameter must be specified together with the TaskCreateTimeLte parameter. Otherwise, it does not take effect.
	//
	// example:
	//
	// 1646792941
	TaskCreateTimeGte *int64 `json:"TaskCreateTimeGte,omitempty" xml:"TaskCreateTimeGte,omitempty"`
	// The end time of the task.
	//
	// > This parameter must be specified together with the TaskCreateTimeGte parameter. Otherwise, it does not take effect.
	//
	// example:
	//
	// 1646792941
	TaskCreateTimeLte *int64 `json:"TaskCreateTimeLte,omitempty" xml:"TaskCreateTimeLte,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 744b27f3-437f-4a8c-a181-f668e492fd24
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The call statuses. Separate multiple statuses with commas (,). Valid values:
	//
	// (Note: The **Succeeded*	- status is subdivided into specific causes. The **Succeeded**: 1 (Connected) status is not displayed. All subsequent returned values are specific subdivided cause types.)
	//
	// - **Executing**: 0 (Dialing).
	//
	// - **Succeeded**: 1 (Connected).
	//
	// - **NoAnswer**: 2 (Not connected - No answer).
	//
	// - **NotExist**: 3 (Not connected - Nonexistent number).
	//
	// - **Busy**: 4 (Not connected - Busy).
	//
	// - **Cancelled**: 5 (Not called - Task stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not connected - Unreachable).
	//
	// - **PoweredOff**: 8 (Not connected - Powered off).
	//
	// - **OutOfService**: 9 (Not connected - Suspended).
	//
	// - **InArrears**: 10 (Not connected - Overdue payment).
	//
	// - **EmptyNumber**: 11 (Not called - Nonexistent number, no call).
	//
	// - **PerDayCallCountLimit**: 12 (Not called - Daily limit exceeded).
	//
	// - **ContactBlockList**: 13 (Not called - Blacklisted).
	//
	// - **CallerNotRegistered**: 14 (Not called - Calling number not registered).
	//
	// - **Terminated**: 15 (Not called - Terminated).
	//
	// - **VerificationCancelled**: 16 (Not called - Canceled due to pre-call verification failure).
	//
	// - **OutOfServiceNoCall**: 17 (Not called - Suspended, no call).
	//
	// - **InArrearsNoCall**: 18 (Not called - Overdue payment, no call).
	//
	// - **CallingNumberNotExist**: 19 (Not called - Calling number does not exist).
	//
	// - **SucceededFinish**: 20 (Connected - Finished normally).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Connected - Bot hung up after unrecognized intent).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Connected - Hung up due to silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Connected - User hung up after unrecognized intent).
	//
	// - **SucceededClientHangUp**: 24 (Connected - User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Connected - Transferred to agent due to matched intent).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Connected - Transferred to agent after unrecognized intent).
	//
	// - **SucceededInoInterAction**: 27 (Connected - No interaction on the user side).
	//
	// - **SucceededError**: 28 (Connected - Interrupted due to system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Connected - Special interception - Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Connected - Special interception - Extension number transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Connected - Special interception - Custom interception).
	//
	// - **HighRiskSipCode**: 32 (Not called - High risk, no call).
	//
	// example:
	//
	// Executing
	TaskStatusStringList *string `json:"TaskStatusStringList,omitempty" xml:"TaskStatusStringList,omitempty"`
	// The user ID, which is the unique identifier of the user.
	//
	// > This field is passed when you upload an outbound call list.
	//
	// > - If the list is uploaded in JSON format, the user ID is the value of referenceId.
	//
	// > - If the list is uploaded in an Excel file, the user ID is the value of contactId.
	//
	// example:
	//
	// C01
	UserIdMatch *string `json:"UserIdMatch,omitempty" xml:"UserIdMatch,omitempty"`
}

func (s SearchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTaskRequest) GoString() string {
	return s.String()
}

func (s *SearchTaskRequest) GetActualTimeGte() *int64 {
	return s.ActualTimeGte
}

func (s *SearchTaskRequest) GetActualTimeLte() *int64 {
	return s.ActualTimeLte
}

func (s *SearchTaskRequest) GetCallDurationGte() *int64 {
	return s.CallDurationGte
}

func (s *SearchTaskRequest) GetCallDurationLte() *int64 {
	return s.CallDurationLte
}

func (s *SearchTaskRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SearchTaskRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *SearchTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchTaskRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *SearchTaskRequest) GetJobGroupNameQuery() *string {
	return s.JobGroupNameQuery
}

func (s *SearchTaskRequest) GetJobId() *string {
	return s.JobId
}

func (s *SearchTaskRequest) GetJobStatusStringList() *string {
	return s.JobStatusStringList
}

func (s *SearchTaskRequest) GetLabelsJson() []*string {
	return s.LabelsJson
}

func (s *SearchTaskRequest) GetOtherId() *string {
	return s.OtherId
}

func (s *SearchTaskRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *SearchTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTaskRequest) GetRecordingDurationGte() *int64 {
	return s.RecordingDurationGte
}

func (s *SearchTaskRequest) GetRecordingDurationLte() *int64 {
	return s.RecordingDurationLte
}

func (s *SearchTaskRequest) GetScriptNameQuery() *string {
	return s.ScriptNameQuery
}

func (s *SearchTaskRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchTaskRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *SearchTaskRequest) GetTaskCreateTimeGte() *int64 {
	return s.TaskCreateTimeGte
}

func (s *SearchTaskRequest) GetTaskCreateTimeLte() *int64 {
	return s.TaskCreateTimeLte
}

func (s *SearchTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SearchTaskRequest) GetTaskStatusStringList() *string {
	return s.TaskStatusStringList
}

func (s *SearchTaskRequest) GetUserIdMatch() *string {
	return s.UserIdMatch
}

func (s *SearchTaskRequest) SetActualTimeGte(v int64) *SearchTaskRequest {
	s.ActualTimeGte = &v
	return s
}

func (s *SearchTaskRequest) SetActualTimeLte(v int64) *SearchTaskRequest {
	s.ActualTimeLte = &v
	return s
}

func (s *SearchTaskRequest) SetCallDurationGte(v int64) *SearchTaskRequest {
	s.CallDurationGte = &v
	return s
}

func (s *SearchTaskRequest) SetCallDurationLte(v int64) *SearchTaskRequest {
	s.CallDurationLte = &v
	return s
}

func (s *SearchTaskRequest) SetCalledNumber(v string) *SearchTaskRequest {
	s.CalledNumber = &v
	return s
}

func (s *SearchTaskRequest) SetCallingNumber(v string) *SearchTaskRequest {
	s.CallingNumber = &v
	return s
}

func (s *SearchTaskRequest) SetInstanceId(v string) *SearchTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTaskRequest) SetJobGroupId(v string) *SearchTaskRequest {
	s.JobGroupId = &v
	return s
}

func (s *SearchTaskRequest) SetJobGroupNameQuery(v string) *SearchTaskRequest {
	s.JobGroupNameQuery = &v
	return s
}

func (s *SearchTaskRequest) SetJobId(v string) *SearchTaskRequest {
	s.JobId = &v
	return s
}

func (s *SearchTaskRequest) SetJobStatusStringList(v string) *SearchTaskRequest {
	s.JobStatusStringList = &v
	return s
}

func (s *SearchTaskRequest) SetLabelsJson(v []*string) *SearchTaskRequest {
	s.LabelsJson = v
	return s
}

func (s *SearchTaskRequest) SetOtherId(v string) *SearchTaskRequest {
	s.OtherId = &v
	return s
}

func (s *SearchTaskRequest) SetPageIndex(v int32) *SearchTaskRequest {
	s.PageIndex = &v
	return s
}

func (s *SearchTaskRequest) SetPageSize(v int32) *SearchTaskRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTaskRequest) SetRecordingDurationGte(v int64) *SearchTaskRequest {
	s.RecordingDurationGte = &v
	return s
}

func (s *SearchTaskRequest) SetRecordingDurationLte(v int64) *SearchTaskRequest {
	s.RecordingDurationLte = &v
	return s
}

func (s *SearchTaskRequest) SetScriptNameQuery(v string) *SearchTaskRequest {
	s.ScriptNameQuery = &v
	return s
}

func (s *SearchTaskRequest) SetSortBy(v string) *SearchTaskRequest {
	s.SortBy = &v
	return s
}

func (s *SearchTaskRequest) SetSortOrder(v string) *SearchTaskRequest {
	s.SortOrder = &v
	return s
}

func (s *SearchTaskRequest) SetTaskCreateTimeGte(v int64) *SearchTaskRequest {
	s.TaskCreateTimeGte = &v
	return s
}

func (s *SearchTaskRequest) SetTaskCreateTimeLte(v int64) *SearchTaskRequest {
	s.TaskCreateTimeLte = &v
	return s
}

func (s *SearchTaskRequest) SetTaskId(v string) *SearchTaskRequest {
	s.TaskId = &v
	return s
}

func (s *SearchTaskRequest) SetTaskStatusStringList(v string) *SearchTaskRequest {
	s.TaskStatusStringList = &v
	return s
}

func (s *SearchTaskRequest) SetUserIdMatch(v string) *SearchTaskRequest {
	s.UserIdMatch = &v
	return s
}

func (s *SearchTaskRequest) Validate() error {
	return dara.Validate(s)
}

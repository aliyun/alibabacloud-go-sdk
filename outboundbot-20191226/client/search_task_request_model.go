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
	// Call start time
	//
	// example:
	//
	// 1646582400000
	ActualTimeGte *int64 `json:"ActualTimeGte,omitempty" xml:"ActualTimeGte,omitempty"`
	// Call end time
	//
	// example:
	//
	// 1643126399000
	ActualTimeLte *int64 `json:"ActualTimeLte,omitempty" xml:"ActualTimeLte,omitempty"`
	// Minimum call duration, in milliseconds
	//
	// example:
	//
	// 12341155
	CallDurationGte *int64 `json:"CallDurationGte,omitempty" xml:"CallDurationGte,omitempty"`
	// Maximum call duration, in milliseconds
	//
	// example:
	//
	// 12341155
	CallDurationLte *int64 `json:"CallDurationLte,omitempty" xml:"CallDurationLte,omitempty"`
	// Called number
	//
	// example:
	//
	// 15126426342
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number
	//
	// example:
	//
	// 051085500215
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// d481cebe-0bb6-4d13-9649-42ce5074fb75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job group ID
	//
	// example:
	//
	// 3a30ae7c-27b2-4305-9444-7185ced9d51f
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job group name
	//
	// example:
	//
	// 国寿财住院保续保_20220301_134130
	JobGroupNameQuery *string `json:"JobGroupNameQuery,omitempty" xml:"JobGroupNameQuery,omitempty"`
	// Job ID
	//
	// example:
	//
	// 11994321-e6bc-47bb-8b1c-8eef8f2f768b
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Job status. Separate multiple statuses with commas. If you specify this parameter, it overrides jobStatusList.
	//
	// - Scheduling: The job is being scheduled.
	//
	// - Executing: The job is running.
	//
	// - Succeeded: The job completed successfully.
	//
	// - Paused: The job was paused.
	//
	// - Failed: The job failed.
	//
	// - Cancelled: The job was cancelled.
	//
	// example:
	//
	// Succeeded
	JobStatusStringList *string `json:"JobStatusStringList,omitempty" xml:"JobStatusStringList,omitempty"`
	// Label-based filter condition for calls
	//
	// > You can only use labels that have specific enumeration values. For example, labels configured with specific values in Large Language Model (LLM) scenarios.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// Other ID
	//
	// **Valid values include the following:**
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
	// Page number
	//
	// > The first page is 0.
	//
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of items per page
	//
	// > If you omit this parameter, the default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Minimum ring duration, in seconds
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
	// Scenario name
	//
	// example:
	//
	// 国寿财
	ScriptNameQuery *string `json:"ScriptNameQuery,omitempty" xml:"ScriptNameQuery,omitempty"`
	// Sort field. Default value: actualTime
	//
	// example:
	//
	// actualTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Sort order. Valid values:
	//
	// - asc (ascending)
	//
	// - desc (descending). Default value.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Start time of the task
	//
	// > You must specify both TaskCreateTimeGte and TaskCreateTimeLte. If you omit either, the filter does not work.
	//
	// example:
	//
	// 1646792941
	TaskCreateTimeGte *int64 `json:"TaskCreateTimeGte,omitempty" xml:"TaskCreateTimeGte,omitempty"`
	// End time of the task
	//
	// > You must specify both TaskCreateTimeGte and TaskCreateTimeLte. If you omit either, the filter does not work.
	//
	// example:
	//
	// 1646792941
	TaskCreateTimeLte *int64 `json:"TaskCreateTimeLte,omitempty" xml:"TaskCreateTimeLte,omitempty"`
	// Task ID
	//
	// example:
	//
	// 744b27f3-437f-4a8c-a181-f668e492fd24
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Call status. Separate multiple statuses with commas.
	//
	// - **Executing**: 0 (Calling).
	//
	// - **Succeeded**: 1 (Connected).
	//
	// - **NoAnswer**: 2 (No answer).
	//
	// - **NotExist**: 3 (Nonexistent number).
	//
	// - **Busy**: 4 (Line busy).
	//
	// - **Cancelled**: 5 (Call canceled due to job stop).
	//
	// - **Failed**: 6 (Call failed).
	//
	// - **NotConnected**: 7 (Cannot connect).
	//
	// - **PoweredOff**: 8 (Phone powered off).
	//
	// - **OutOfService**: 9 (Called number out of service).
	//
	// - **InArrears**: 10 (Called number overdue payment).
	//
	// - **EmptyNumber**: 11 (Empty number, no outbound call).
	//
	// - **PerDayCallCountLimit**: 12 (Daily call limit exceeded).
	//
	// - **ContactBlockList**: 13 (Blacklisted).
	//
	// - **CallerNotRegistered**: 14 (Caller number not registered).
	//
	// - **Terminated**: 15 (Call terminated).
	//
	// - **VerificationCancelled**: 16 (Call canceled due to pre-call validation failure).
	//
	// - **OutOfServiceNoCall**: 17 (Called number out of service, no outbound call).
	//
	// - **InArrearsNoCall**: 18 (Called number overdue payment, no outbound call).
	//
	// - **CallingNumberNotExist**: 19 (Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Connected and ended normally).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Connected and robot hung up after rejection).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Connected and robot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Connected and client hung up after rejection).
	//
	// - **SucceededClientHangUp**: 24 (Connected and client hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Connected and transferred to agent based on intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Connected and transferred to agent after rejection).
	//
	// - **SucceededInoInterAction**: 27 (Connected and no interaction from client side).
	//
	// - **SucceededError**: 28 (Connected but system error interrupted).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Connected but intercepted by voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Connected but intercepted by extension transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Connected but intercepted by custom rule).
	//
	// - **HighRiskSipCode**: 32 (High-risk SIP code, no outbound call).
	//
	// example:
	//
	// Executing
	TaskStatusStringList *string `json:"TaskStatusStringList,omitempty" xml:"TaskStatusStringList,omitempty"`
	// User ID. A unique identifier for a user.
	//
	// > This field is passed when you upload an outbound call list.
	//
	// >
	//
	// > - If you upload the list in JSON format, the user ID is the value of referenceId.
	//
	// >
	//
	// > - If you upload the list as an Excel file, the user ID is the value of contactId.
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

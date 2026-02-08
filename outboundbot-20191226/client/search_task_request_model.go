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
	// Call Start Time
	//
	// example:
	//
	// 1646582400000
	ActualTimeGte *int64 `json:"ActualTimeGte,omitempty" xml:"ActualTimeGte,omitempty"`
	// Call End Time
	//
	// example:
	//
	// 1643126399000
	ActualTimeLte *int64 `json:"ActualTimeLte,omitempty" xml:"ActualTimeLte,omitempty"`
	// Minimum call duration for search, in milliseconds
	//
	// example:
	//
	// 12341155
	CallDurationGte *int64 `json:"CallDurationGte,omitempty" xml:"CallDurationGte,omitempty"`
	// Maximum call duration for search, in milliseconds
	//
	// example:
	//
	// 12341155
	CallDurationLte *int64 `json:"CallDurationLte,omitempty" xml:"CallDurationLte,omitempty"`
	// Called Number
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
	// Task group ID
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
	// Job status. Separate multiple values with commas. When this field is not empty, it overwrites jobStatusList.
	//
	// - Scheduling: Scheduling.
	//
	// - Executing: Executing.
	//
	// - Succeeded: End – Reached.
	//
	// - Paused: Paused.
	//
	// - Failed: End – Not Reached.
	//
	// - Cancelled: Cancelled.
	//
	// example:
	//
	// Succeeded
	JobStatusStringList *string `json:"JobStatusStringList,omitempty" xml:"JobStatusStringList,omitempty"`
	// Tag filtering conditions for calls
	//
	// > This condition supports filtering only by tags that have specific enumeration tag values configured, specifically those tags with explicit tag values configured in the LLM scenario.
	LabelsJson []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// Other ID
	//
	// **Includes the following:**
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
	// Page index
	//
	// > The page index for this API starts from 0 by default.
	//
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of items
	//
	// > When this parameter is empty, the default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Minimum ringing duration for search
	//
	// example:
	//
	// 10
	RecordingDurationGte *int64 `json:"RecordingDurationGte,omitempty" xml:"RecordingDurationGte,omitempty"`
	// Maximum ringing duration for search
	//
	// example:
	//
	// 60
	RecordingDurationLte *int64 `json:"RecordingDurationLte,omitempty" xml:"RecordingDurationLte,omitempty"`
	// Scenario Name
	//
	// example:
	//
	// 国寿财
	ScriptNameQuery *string `json:"ScriptNameQuery,omitempty" xml:"ScriptNameQuery,omitempty"`
	// Sorting method. The default is actualTime.
	//
	// example:
	//
	// actualTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Sorting order. Valid values:
	//
	// - asc (ascending)
	//
	// - desc (default): descending
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Job start time
	//
	// > This parameter must be used together with the job end time parameter TaskCreateTimeLte. Otherwise, it will be invalid.
	//
	// example:
	//
	// 1646792941
	TaskCreateTimeGte *int64 `json:"TaskCreateTimeGte,omitempty" xml:"TaskCreateTimeGte,omitempty"`
	// Job end time
	//
	// > This parameter must be used together with the job start time parameter TaskCreateTimeGte. Otherwise, it will be invalid.
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
	// Call status. Separate multiple values with commas.
	//
	// Value range:
	//
	// (Note: The **Succeeded*	- status has been subdivided by reason. The generic **Succeeded**: 1 (Answered) status will not be returned. Instead, only the specific subdivided reason types will be returned.)
	//
	// - **Executing**: 0 (Calling).
	//
	// - **Succeeded**: 1 (Answered).
	//
	// - **NoAnswer**: 2 (Not answered – No response).
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
	// - **OutOfService**: 9 (Not answered – Called party out of service).
	//
	// - **InArrears**: 10 (Not answered – Called party has overdue payment).
	//
	// - **EmptyNumber**: 11 (Not dialed – Nonexistent number, no call made).
	//
	// - **PerDayCallCountLimit**: 12 (Not dialed – Exceeded daily call limit).
	//
	// - **ContactBlockList**: 13 (Not dialed – Blacklist).
	//
	// - **CallerNotRegistered**: 14 (Not dialed – Caller number not registered).
	//
	// - **Terminated**: 15 (Not dialed – Terminated).
	//
	// - **VerificationCancelled**: 16 (Not dialed – Cancelled due to pre-call authentication failure).
	//
	// - **OutOfServiceNoCall**: 17 (Not dialed – Called party out of service, no call made).
	//
	// - **InArrearsNoCall**: 18 (Not dialed – Called party has overdue payment, no call made).
	//
	// - **CallingNumberNotExist**: 19 (Not dialed – Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Answered – Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Answered – Bot hung up after unrecognized input).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Answered – Bot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Answered – User hung up after unrecognized input).
	//
	// - **SucceededClientHangUp**: 24 (Answered – User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Answered – Transferred to agent due to intent hit).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Answered – Transferred to agent after unrecognized input).
	//
	// - **SucceededInoInterAction**: 27 (Answered – No interaction from user side).
	//
	// - **SucceededError**: 28 (Answered – Interrupted due to system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Answered – Intercepted under special condition – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Answered – Intercepted under special condition – Extension transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Answered – Intercepted under special condition – Custom intercept).
	//
	// - **HighRiskSipCode**: 32 (Not dialed – High-risk, no call made)
	//
	// example:
	//
	// Executing
	TaskStatusStringList *string `json:"TaskStatusStringList,omitempty" xml:"TaskStatusStringList,omitempty"`
	// User ID, the unique identifier of the user
	//
	// > This field is provided when uploading an outbound call list.
	//
	// > - If uploaded via JSON, the User ID corresponds to the value of referenceId.
	//
	// > - If uploaded via an Excel file, the User ID corresponds to the value of contactId.
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

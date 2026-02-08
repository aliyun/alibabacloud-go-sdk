// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SearchTaskResponseBody
	GetHttpStatusCode() *int32
	SetLabels(v []*SearchTaskResponseBodyLabels) *SearchTaskResponseBody
	GetLabels() []*SearchTaskResponseBodyLabels
	SetMessage(v string) *SearchTaskResponseBody
	GetMessage() *string
	SetPageIndex(v int32) *SearchTaskResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *SearchTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchTaskResponseBody
	GetRequestId() *string
	SetSearchTaskInfoList(v []*SearchTaskResponseBodySearchTaskInfoList) *SearchTaskResponseBody
	GetSearchTaskInfoList() []*SearchTaskResponseBodySearchTaskInfoList
	SetSuccess(v bool) *SearchTaskResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *SearchTaskResponseBody
	GetTotal() *int64
	SetVariableNames(v []*string) *SearchTaskResponseBody
	GetVariableNames() []*string
}

type SearchTaskResponseBody struct {
	// Request status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Tag information that can be used as filter conditions.
	//
	// > Displays all tags with enumeration values in this task group.
	Labels []*SearchTaskResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 0
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Job list
	//
	// example:
	//
	// []
	SearchTaskInfoList []*SearchTaskResponseBodySearchTaskInfoList `json:"SearchTaskInfoList,omitempty" xml:"SearchTaskInfoList,omitempty" type:"Repeated"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Full set of tag name keys
	VariableNames []*string `json:"VariableNames,omitempty" xml:"VariableNames,omitempty" type:"Repeated"`
}

func (s SearchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchTaskResponseBody) GetLabels() []*SearchTaskResponseBodyLabels {
	return s.Labels
}

func (s *SearchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchTaskResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *SearchTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTaskResponseBody) GetSearchTaskInfoList() []*SearchTaskResponseBodySearchTaskInfoList {
	return s.SearchTaskInfoList
}

func (s *SearchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchTaskResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchTaskResponseBody) GetVariableNames() []*string {
	return s.VariableNames
}

func (s *SearchTaskResponseBody) SetCode(v string) *SearchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTaskResponseBody) SetHttpStatusCode(v int32) *SearchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchTaskResponseBody) SetLabels(v []*SearchTaskResponseBodyLabels) *SearchTaskResponseBody {
	s.Labels = v
	return s
}

func (s *SearchTaskResponseBody) SetMessage(v string) *SearchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTaskResponseBody) SetPageIndex(v int32) *SearchTaskResponseBody {
	s.PageIndex = &v
	return s
}

func (s *SearchTaskResponseBody) SetPageSize(v int32) *SearchTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchTaskResponseBody) SetRequestId(v string) *SearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTaskResponseBody) SetSearchTaskInfoList(v []*SearchTaskResponseBodySearchTaskInfoList) *SearchTaskResponseBody {
	s.SearchTaskInfoList = v
	return s
}

func (s *SearchTaskResponseBody) SetSuccess(v bool) *SearchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTaskResponseBody) SetTotal(v int64) *SearchTaskResponseBody {
	s.Total = &v
	return s
}

func (s *SearchTaskResponseBody) SetVariableNames(v []*string) *SearchTaskResponseBody {
	s.VariableNames = v
	return s
}

func (s *SearchTaskResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchTaskInfoList != nil {
		for _, item := range s.SearchTaskInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTaskResponseBodyLabels struct {
	// Tag Name
	//
	// example:
	//
	// 是否满意
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of tag values
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s SearchTaskResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s SearchTaskResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *SearchTaskResponseBodyLabels) GetName() *string {
	return s.Name
}

func (s *SearchTaskResponseBodyLabels) GetValueList() []*string {
	return s.ValueList
}

func (s *SearchTaskResponseBodyLabels) SetName(v string) *SearchTaskResponseBodyLabels {
	s.Name = &v
	return s
}

func (s *SearchTaskResponseBodyLabels) SetValueList(v []*string) *SearchTaskResponseBodyLabels {
	s.ValueList = v
	return s
}

func (s *SearchTaskResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}

type SearchTaskResponseBodySearchTaskInfoList struct {
	// Actual running time
	//
	// example:
	//
	// 1643436089677
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Call duration, in milliseconds.
	//
	// example:
	//
	// 46000
	CallDuration *int32 `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	// Call duration
	//
	// example:
	//
	// 47″
	CallDurationDisplay *string `json:"CallDurationDisplay,omitempty" xml:"CallDurationDisplay,omitempty"`
	// Called number
	//
	// example:
	//
	// 15205879599
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number
	//
	// example:
	//
	// 02152739269
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Abnormal condition
	//
	// example:
	//
	// [{"code":"OutboundCallError.SipCodeError", "params":[{"key":"SipCode","value":"500"}]}]
	DialException *string `json:"DialException,omitempty" xml:"DialException,omitempty"`
	// Field used for displaying remarks on the page
	DialExceptionCodes []*string `json:"DialExceptionCodes,omitempty" xml:"DialExceptionCodes,omitempty" type:"Repeated"`
	// Abnormal condition
	//
	// example:
	//
	// ["OutboundCallError.SipTrunkError"]
	DialExceptionOld *string `json:"DialExceptionOld,omitempty" xml:"DialExceptionOld,omitempty"`
	// Whether there is an acknowledgement
	//
	// example:
	//
	// true
	HasAnswered *bool `json:"HasAnswered,omitempty" xml:"HasAnswered,omitempty"`
	// Hang up due to rejection
	//
	// example:
	//
	// true
	HasHangUpByRejection *bool `json:"HasHangUpByRejection,omitempty" xml:"HasHangUpByRejection,omitempty"`
	// Whether the last playback was completed before hang-up
	//
	// example:
	//
	// true
	HasLastPlaybackCompleted *bool `json:"HasLastPlaybackCompleted,omitempty" xml:"HasLastPlaybackCompleted,omitempty"`
	// Conversation completed
	//
	// example:
	//
	// true
	HasReachedEndOfFlow *bool `json:"HasReachedEndOfFlow,omitempty" xml:"HasReachedEndOfFlow,omitempty"`
	// Instance ID
	//
	// example:
	//
	// 73df6283-26b2-402d-bad0-ffa489923ea1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job group ID
	//
	// example:
	//
	// 37db3113-ad34-4ba3-b930-468f016bbf95
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job group name
	//
	// example:
	//
	// 85成新（有笔记划线）2021_11_14_18_00_51
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// Job ID
	//
	// example:
	//
	// 6203248e-e652-4ef8-a1eb-586ed7b54dc2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Job status
	//
	//
	//
	// - 0: Scheduling
	//
	// - 1: Executing
	//
	// - 2: Succeeded
	//
	// - 3: Paused
	//
	// - 4: Failed
	//
	// - 5: Cancelled
	//
	// example:
	//
	// 0
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// Job status display content
	//
	// - 0: Scheduling
	//
	// - 1: Executing
	//
	// - 2: Succeeded
	//
	// - 3: Paused
	//
	// - 4: Failed
	//
	// - 5: Cancelled
	//
	// example:
	//
	// 调度中
	JobStatusName *string `json:"JobStatusName,omitempty" xml:"JobStatusName,omitempty"`
	// Job Status
	//
	// - Scheduling: Scheduling (0)
	//
	// - Executing: Executing (1)
	//
	// - Succeeded: End – Reached (2)
	//
	// - Paused: Paused (3)
	//
	// - Failed: Failed – Busy (4)
	//
	// - Cancelled: Cancelled (5)
	//
	// example:
	//
	// Scheduling
	JobStatusString *string `json:"JobStatusString,omitempty" xml:"JobStatusString,omitempty"`
	// The label hit status of the current outbound call
	Labels []*SearchTaskResponseBodySearchTaskInfoListLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Ringing duration, in seconds.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// Scenario Name
	//
	// example:
	//
	// 慢病线索
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Job creation time
	//
	// example:
	//
	// 1646792941
	TaskCreateTime *int64 `json:"TaskCreateTime,omitempty" xml:"TaskCreateTime,omitempty"`
	// Reason for job termination
	//
	// - FINISHED(1, "Normal completion"),
	//
	// - CHATBOT_HANGUP_AFTER_NOANSWER(2, "Chatbot hang-up after no answer"),
	//
	// - CHATBOT_HANGUP_AFTER_SILENCE(3, "Chatbot hang-up due to silence timeout"),
	//
	// - CLIENT_HANGUP_AFTER_NOANSWER(4, "User hang-up after no answer"),
	//
	// - CLIENT_HANGUP(5, "User hang-up without reason"),
	//
	// - TRANSFER_BY_INTENT(6, "Transfer to agent after intent hit"),
	//
	// - TRANSFER_AFTER_NOANSWER(7, "Transfer to agent after no answer"),
	//
	// - NO_INTERACTION(8, "No interaction from user side"),
	//
	// - ERROR(9, "System Exception break"),
	//
	// - SPECIAL_INTERCEPT_VOICE_ASSISTANT(10, "Special case interception – voice assistant"),
	//
	// - SPECIAL_INTERCEPT_EXTENSION_NUMBER_TRANSFER(11, "Special case interception – extension number transfer");
	//
	// example:
	//
	// OutOfService
	TaskEndReason *int32 `json:"TaskEndReason,omitempty" xml:"TaskEndReason,omitempty"`
	// Job ID
	//
	// example:
	//
	// 479aea04-3a92-4ac3-935d-c8798c667850
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Value range
	//
	// (Note: The **Succeeded*	- status has been further categorized by reason. The generic **Succeeded**: 1 (Connected) status will not be displayed. Subsequent responses will only show specific sub-reason types.)
	//
	// - **Executing**: 0 (Calling in progress).
	//
	// - **Succeeded**: 1 (Connected).
	//
	// - **NoAnswer**: 2 (Not connected – No answer).
	//
	// - **NotExist**: 3 (Not connected – Nonexistent number).
	//
	// - **Busy**: 4 (Not connected – Line busy).
	//
	// - **Cancelled**: 5 (Not dialed – Job stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not connected – Unable to reach).
	//
	// - **PoweredOff**: 8 (Not connected – Powered off).
	//
	// - **OutOfService**: 9 (Not connected – Called party suspended).
	//
	// - **InArrears**: 10 (Not connected – Called party has overdue payment).
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
	// - **VerificationCancelled**: 16 (Not dialed – Cancelled due to failed pre-call authentication).
	//
	// - **OutOfServiceNoCall**: 17 (Not dialed – Called party suspended, no call made).
	//
	// - **InArrearsNoCall**: 18 (Not dialed – Called party has overdue payment, no call made).
	//
	// - **CallingNumberNotExist**: 19 (Not dialed – Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Connected – Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Connected – Bot hung up after unrecognized input).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Connected – Bot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Connected – User hung up after unrecognized input).
	//
	// - **SucceededClientHangUp**: 24 (Connected – User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Connected – Transferred to agent due to intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Connected – Transferred to agent after unrecognized input).
	//
	// - **SucceededInoInterAction**: 27 (Connected – No interaction from user side).
	//
	// - **SucceededError**: 28 (Connected – Interrupted due to system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Connected – Blocked under special circumstances – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Connected – Blocked under special circumstances – Extension transfer).
	//
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Call status display content
	//
	// Value range
	//
	// (Note: The **Succeeded*	- status has been further categorized by reason and will no longer be displayed as **Succeeded**: 1 (Answered). Subsequent responses will use specific detailed reason types.)
	//
	// - **Executing**: 0 (Calling in progress).
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
	// - **OutOfService**: 9 (Not answered – Called party suspended).
	//
	// - **InArrears**: 10 (Not answered – Called party has overdue payment).
	//
	// - **EmptyNumber**: 11 (Not dialed – Nonexistent number, not called).
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
	// - **OutOfServiceNoCall**: 17 (Not dialed – Called party suspended, not called).
	//
	// - **InArrearsNoCall**: 18 (Not dialed – Called party has overdue payment, not called).
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
	// - **SucceededTransferByIntent**: 25 (Answered – Transferred to agent due to intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Answered – Transferred to agent after unrecognized input).
	//
	// - **SucceededInoInterAction**: 27 (Answered – No interaction from user side).
	//
	// - **SucceededError**: 28 (Answered – Call interrupted due to system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Answered – Intercepted under special circumstances – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Answered – Intercepted under special circumstances – Extension transfer).
	//
	// example:
	//
	// 正在拨打
	TaskStatusName *string `json:"TaskStatusName,omitempty" xml:"TaskStatusName,omitempty"`
	// Task Status
	//
	// Value range:
	//
	// (Note: The **Succeeded*	- status has been further categorized by reason. The generic **Succeeded**: 1 (Connected) status will no longer be displayed. Subsequent responses will always return a specific detailed reason.)
	//
	// - **Executing**: 0 (Calling in progress).
	//
	// - **Succeeded**: 1 (Connected).
	//
	// - **NoAnswer**: 2 (Not connected – No answer).
	//
	// - **NotExist**: 3 (Not connected – Nonexistent number).
	//
	// - **Busy**: 4 (Not connected – Line busy).
	//
	// - **Cancelled**: 5 (Not dialed – Job stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not connected – Unable to reach).
	//
	// - **PoweredOff**: 8 (Not connected – Powered off).
	//
	// - **OutOfService**: 9 (Not connected – Called party suspended).
	//
	// - **InArrears**: 10 (Not connected – Called party has overdue payment).
	//
	// - **EmptyNumber**: 11 (Not dialed – Nonexistent number, not called).
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
	// - **OutOfServiceNoCall**: 17 (Not dialed – Called party suspended, not called).
	//
	// - **InArrearsNoCall**: 18 (Not dialed – Called party has overdue payment, not called).
	//
	// - **CallingNumberNotExist**: 19 (Not dialed – Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Connected – Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Connected – Bot hung up after unrecognized input).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Connected – Bot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Connected – User hung up after unrecognized input).
	//
	// - **SucceededClientHangUp**: 24 (Connected – User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Connected – Transferred to agent due to intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Connected – Transferred to agent after unrecognized input).
	//
	// - **SucceededInoInterAction**: 27 (Connected – No interaction from user side).
	//
	// - **SucceededError**: 28 (Connected – Interrupted due to system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Connected – Intercepted under special condition – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Connected – Intercepted under special condition – Extension transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Connected – Intercepted under special condition – Custom intercept).
	//
	// - **HighRiskSipCode**: 32 (Not dialed – High-risk, not called).
	//
	// example:
	//
	// Executing
	TaskStatusString *string `json:"TaskStatusString,omitempty" xml:"TaskStatusString,omitempty"`
	// User ID
	//
	// example:
	//
	// 12334134
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Username
	//
	// example:
	//
	// xxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SearchTaskResponseBodySearchTaskInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchTaskResponseBodySearchTaskInfoList) GoString() string {
	return s.String()
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetCallDuration() *int32 {
	return s.CallDuration
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetCallDurationDisplay() *string {
	return s.CallDurationDisplay
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetDialException() *string {
	return s.DialException
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetDialExceptionCodes() []*string {
	return s.DialExceptionCodes
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetDialExceptionOld() *string {
	return s.DialExceptionOld
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetHasAnswered() *bool {
	return s.HasAnswered
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetHasHangUpByRejection() *bool {
	return s.HasHangUpByRejection
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetHasLastPlaybackCompleted() *bool {
	return s.HasLastPlaybackCompleted
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetHasReachedEndOfFlow() *bool {
	return s.HasReachedEndOfFlow
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetJobId() *string {
	return s.JobId
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetJobStatusName() *string {
	return s.JobStatusName
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetJobStatusString() *string {
	return s.JobStatusString
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetLabels() []*SearchTaskResponseBodySearchTaskInfoListLabels {
	return s.Labels
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetScriptName() *string {
	return s.ScriptName
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetTaskCreateTime() *int64 {
	return s.TaskCreateTime
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetTaskEndReason() *int32 {
	return s.TaskEndReason
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetTaskId() *string {
	return s.TaskId
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetTaskStatusName() *string {
	return s.TaskStatusName
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetTaskStatusString() *string {
	return s.TaskStatusString
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetUserId() *string {
	return s.UserId
}

func (s *SearchTaskResponseBodySearchTaskInfoList) GetUserName() *string {
	return s.UserName
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetActualTime(v int64) *SearchTaskResponseBodySearchTaskInfoList {
	s.ActualTime = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetCallDuration(v int32) *SearchTaskResponseBodySearchTaskInfoList {
	s.CallDuration = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetCallDurationDisplay(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.CallDurationDisplay = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetCalledNumber(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.CalledNumber = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetCallingNumber(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.CallingNumber = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetDialException(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.DialException = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetDialExceptionCodes(v []*string) *SearchTaskResponseBodySearchTaskInfoList {
	s.DialExceptionCodes = v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetDialExceptionOld(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.DialExceptionOld = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetHasAnswered(v bool) *SearchTaskResponseBodySearchTaskInfoList {
	s.HasAnswered = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetHasHangUpByRejection(v bool) *SearchTaskResponseBodySearchTaskInfoList {
	s.HasHangUpByRejection = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetHasLastPlaybackCompleted(v bool) *SearchTaskResponseBodySearchTaskInfoList {
	s.HasLastPlaybackCompleted = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetHasReachedEndOfFlow(v bool) *SearchTaskResponseBodySearchTaskInfoList {
	s.HasReachedEndOfFlow = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetInstanceId(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.InstanceId = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetJobGroupId(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.JobGroupId = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetJobGroupName(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.JobGroupName = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetJobId(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.JobId = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetJobStatus(v int32) *SearchTaskResponseBodySearchTaskInfoList {
	s.JobStatus = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetJobStatusName(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.JobStatusName = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetJobStatusString(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.JobStatusString = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetLabels(v []*SearchTaskResponseBodySearchTaskInfoListLabels) *SearchTaskResponseBodySearchTaskInfoList {
	s.Labels = v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetRecordingDuration(v int32) *SearchTaskResponseBodySearchTaskInfoList {
	s.RecordingDuration = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetScriptName(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.ScriptName = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetTaskCreateTime(v int64) *SearchTaskResponseBodySearchTaskInfoList {
	s.TaskCreateTime = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetTaskEndReason(v int32) *SearchTaskResponseBodySearchTaskInfoList {
	s.TaskEndReason = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetTaskId(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.TaskId = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetTaskStatus(v int32) *SearchTaskResponseBodySearchTaskInfoList {
	s.TaskStatus = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetTaskStatusName(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.TaskStatusName = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetTaskStatusString(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.TaskStatusString = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetUserId(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.UserId = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) SetUserName(v string) *SearchTaskResponseBodySearchTaskInfoList {
	s.UserName = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoList) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTaskResponseBodySearchTaskInfoListLabels struct {
	// Tag Name
	//
	// example:
	//
	// 是否满意
	K *string `json:"K,omitempty" xml:"K,omitempty"`
	// The hit tag value
	//
	// example:
	//
	// 是
	V *string `json:"V,omitempty" xml:"V,omitempty"`
}

func (s SearchTaskResponseBodySearchTaskInfoListLabels) String() string {
	return dara.Prettify(s)
}

func (s SearchTaskResponseBodySearchTaskInfoListLabels) GoString() string {
	return s.String()
}

func (s *SearchTaskResponseBodySearchTaskInfoListLabels) GetK() *string {
	return s.K
}

func (s *SearchTaskResponseBodySearchTaskInfoListLabels) GetV() *string {
	return s.V
}

func (s *SearchTaskResponseBodySearchTaskInfoListLabels) SetK(v string) *SearchTaskResponseBodySearchTaskInfoListLabels {
	s.K = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoListLabels) SetV(v string) *SearchTaskResponseBodySearchTaskInfoListLabels {
	s.V = &v
	return s
}

func (s *SearchTaskResponseBodySearchTaskInfoListLabels) Validate() error {
	return dara.Validate(s)
}

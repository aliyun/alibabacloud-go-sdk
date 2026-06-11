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
	// Labels available for filtering.
	//
	// > Displays all labels with enumeration values in this job group.
	Labels []*SearchTaskResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Response message
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
	// List of tasks
	//
	// example:
	//
	// []
	SearchTaskInfoList []*SearchTaskResponseBodySearchTaskInfoList `json:"SearchTaskInfoList,omitempty" xml:"SearchTaskInfoList,omitempty" type:"Repeated"`
	// Indicates whether the request succeeded
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
	// Complete list of label keys
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
	// Label name
	//
	// example:
	//
	// 是否满意
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of label values
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
	// Actual execution time
	//
	// example:
	//
	// 1643436089677
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Call duration, in milliseconds
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
	// Exception details
	//
	// example:
	//
	// [{"code":"OutboundCallError.SipCodeError", "params":[{"key":"SipCode","value":"500"}]}]
	DialException *string `json:"DialException,omitempty" xml:"DialException,omitempty"`
	// Exception codes shown on the remarks page
	DialExceptionCodes []*string `json:"DialExceptionCodes,omitempty" xml:"DialExceptionCodes,omitempty" type:"Repeated"`
	// Exception details
	//
	// example:
	//
	// ["OutboundCallError.SipTrunkError"]
	DialExceptionOld *string `json:"DialExceptionOld,omitempty" xml:"DialExceptionOld,omitempty"`
	// Indicates whether the called party answered
	//
	// example:
	//
	// true
	HasAnswered *bool `json:"HasAnswered,omitempty" xml:"HasAnswered,omitempty"`
	// Indicates whether the call ended due to rejection
	//
	// example:
	//
	// true
	HasHangUpByRejection *bool `json:"HasHangUpByRejection,omitempty" xml:"HasHangUpByRejection,omitempty"`
	// Indicates whether the last audio playback completed before hangup
	//
	// example:
	//
	// true
	HasLastPlaybackCompleted *bool `json:"HasLastPlaybackCompleted,omitempty" xml:"HasLastPlaybackCompleted,omitempty"`
	// Indicates whether the conversation ended
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
	// Display text for job status
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
	// Job status
	//
	// - Scheduling (0)
	//
	// - Executing (1)
	//
	// - Completed—Reached (2)
	//
	// - Paused (3)
	//
	// - Failed—Line busy (4)
	//
	// - Cancelled (5)
	//
	// example:
	//
	// Scheduling
	JobStatusString *string `json:"JobStatusString,omitempty" xml:"JobStatusString,omitempty"`
	// Labels matched for this outbound call
	Labels []*SearchTaskResponseBodySearchTaskInfoListLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Ringing duration, in seconds
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// Scenario name
	//
	// example:
	//
	// 慢病线索
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Task creation time
	//
	// example:
	//
	// 1646792941
	TaskCreateTime *int64 `json:"TaskCreateTime,omitempty" xml:"TaskCreateTime,omitempty"`
	// Reason why the task ended
	//
	// - FINISHED(1,"Normal completion")
	//
	// - CHATBOT_HANGUP_AFTER_NOANSWER(2, "Robot hangup after rejection")
	//
	// - CHATBOT_HANGUP_AFTER_SILENCE(3, "Robot hangup after silence timeout")
	//
	// - CLIENT_HANGUP_AFTER_NOANSWER(4, "Client hangup after rejection")
	//
	// - CLIENT_HANGUP(5, "Client hangup without reason")
	//
	// - TRANSFER_BY_INTENT(6, "Transfer to agent based on intent match")
	//
	// - TRANSFER_AFTER_NOANSWER(7, "Transfer to agent after rejection")
	//
	// - INO_INTERACTION(8, "No interaction from client side")
	//
	// - ERROR(9, "System error interrupted")
	//
	// - SPECIAL_INTERCEPT_VOICE_ASSISTANT(10, "Intercepted by voice assistant")
	//
	// - SPECIAL_INTERCEPT_EXTENSION_NUMBER_TRANSFER(11, "Intercepted by extension transfer")
	//
	// example:
	//
	// OutOfService
	TaskEndReason *int32 `json:"TaskEndReason,omitempty" xml:"TaskEndReason,omitempty"`
	// Task ID
	//
	// example:
	//
	// 479aea04-3a92-4ac3-935d-c8798c667850
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Valid values:
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
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Display text for task status
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
	// example:
	//
	// 正在拨打
	TaskStatusName *string `json:"TaskStatusName,omitempty" xml:"TaskStatusName,omitempty"`
	// Task status
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
	// Label name
	//
	// example:
	//
	// 是否满意
	K *string `json:"K,omitempty" xml:"K,omitempty"`
	// Matched label value
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

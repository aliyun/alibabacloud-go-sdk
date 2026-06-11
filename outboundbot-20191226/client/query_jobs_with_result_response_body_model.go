// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsWithResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryJobsWithResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QueryJobsWithResultResponseBody
	GetHttpStatusCode() *int32
	SetJobs(v *QueryJobsWithResultResponseBodyJobs) *QueryJobsWithResultResponseBody
	GetJobs() *QueryJobsWithResultResponseBodyJobs
	SetLabels(v []*QueryJobsWithResultResponseBodyLabels) *QueryJobsWithResultResponseBody
	GetLabels() []*QueryJobsWithResultResponseBodyLabels
	SetMessage(v string) *QueryJobsWithResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryJobsWithResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryJobsWithResultResponseBody
	GetSuccess() *bool
	SetVariableNames(v []*string) *QueryJobsWithResultResponseBody
	GetVariableNames() []*string
}

type QueryJobsWithResultResponseBody struct {
	// The response code. The value `OK` indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The paginated list of jobs.
	Jobs *QueryJobsWithResultResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Struct"`
	// The tag information that can be used as filter criteria.
	//
	// > Displays information for all tags in the job group that have enumerated values.
	Labels []*QueryJobsWithResultResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9bdaa1d1-a036-4451-ab11-ca0373679091
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The names of tags to be collected.
	VariableNames []*string `json:"VariableNames,omitempty" xml:"VariableNames,omitempty" type:"Repeated"`
}

func (s QueryJobsWithResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryJobsWithResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryJobsWithResultResponseBody) GetJobs() *QueryJobsWithResultResponseBodyJobs {
	return s.Jobs
}

func (s *QueryJobsWithResultResponseBody) GetLabels() []*QueryJobsWithResultResponseBodyLabels {
	return s.Labels
}

func (s *QueryJobsWithResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryJobsWithResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryJobsWithResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryJobsWithResultResponseBody) GetVariableNames() []*string {
	return s.VariableNames
}

func (s *QueryJobsWithResultResponseBody) SetCode(v string) *QueryJobsWithResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetHttpStatusCode(v int32) *QueryJobsWithResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetJobs(v *QueryJobsWithResultResponseBodyJobs) *QueryJobsWithResultResponseBody {
	s.Jobs = v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetLabels(v []*QueryJobsWithResultResponseBodyLabels) *QueryJobsWithResultResponseBody {
	s.Labels = v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetMessage(v string) *QueryJobsWithResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetRequestId(v string) *QueryJobsWithResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetSuccess(v bool) *QueryJobsWithResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryJobsWithResultResponseBody) SetVariableNames(v []*string) *QueryJobsWithResultResponseBody {
	s.VariableNames = v
	return s
}

func (s *QueryJobsWithResultResponseBody) Validate() error {
	if s.Jobs != nil {
		if err := s.Jobs.Validate(); err != nil {
			return err
		}
	}
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

type QueryJobsWithResultResponseBodyJobs struct {
	// A list of jobs.
	List []*QueryJobsWithResultResponseBodyJobsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
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
	// The total number of entries.
	//
	// example:
	//
	// 1
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobs) GetList() []*QueryJobsWithResultResponseBodyJobsList {
	return s.List
}

func (s *QueryJobsWithResultResponseBodyJobs) GetPageCount() *int32 {
	return s.PageCount
}

func (s *QueryJobsWithResultResponseBodyJobs) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryJobsWithResultResponseBodyJobs) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryJobsWithResultResponseBodyJobs) GetRowCount() *int32 {
	return s.RowCount
}

func (s *QueryJobsWithResultResponseBodyJobs) SetList(v []*QueryJobsWithResultResponseBodyJobsList) *QueryJobsWithResultResponseBodyJobs {
	s.List = v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobs) SetPageCount(v int32) *QueryJobsWithResultResponseBodyJobs {
	s.PageCount = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobs) SetPageNumber(v int32) *QueryJobsWithResultResponseBodyJobs {
	s.PageNumber = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobs) SetPageSize(v int32) *QueryJobsWithResultResponseBodyJobs {
	s.PageSize = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobs) SetRowCount(v int32) *QueryJobsWithResultResponseBodyJobs {
	s.RowCount = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobs) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobsWithResultResponseBodyJobsList struct {
	// The job ID.
	//
	// example:
	//
	// cc231a1d-3c05-4739-8926-193ecf4097ba
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The reason the job failed.
	//
	// example:
	//
	// - Unknown（未知错误）
	//
	// - NoAnswer（无人接听）
	//
	// - InvalidStrategy（无效的策略，策略配置不正确）
	//
	// - TimeUp（调度时发现超时）
	//
	// - NoStrategy（策略为空或没有找到）
	//
	// - CallFailed（呼叫失败）
	//
	// -PerDayCallCountLimit（号码每日呼叫次数限制）
	//
	// - ContactBlockList（禁止外呼名单）
	//
	// - EmptyNumber（空号不再外呼）
	//
	// - JobPerDayCallCountLimit（号码每日呼叫次数限制）
	//
	// - VerificationCancelled（呼叫前验证不通过取消）
	//
	// - ContactSuspended（止呼）
	//
	// - InArrears（欠费）
	//
	// - OutOfService（停机）
	//
	// - NoneRepeatableJobMaxAttemptCountLimit（ 任务最大尝试次数， 当RepeatBy（@see ）为None时生效）
	JobFailureReason *string `json:"JobFailureReason,omitempty" xml:"JobFailureReason,omitempty"`
	// The most recent call.
	LatestTask *QueryJobsWithResultResponseBodyJobsListLatestTask `json:"LatestTask,omitempty" xml:"LatestTask,omitempty" type:"Struct"`
	// The job status. Valid values:
	//
	// - `Scheduling`: The job is being scheduled.
	//
	// - `Executing`: The job is in progress.
	//
	// - `Succeeded`: The job succeeded, and the contact was reached.
	//
	// - `Paused`: The job is paused.
	//
	// - `Failed`: The job failed because the contact was not reached.
	//
	// - `Cancelled`: The job was cancelled by a user.
	//
	// - `Drafted`: The job is a draft.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the job status.
	//
	// example:
	//
	// 结束-已触达
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobsList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobsList) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobsList) GetId() *string {
	return s.Id
}

func (s *QueryJobsWithResultResponseBodyJobsList) GetJobFailureReason() *string {
	return s.JobFailureReason
}

func (s *QueryJobsWithResultResponseBodyJobsList) GetLatestTask() *QueryJobsWithResultResponseBodyJobsListLatestTask {
	return s.LatestTask
}

func (s *QueryJobsWithResultResponseBodyJobsList) GetStatus() *string {
	return s.Status
}

func (s *QueryJobsWithResultResponseBodyJobsList) GetStatusName() *string {
	return s.StatusName
}

func (s *QueryJobsWithResultResponseBodyJobsList) SetId(v string) *QueryJobsWithResultResponseBodyJobsList {
	s.Id = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsList) SetJobFailureReason(v string) *QueryJobsWithResultResponseBodyJobsList {
	s.JobFailureReason = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsList) SetLatestTask(v *QueryJobsWithResultResponseBodyJobsListLatestTask) *QueryJobsWithResultResponseBodyJobsList {
	s.LatestTask = v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsList) SetStatus(v string) *QueryJobsWithResultResponseBodyJobsList {
	s.Status = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsList) SetStatusName(v string) *QueryJobsWithResultResponseBodyJobsList {
	s.StatusName = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsList) Validate() error {
	if s.LatestTask != nil {
		if err := s.LatestTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobsWithResultResponseBodyJobsListLatestTask struct {
	// The call duration, in milliseconds.
	//
	// example:
	//
	// 40000
	CallDuration *int32 `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	// The call duration, formatted for display.
	//
	// example:
	//
	// 40
	CallDurationDisplay *string `json:"CallDurationDisplay,omitempty" xml:"CallDurationDisplay,omitempty"`
	// The time when the call was made. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1615363297000
	CallTime *int64 `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// The contact.
	Contact *QueryJobsWithResultResponseBodyJobsListLatestTaskContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The error codes for the call.
	DialExceptionCodes []*QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes `json:"DialExceptionCodes,omitempty" xml:"DialExceptionCodes,omitempty" type:"Repeated"`
	// The custom tags (key-value pairs) associated with the call.
	Extras []*QueryJobsWithResultResponseBodyJobsListLatestTaskExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Repeated"`
	// example:
	//
	// 被叫用户
	HangUpDirection *string `json:"HangUpDirection,omitempty" xml:"HangUpDirection,omitempty"`
	// Indicates whether the call was answered.
	//
	// example:
	//
	// true
	HasAnswered *bool `json:"HasAnswered,omitempty" xml:"HasAnswered,omitempty"`
	// Indicates whether the call was hung up due to rejection.
	//
	// example:
	//
	// false
	HasHangUpByRejection *bool `json:"HasHangUpByRejection,omitempty" xml:"HasHangUpByRejection,omitempty"`
	// Indicates whether the last voice playback was completed before the call was hung up.
	//
	// example:
	//
	// true
	HasLastPlaybackCompleted *bool `json:"HasLastPlaybackCompleted,omitempty" xml:"HasLastPlaybackCompleted,omitempty"`
	// Indicates whether the call flow was completed.
	//
	// example:
	//
	// true
	HasReachedEndOfFlow *bool `json:"HasReachedEndOfFlow,omitempty" xml:"HasReachedEndOfFlow,omitempty"`
	// The call status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the call status.
	//
	// example:
	//
	// 未呼出-超出每日上限
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// Information about tags that were matched. This parameter is returned for flows created using the legacy canvas.
	TagHits []*QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits `json:"TagHits,omitempty" xml:"TagHits,omitempty" type:"Repeated"`
	// The call end reason.
	//
	// example:
	//
	// FINISHED
	TaskEndReason *string `json:"TaskEndReason,omitempty" xml:"TaskEndReason,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTask) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTask) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetCallDuration() *int32 {
	return s.CallDuration
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetCallDurationDisplay() *string {
	return s.CallDurationDisplay
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetCallTime() *int64 {
	return s.CallTime
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetContact() *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	return s.Contact
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetDialExceptionCodes() []*QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes {
	return s.DialExceptionCodes
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetExtras() []*QueryJobsWithResultResponseBodyJobsListLatestTaskExtras {
	return s.Extras
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetHangUpDirection() *string {
	return s.HangUpDirection
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetHasAnswered() *bool {
	return s.HasAnswered
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetHasHangUpByRejection() *bool {
	return s.HasHangUpByRejection
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetHasLastPlaybackCompleted() *bool {
	return s.HasLastPlaybackCompleted
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetHasReachedEndOfFlow() *bool {
	return s.HasReachedEndOfFlow
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetStatus() *string {
	return s.Status
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetStatusName() *string {
	return s.StatusName
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetTagHits() []*QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits {
	return s.TagHits
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) GetTaskEndReason() *string {
	return s.TaskEndReason
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetCallDuration(v int32) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.CallDuration = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetCallDurationDisplay(v string) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.CallDurationDisplay = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetCallTime(v int64) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.CallTime = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetContact(v *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.Contact = v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetDialExceptionCodes(v []*QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.DialExceptionCodes = v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetExtras(v []*QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.Extras = v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetHangUpDirection(v string) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.HangUpDirection = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetHasAnswered(v bool) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.HasAnswered = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetHasHangUpByRejection(v bool) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.HasHangUpByRejection = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetHasLastPlaybackCompleted(v bool) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.HasLastPlaybackCompleted = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetHasReachedEndOfFlow(v bool) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.HasReachedEndOfFlow = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetStatus(v string) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.Status = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetStatusName(v string) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.StatusName = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetTagHits(v []*QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.TagHits = v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) SetTaskEndReason(v string) *QueryJobsWithResultResponseBodyJobsListLatestTask {
	s.TaskEndReason = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTask) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	if s.DialExceptionCodes != nil {
		for _, item := range s.DialExceptionCodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Extras != nil {
		for _, item := range s.Extras {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagHits != nil {
		for _, item := range s.TagHits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobsWithResultResponseBodyJobsListLatestTaskContact struct {
	// The honorific for the contact, such as Mr. or Ms. If not provided, it defaults to the value of the `Name` parameter.
	//
	// > This is a custom parameter. You can specify it in the JSON object when you start an outbound job.
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// The unique ID of the contact, generated by the system when the contact list is uploaded.
	//
	// example:
	//
	// 63860deb-6218-45df-b1e0-76f2b166e790
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The job ID.
	//
	// example:
	//
	// cc231a1d-3c05-4739-8926-193ecf4097ba
	JobUuid *string `json:"JobUuid,omitempty" xml:"JobUuid,omitempty"`
	// The name of the contact. For example, John Smith.
	//
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 1882020****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The phone number specified by the callee during the conversation. If provided, Outbound Bot calls this number.
	//
	// > This is a custom parameter. You can specify it in the JSON object when you start an outbound job.
	//
	// example:
	//
	// 134123****
	PreferredPhoneNumber *string `json:"PreferredPhoneNumber,omitempty" xml:"PreferredPhoneNumber,omitempty"`
	// Your custom ID for the contact. This ID helps you uniquely identify the contact and prevent issues caused by duplicate names.
	//
	// example:
	//
	// C01
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The role of the contact in the outbound calling scenario. This parameter is optional. For example, in a debt collection scenario, roles can include `borrower`, `co-borrower`, or `guarantor`.
	//
	// > This is a custom parameter. You can specify it in the JSON object when starting a job.
	//
	// example:
	//
	// 借款人
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The call attempt number.
	//
	// example:
	//
	// 151
	Round *int32 `json:"Round,omitempty" xml:"Round,omitempty"`
	// The status of the contact. For example, `Available` (the contact is available), `WrongNumber` (the phone number is incorrect), or `DoesNotExist` (the phone number does not exist).
	//
	// > This is a custom parameter. You can specify it in the JSON object when you start an outbound job.
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskContact) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetHonorific() *string {
	return s.Honorific
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetId() *string {
	return s.Id
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetJobUuid() *string {
	return s.JobUuid
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetName() *string {
	return s.Name
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetPreferredPhoneNumber() *string {
	return s.PreferredPhoneNumber
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetRole() *string {
	return s.Role
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetRound() *int32 {
	return s.Round
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) GetState() *string {
	return s.State
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetHonorific(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.Honorific = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetId(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.Id = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetJobUuid(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.JobUuid = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetName(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.Name = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetPhoneNumber(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.PhoneNumber = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetPreferredPhoneNumber(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.PreferredPhoneNumber = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetReferenceId(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.ReferenceId = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetRole(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.Role = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetRound(v int32) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.Round = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) SetState(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskContact {
	s.State = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskContact) Validate() error {
	return dara.Validate(s)
}

type QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes struct {
	// The error code.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message
	//
	// example:
	//
	// 无
	Hint *string `json:"Hint,omitempty" xml:"Hint,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) GetCode() *string {
	return s.Code
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) GetHint() *string {
	return s.Hint
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) SetCode(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes {
	s.Code = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) SetHint(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes {
	s.Hint = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskDialExceptionCodes) Validate() error {
	return dara.Validate(s)
}

type QueryJobsWithResultResponseBodyJobsListLatestTaskExtras struct {
	// The tag key.
	//
	// example:
	//
	// 是否已经接通
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 是
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) GetKey() *string {
	return s.Key
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) GetValue() *string {
	return s.Value
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) SetKey(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras {
	s.Key = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) SetValue(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras {
	s.Value = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskExtras) Validate() error {
	return dara.Validate(s)
}

type QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits struct {
	// The tag group to which the tag belongs.
	//
	// example:
	//
	// 动物
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 猫幼年期,猫幼年期
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) GetTagGroup() *string {
	return s.TagGroup
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) GetTagName() *string {
	return s.TagName
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) SetTagGroup(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits {
	s.TagGroup = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) SetTagName(v string) *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits {
	s.TagName = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyJobsListLatestTaskTagHits) Validate() error {
	return dara.Validate(s)
}

type QueryJobsWithResultResponseBodyLabels struct {
	// Tag Name
	//
	// example:
	//
	// 是否满意
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Tag value list.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s QueryJobsWithResultResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponseBodyLabels) GetName() *string {
	return s.Name
}

func (s *QueryJobsWithResultResponseBodyLabels) GetValueList() []*string {
	return s.ValueList
}

func (s *QueryJobsWithResultResponseBodyLabels) SetName(v string) *QueryJobsWithResultResponseBodyLabels {
	s.Name = &v
	return s
}

func (s *QueryJobsWithResultResponseBodyLabels) SetValueList(v []*string) *QueryJobsWithResultResponseBodyLabels {
	s.ValueList = v
	return s
}

func (s *QueryJobsWithResultResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}

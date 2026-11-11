// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *GetConnectionTicketResponseBody
	GetAppInstanceId() *string
	SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody
	GetAppInstancePersistentId() *string
	SetAvatarId(v string) *GetConnectionTicketResponseBody
	GetAvatarId() *string
	SetBindQueueInfo(v *GetConnectionTicketResponseBodyBindQueueInfo) *GetConnectionTicketResponseBody
	GetBindQueueInfo() *GetConnectionTicketResponseBodyBindQueueInfo
	SetCode(v string) *GetConnectionTicketResponseBody
	GetCode() *string
	SetLoginToken(v string) *GetConnectionTicketResponseBody
	GetLoginToken() *string
	SetMessage(v string) *GetConnectionTicketResponseBody
	GetMessage() *string
	SetNextPollIntervalMs(v int32) *GetConnectionTicketResponseBody
	GetNextPollIntervalMs() *int32
	SetOsType(v string) *GetConnectionTicketResponseBody
	GetOsType() *string
	SetPolicy(v *GetConnectionTicketResponseBodyPolicy) *GetConnectionTicketResponseBody
	GetPolicy() *GetConnectionTicketResponseBodyPolicy
	SetRegionId(v string) *GetConnectionTicketResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetConnectionTicketResponseBody
	GetRequestId() *string
	SetRetryTimes(v int32) *GetConnectionTicketResponseBody
	GetRetryTimes() *int32
	SetTaskId(v string) *GetConnectionTicketResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *GetConnectionTicketResponseBody
	GetTaskStatus() *string
	SetTenantId(v int64) *GetConnectionTicketResponseBody
	GetTenantId() *int64
	SetTicket(v string) *GetConnectionTicketResponseBody
	GetTicket() *string
}

type GetConnectionTicketResponseBody struct {
	// The delivery group ID.
	//
	// example:
	//
	// aig-53fvrq1oa******
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The application instance ID.
	//
	// example:
	//
	// ai-gc1gemx6vp******
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The persistent session ID.
	//
	// example:
	//
	// p-0cdr5c9kl******
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The user avatar ID.
	//
	// example:
	//
	// abc
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The queue information.
	//
	// > This parameter is not available.
	BindQueueInfo *GetConnectionTicketResponseBodyBindQueueInfo `json:"BindQueueInfo,omitempty" xml:"BindQueueInfo,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InternalError.TicketGenInternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logon token.
	//
	// example:
	//
	// v2c4e2ef03d62******
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The error message.
	//
	// example:
	//
	// reenter app instance failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The next polling interval, in milliseconds. If this parameter is empty, the value is 1000 ms.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 500
	NextPollIntervalMs *int32 `json:"NextPollIntervalMs,omitempty" xml:"NextPollIntervalMs,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The session policy.
	Policy *GetConnectionTicketResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the region to which the delivery group belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 3
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f3d1b31c-605e-******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Running
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The tenant ID, which is the Alibaba Cloud account UID.
	//
	// example:
	//
	// 1126819517******
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The connection ticket.
	//
	// example:
	//
	// n7n9bqZlPrvgUOPY0tOfVYGBBcdoPoH3******
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetConnectionTicketResponseBody) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *GetConnectionTicketResponseBody) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *GetConnectionTicketResponseBody) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetConnectionTicketResponseBody) GetBindQueueInfo() *GetConnectionTicketResponseBodyBindQueueInfo {
	return s.BindQueueInfo
}

func (s *GetConnectionTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConnectionTicketResponseBody) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GetConnectionTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConnectionTicketResponseBody) GetNextPollIntervalMs() *int32 {
	return s.NextPollIntervalMs
}

func (s *GetConnectionTicketResponseBody) GetOsType() *string {
	return s.OsType
}

func (s *GetConnectionTicketResponseBody) GetPolicy() *GetConnectionTicketResponseBodyPolicy {
	return s.Policy
}

func (s *GetConnectionTicketResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConnectionTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectionTicketResponseBody) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *GetConnectionTicketResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetConnectionTicketResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetConnectionTicketResponseBody) GetTicket() *string {
	return s.Ticket
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAvatarId(v string) *GetConnectionTicketResponseBody {
	s.AvatarId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetBindQueueInfo(v *GetConnectionTicketResponseBodyBindQueueInfo) *GetConnectionTicketResponseBody {
	s.BindQueueInfo = v
	return s
}

func (s *GetConnectionTicketResponseBody) SetCode(v string) *GetConnectionTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetLoginToken(v string) *GetConnectionTicketResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetMessage(v string) *GetConnectionTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetNextPollIntervalMs(v int32) *GetConnectionTicketResponseBody {
	s.NextPollIntervalMs = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetOsType(v string) *GetConnectionTicketResponseBody {
	s.OsType = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetPolicy(v *GetConnectionTicketResponseBodyPolicy) *GetConnectionTicketResponseBody {
	s.Policy = v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRegionId(v string) *GetConnectionTicketResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRetryTimes(v int32) *GetConnectionTicketResponseBody {
	s.RetryTimes = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTenantId(v int64) *GetConnectionTicketResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetConnectionTicketResponseBody) Validate() error {
	if s.BindQueueInfo != nil {
		if err := s.BindQueueInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConnectionTicketResponseBodyBindQueueInfo struct {
	// The queue task status.
	//
	// example:
	//
	// Waiting
	QueueStatus *string `json:"QueueStatus,omitempty" xml:"QueueStatus,omitempty"`
	// The ranking information.
	//
	// example:
	//
	// 1
	Rank *int32 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The timeout period for manual connection after the turn is reached.
	//
	// example:
	//
	// 120
	ReadyTimeout *int64 `json:"ReadyTimeout,omitempty" xml:"ReadyTimeout,omitempty"`
	// The remaining wait time. Unit: seconds.
	//
	// example:
	//
	// 100
	RemainingTime *int64 `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty"`
	// The queue request key.
	//
	// example:
	//
	// alice******
	RequestKey *string `json:"RequestKey,omitempty" xml:"RequestKey,omitempty"`
	// The queue ID.
	//
	// example:
	//
	// aig-53fvrq1oa******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The time already waited in the queue. Unit: seconds.
	//
	// example:
	//
	// 100
	WaitTime *int64 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s GetConnectionTicketResponseBodyBindQueueInfo) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketResponseBodyBindQueueInfo) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetQueueStatus() *string {
	return s.QueueStatus
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetRank() *int32 {
	return s.Rank
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetReadyTimeout() *int64 {
	return s.ReadyTimeout
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetRemainingTime() *int64 {
	return s.RemainingTime
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetRequestKey() *string {
	return s.RequestKey
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetTargetId() *string {
	return s.TargetId
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) GetWaitTime() *int64 {
	return s.WaitTime
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetQueueStatus(v string) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.QueueStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetRank(v int32) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.Rank = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetReadyTimeout(v int64) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.ReadyTimeout = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetRemainingTime(v int64) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.RemainingTime = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetRequestKey(v string) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.RequestKey = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetTargetId(v string) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.TargetId = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetWaitTime(v int64) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.WaitTime = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) Validate() error {
	return dara.Validate(s)
}

type GetConnectionTicketResponseBodyPolicy struct {
	// Specifies whether to use adaptive resolution.
	//
	// >
	//
	// > - `true`: The session resolution changes with the terminal display area. In this case, `SessionResolutionWidth` and `SessionResolutionHeight` are the maximum values of the resolution change.
	//
	// > - `false`: The session resolution does not change with the terminal display area. In this case, the resolution is fixed to the values of `SessionResolutionWidth` and `SessionResolutionHeight`.
	//
	// example:
	//
	// true
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" xml:"ResolutionAdaptive,omitempty"`
	// The resolution height, in pixels.
	//
	// example:
	//
	// 1080
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution width, in pixels.
	//
	// example:
	//
	// 1920
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s GetConnectionTicketResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBodyPolicy) GetResolutionAdaptive() *string {
	return s.ResolutionAdaptive
}

func (s *GetConnectionTicketResponseBodyPolicy) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *GetConnectionTicketResponseBodyPolicy) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionAdaptive(v string) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionAdaptive = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionHeight(v int32) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionHeight = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionWidth(v int32) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionWidth = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}

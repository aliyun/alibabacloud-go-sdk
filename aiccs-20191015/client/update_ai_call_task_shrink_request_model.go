// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiCallTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallDayShrink(v string) *UpdateAiCallTaskShrinkRequest
	GetCallDayShrink() *string
	SetCallRetryInterval(v int64) *UpdateAiCallTaskShrinkRequest
	GetCallRetryInterval() *int64
	SetCallRetryReasonShrink(v string) *UpdateAiCallTaskShrinkRequest
	GetCallRetryReasonShrink() *string
	SetCallRetryTimes(v int64) *UpdateAiCallTaskShrinkRequest
	GetCallRetryTimes() *int64
	SetCallTimeShrink(v string) *UpdateAiCallTaskShrinkRequest
	GetCallTimeShrink() *string
	SetLineEncoding(v string) *UpdateAiCallTaskShrinkRequest
	GetLineEncoding() *string
	SetLinePhoneNum(v string) *UpdateAiCallTaskShrinkRequest
	GetLinePhoneNum() *string
	SetMissCallRetry(v bool) *UpdateAiCallTaskShrinkRequest
	GetMissCallRetry() *bool
	SetOwnerId(v int64) *UpdateAiCallTaskShrinkRequest
	GetOwnerId() *int64
	SetPhoneType(v int64) *UpdateAiCallTaskShrinkRequest
	GetPhoneType() *int64
	SetResourceOwnerAccount(v string) *UpdateAiCallTaskShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAiCallTaskShrinkRequest
	GetResourceOwnerId() *int64
	SetSource(v int64) *UpdateAiCallTaskShrinkRequest
	GetSource() *int64
	SetStartType(v string) *UpdateAiCallTaskShrinkRequest
	GetStartType() *string
	SetTaskCps(v int64) *UpdateAiCallTaskShrinkRequest
	GetTaskCps() *int64
	SetTaskId(v string) *UpdateAiCallTaskShrinkRequest
	GetTaskId() *string
	SetTaskName(v string) *UpdateAiCallTaskShrinkRequest
	GetTaskName() *string
	SetTaskStartTime(v int64) *UpdateAiCallTaskShrinkRequest
	GetTaskStartTime() *int64
	SetVirtualNumber(v string) *UpdateAiCallTaskShrinkRequest
	GetVirtualNumber() *string
}

type UpdateAiCallTaskShrinkRequest struct {
	// The days of the week when calls can be made.
	//
	// This parameter is required.
	CallDayShrink *string `json:"CallDay,omitempty" xml:"CallDay,omitempty"`
	// The retry interval in minutes. The maximum value is 120.
	//
	// example:
	//
	// 25
	CallRetryInterval *int64 `json:"CallRetryInterval,omitempty" xml:"CallRetryInterval,omitempty"`
	// The call failure statuses that trigger a retry.
	CallRetryReasonShrink *string `json:"CallRetryReason,omitempty" xml:"CallRetryReason,omitempty"`
	// The number of retries. The maximum value is 3.
	//
	// example:
	//
	// 2
	CallRetryTimes *int64 `json:"CallRetryTimes,omitempty" xml:"CallRetryTimes,omitempty"`
	// The callable time windows.
	//
	// This parameter is required.
	CallTimeShrink *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	LineEncoding *string `json:"LineEncoding,omitempty" xml:"LineEncoding,omitempty"`
	// example:
	//
	// 示例值示例值
	LinePhoneNum *string `json:"LinePhoneNum,omitempty" xml:"LinePhoneNum,omitempty"`
	// Specifies whether to enable retry. Valid values:
	//
	// - `true`: Yes.
	//
	// - `false` (default): No.
	//
	// example:
	//
	// true
	MissCallRetry *bool  `json:"MissCallRetry,omitempty" xml:"MissCallRetry,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 53
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 31
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The startup method. Valid values:
	//
	// - `IMMEDIATE`: Start immediately.
	//
	// - `SCHEDULE`: Start at a specified time.
	//
	// This parameter is required.
	//
	// example:
	//
	// IMMEDIATE
	StartType *string `json:"StartType,omitempty" xml:"StartType,omitempty"`
	// The number of concurrent calls per second (CPS) for the task. The maximum value is 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 57
	TaskCps *int64 `json:"TaskCps,omitempty" xml:"TaskCps,omitempty"`
	// The ID of the task to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1187**************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name. The name must be unique within an Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试任务
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The scheduled start time for the task, specified as a Unix timestamp in milliseconds. This parameter is required when `StartType` is set to `SCHEDULE`.
	//
	// example:
	//
	// 1748923429000
	TaskStartTime *int64 `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// The service instance used for outbound calls.
	//
	// example:
	//
	// 0537022*****
	VirtualNumber *string `json:"VirtualNumber,omitempty" xml:"VirtualNumber,omitempty"`
}

func (s UpdateAiCallTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiCallTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallDayShrink() *string {
	return s.CallDayShrink
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallRetryInterval() *int64 {
	return s.CallRetryInterval
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallRetryReasonShrink() *string {
	return s.CallRetryReasonShrink
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallRetryTimes() *int64 {
	return s.CallRetryTimes
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallTimeShrink() *string {
	return s.CallTimeShrink
}

func (s *UpdateAiCallTaskShrinkRequest) GetLineEncoding() *string {
	return s.LineEncoding
}

func (s *UpdateAiCallTaskShrinkRequest) GetLinePhoneNum() *string {
	return s.LinePhoneNum
}

func (s *UpdateAiCallTaskShrinkRequest) GetMissCallRetry() *bool {
	return s.MissCallRetry
}

func (s *UpdateAiCallTaskShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAiCallTaskShrinkRequest) GetPhoneType() *int64 {
	return s.PhoneType
}

func (s *UpdateAiCallTaskShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAiCallTaskShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAiCallTaskShrinkRequest) GetSource() *int64 {
	return s.Source
}

func (s *UpdateAiCallTaskShrinkRequest) GetStartType() *string {
	return s.StartType
}

func (s *UpdateAiCallTaskShrinkRequest) GetTaskCps() *int64 {
	return s.TaskCps
}

func (s *UpdateAiCallTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAiCallTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateAiCallTaskShrinkRequest) GetTaskStartTime() *int64 {
	return s.TaskStartTime
}

func (s *UpdateAiCallTaskShrinkRequest) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallDayShrink(v string) *UpdateAiCallTaskShrinkRequest {
	s.CallDayShrink = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallRetryInterval(v int64) *UpdateAiCallTaskShrinkRequest {
	s.CallRetryInterval = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallRetryReasonShrink(v string) *UpdateAiCallTaskShrinkRequest {
	s.CallRetryReasonShrink = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallRetryTimes(v int64) *UpdateAiCallTaskShrinkRequest {
	s.CallRetryTimes = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallTimeShrink(v string) *UpdateAiCallTaskShrinkRequest {
	s.CallTimeShrink = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetLineEncoding(v string) *UpdateAiCallTaskShrinkRequest {
	s.LineEncoding = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetLinePhoneNum(v string) *UpdateAiCallTaskShrinkRequest {
	s.LinePhoneNum = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetMissCallRetry(v bool) *UpdateAiCallTaskShrinkRequest {
	s.MissCallRetry = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetOwnerId(v int64) *UpdateAiCallTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetPhoneType(v int64) *UpdateAiCallTaskShrinkRequest {
	s.PhoneType = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAiCallTaskShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetResourceOwnerId(v int64) *UpdateAiCallTaskShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetSource(v int64) *UpdateAiCallTaskShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetStartType(v string) *UpdateAiCallTaskShrinkRequest {
	s.StartType = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetTaskCps(v int64) *UpdateAiCallTaskShrinkRequest {
	s.TaskCps = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetTaskId(v string) *UpdateAiCallTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetTaskName(v string) *UpdateAiCallTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetTaskStartTime(v int64) *UpdateAiCallTaskShrinkRequest {
	s.TaskStartTime = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetVirtualNumber(v string) *UpdateAiCallTaskShrinkRequest {
	s.VirtualNumber = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

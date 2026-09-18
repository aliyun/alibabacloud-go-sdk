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
	SetCallExpireDate(v string) *UpdateAiCallTaskShrinkRequest
	GetCallExpireDate() *string
	SetCallExpireMinutes(v int64) *UpdateAiCallTaskShrinkRequest
	GetCallExpireMinutes() *int64
	SetCallExpireType(v int64) *UpdateAiCallTaskShrinkRequest
	GetCallExpireType() *int64
	SetCallRetryInterval(v int64) *UpdateAiCallTaskShrinkRequest
	GetCallRetryInterval() *int64
	SetCallRetryReasonShrink(v string) *UpdateAiCallTaskShrinkRequest
	GetCallRetryReasonShrink() *string
	SetCallRetryTimes(v int64) *UpdateAiCallTaskShrinkRequest
	GetCallRetryTimes() *int64
	SetCallTimeShrink(v string) *UpdateAiCallTaskShrinkRequest
	GetCallTimeShrink() *string
	SetCallableTimeShrink(v string) *UpdateAiCallTaskShrinkRequest
	GetCallableTimeShrink() *string
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
	// The available call days.
	//
	// This parameter is required.
	CallDayShrink *string `json:"CallDay,omitempty" xml:"CallDay,omitempty"`
	// The expiration date of outbound call details (the specific deadline).
	//
	// example:
	//
	// 2026-07-30 20:00:20
	CallExpireDate *string `json:"CallExpireDate,omitempty" xml:"CallExpireDate,omitempty"`
	// The expiration duration of outbound call details. Unit: minutes.
	//
	// example:
	//
	// 10
	CallExpireMinutes *int64 `json:"CallExpireMinutes,omitempty" xml:"CallExpireMinutes,omitempty"`
	// The outbound call validity type. Valid values:
	//
	// 0: permanently valid.
	//
	// 1: valid for a specified duration after import.
	//
	// 2: valid until a specified date.
	//
	// example:
	//
	// 0
	CallExpireType *int64 `json:"CallExpireType,omitempty" xml:"CallExpireType,omitempty"`
	// The retry interval. Unit: minutes. The maximum value is 120 minutes.
	//
	// example:
	//
	// 25
	CallRetryInterval *int64 `json:"CallRetryInterval,omitempty" xml:"CallRetryInterval,omitempty"`
	// The reasons for retry upon failure.
	CallRetryReasonShrink *string `json:"CallRetryReason,omitempty" xml:"CallRetryReason,omitempty"`
	// The number of retries. The maximum value is 3.
	//
	// example:
	//
	// 2
	CallRetryTimes *int64 `json:"CallRetryTimes,omitempty" xml:"CallRetryTimes,omitempty"`
	// The available call time periods.
	//
	// This parameter is required.
	CallTimeShrink     *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	CallableTimeShrink *string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	// The line encoding.
	//
	// example:
	//
	// JILIANG_***_***_NET
	LineEncoding *string `json:"LineEncoding,omitempty" xml:"LineEncoding,omitempty"`
	// The customer-provided line number.
	//
	// example:
	//
	// 152****3120
	LinePhoneNum *string `json:"LinePhoneNum,omitempty" xml:"LinePhoneNum,omitempty"`
	// Specifies whether to enable retry. Valid values:
	//
	// - true: Enabled.
	//
	// - false (default): Disabled.
	//
	// example:
	//
	// true
	MissCallRetry *bool  `json:"MissCallRetry,omitempty" xml:"MissCallRetry,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number type. This parameter is used when the creation source is engine-based.
	//
	// 0: Alibaba Cloud number.
	//
	// 1: Customer-provided line.
	//
	// example:
	//
	// 0
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The creation source. Valid values:
	//
	// 0: created by agent.
	//
	// 1: created by engine.
	//
	// example:
	//
	// Cannot be modified. Leave this parameter empty
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start mode. Valid values:
	//
	// - IMMEDIATE: starts immediately.
	//
	// - SCHEDULE: starts at a scheduled time.
	//
	// This parameter is required.
	//
	// example:
	//
	// IMMEDIATE
	StartType *string `json:"StartType,omitempty" xml:"StartType,omitempty"`
	// The task concurrency. The maximum value is 500.
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
	// The task name. The name must be unique within the same account.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestTask
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The preset start time of the task. The value is a UNIX timestamp in milliseconds. This parameter is valid and required when the StartType parameter is set to SCHEDULE. The task automatically starts at the time specified by this parameter.
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

func (s *UpdateAiCallTaskShrinkRequest) GetCallExpireDate() *string {
	return s.CallExpireDate
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallExpireMinutes() *int64 {
	return s.CallExpireMinutes
}

func (s *UpdateAiCallTaskShrinkRequest) GetCallExpireType() *int64 {
	return s.CallExpireType
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

func (s *UpdateAiCallTaskShrinkRequest) GetCallableTimeShrink() *string {
	return s.CallableTimeShrink
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

func (s *UpdateAiCallTaskShrinkRequest) SetCallExpireDate(v string) *UpdateAiCallTaskShrinkRequest {
	s.CallExpireDate = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallExpireMinutes(v int64) *UpdateAiCallTaskShrinkRequest {
	s.CallExpireMinutes = &v
	return s
}

func (s *UpdateAiCallTaskShrinkRequest) SetCallExpireType(v int64) *UpdateAiCallTaskShrinkRequest {
	s.CallExpireType = &v
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

func (s *UpdateAiCallTaskShrinkRequest) SetCallableTimeShrink(v string) *UpdateAiCallTaskShrinkRequest {
	s.CallableTimeShrink = &v
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

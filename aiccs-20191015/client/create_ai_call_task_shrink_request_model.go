// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiCallTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateAiCallTaskShrinkRequest
	GetAgentId() *string
	SetApplicationCode(v string) *CreateAiCallTaskShrinkRequest
	GetApplicationCode() *string
	SetCallDayShrink(v string) *CreateAiCallTaskShrinkRequest
	GetCallDayShrink() *string
	SetCallExpireDate(v string) *CreateAiCallTaskShrinkRequest
	GetCallExpireDate() *string
	SetCallExpireMinutes(v int64) *CreateAiCallTaskShrinkRequest
	GetCallExpireMinutes() *int64
	SetCallExpireType(v int64) *CreateAiCallTaskShrinkRequest
	GetCallExpireType() *int64
	SetCallRetryInterval(v int64) *CreateAiCallTaskShrinkRequest
	GetCallRetryInterval() *int64
	SetCallRetryReasonShrink(v string) *CreateAiCallTaskShrinkRequest
	GetCallRetryReasonShrink() *string
	SetCallRetryTimes(v int64) *CreateAiCallTaskShrinkRequest
	GetCallRetryTimes() *int64
	SetCallTimeShrink(v string) *CreateAiCallTaskShrinkRequest
	GetCallTimeShrink() *string
	SetCallableTimeShrink(v string) *CreateAiCallTaskShrinkRequest
	GetCallableTimeShrink() *string
	SetLineEncoding(v string) *CreateAiCallTaskShrinkRequest
	GetLineEncoding() *string
	SetLinePhoneNum(v string) *CreateAiCallTaskShrinkRequest
	GetLinePhoneNum() *string
	SetMissCallRetry(v bool) *CreateAiCallTaskShrinkRequest
	GetMissCallRetry() *bool
	SetOwnerId(v int64) *CreateAiCallTaskShrinkRequest
	GetOwnerId() *int64
	SetPhoneType(v int64) *CreateAiCallTaskShrinkRequest
	GetPhoneType() *int64
	SetResourceOwnerAccount(v string) *CreateAiCallTaskShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAiCallTaskShrinkRequest
	GetResourceOwnerId() *int64
	SetSource(v int64) *CreateAiCallTaskShrinkRequest
	GetSource() *int64
	SetStartType(v string) *CreateAiCallTaskShrinkRequest
	GetStartType() *string
	SetTaskCps(v int64) *CreateAiCallTaskShrinkRequest
	GetTaskCps() *int64
	SetTaskName(v string) *CreateAiCallTaskShrinkRequest
	GetTaskName() *string
	SetTaskStartTime(v int64) *CreateAiCallTaskShrinkRequest
	GetTaskStartTime() *int64
	SetVirtualNumber(v string) *CreateAiCallTaskShrinkRequest
	GetVirtualNumber() *string
}

type CreateAiCallTaskShrinkRequest struct {
	// The code of the agent that is already online.
	//
	// example:
	//
	// 1180**************
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application code. This parameter is used when the creation source is engine.
	//
	// example:
	//
	// 025****C98
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// The list of callable days.
	//
	// This parameter is required.
	CallDayShrink *string `json:"CallDay,omitempty" xml:"CallDay,omitempty"`
	// The expiration date of outbound call details (specific deadline).
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
	// The retry interval. Unit: minutes. The maximum value is 720 minutes.
	//
	// example:
	//
	// 32
	CallRetryInterval *int64 `json:"CallRetryInterval,omitempty" xml:"CallRetryInterval,omitempty"`
	// The list of failure retry reasons.
	CallRetryReasonShrink *string `json:"CallRetryReason,omitempty" xml:"CallRetryReason,omitempty"`
	// The number of retries. The maximum value is 3.
	//
	// example:
	//
	// 2
	CallRetryTimes *int64 `json:"CallRetryTimes,omitempty" xml:"CallRetryTimes,omitempty"`
	// The list of callable time periods.
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
	// - true: enabled.
	//
	// - false (default): disabled.
	//
	// example:
	//
	// false
	MissCallRetry *bool  `json:"MissCallRetry,omitempty" xml:"MissCallRetry,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number type. This parameter is used when the creation source is engine. Valid values:
	//
	// - 0: Alibaba Cloud number.
	//
	// - 1: customer-provided line.
	//
	// example:
	//
	// 0
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The creation source. Valid values:
	//
	// - 0: created by agent.
	//
	// - 1: created by engine.
	//
	// example:
	//
	// 0
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
	// SCHEDULE
	StartType *string `json:"StartType,omitempty" xml:"StartType,omitempty"`
	// The task concurrency. The maximum value is 500.
	//
	// example:
	//
	// 75
	TaskCps *int64 `json:"TaskCps,omitempty" xml:"TaskCps,omitempty"`
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
	// 12313123133
	TaskStartTime *int64 `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// The service instance used for outbound calls.
	//
	// example:
	//
	// 032712122*****
	VirtualNumber *string `json:"VirtualNumber,omitempty" xml:"VirtualNumber,omitempty"`
}

func (s CreateAiCallTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiCallTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAiCallTaskShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAiCallTaskShrinkRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *CreateAiCallTaskShrinkRequest) GetCallDayShrink() *string {
	return s.CallDayShrink
}

func (s *CreateAiCallTaskShrinkRequest) GetCallExpireDate() *string {
	return s.CallExpireDate
}

func (s *CreateAiCallTaskShrinkRequest) GetCallExpireMinutes() *int64 {
	return s.CallExpireMinutes
}

func (s *CreateAiCallTaskShrinkRequest) GetCallExpireType() *int64 {
	return s.CallExpireType
}

func (s *CreateAiCallTaskShrinkRequest) GetCallRetryInterval() *int64 {
	return s.CallRetryInterval
}

func (s *CreateAiCallTaskShrinkRequest) GetCallRetryReasonShrink() *string {
	return s.CallRetryReasonShrink
}

func (s *CreateAiCallTaskShrinkRequest) GetCallRetryTimes() *int64 {
	return s.CallRetryTimes
}

func (s *CreateAiCallTaskShrinkRequest) GetCallTimeShrink() *string {
	return s.CallTimeShrink
}

func (s *CreateAiCallTaskShrinkRequest) GetCallableTimeShrink() *string {
	return s.CallableTimeShrink
}

func (s *CreateAiCallTaskShrinkRequest) GetLineEncoding() *string {
	return s.LineEncoding
}

func (s *CreateAiCallTaskShrinkRequest) GetLinePhoneNum() *string {
	return s.LinePhoneNum
}

func (s *CreateAiCallTaskShrinkRequest) GetMissCallRetry() *bool {
	return s.MissCallRetry
}

func (s *CreateAiCallTaskShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAiCallTaskShrinkRequest) GetPhoneType() *int64 {
	return s.PhoneType
}

func (s *CreateAiCallTaskShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAiCallTaskShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAiCallTaskShrinkRequest) GetSource() *int64 {
	return s.Source
}

func (s *CreateAiCallTaskShrinkRequest) GetStartType() *string {
	return s.StartType
}

func (s *CreateAiCallTaskShrinkRequest) GetTaskCps() *int64 {
	return s.TaskCps
}

func (s *CreateAiCallTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateAiCallTaskShrinkRequest) GetTaskStartTime() *int64 {
	return s.TaskStartTime
}

func (s *CreateAiCallTaskShrinkRequest) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *CreateAiCallTaskShrinkRequest) SetAgentId(v string) *CreateAiCallTaskShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetApplicationCode(v string) *CreateAiCallTaskShrinkRequest {
	s.ApplicationCode = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallDayShrink(v string) *CreateAiCallTaskShrinkRequest {
	s.CallDayShrink = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallExpireDate(v string) *CreateAiCallTaskShrinkRequest {
	s.CallExpireDate = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallExpireMinutes(v int64) *CreateAiCallTaskShrinkRequest {
	s.CallExpireMinutes = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallExpireType(v int64) *CreateAiCallTaskShrinkRequest {
	s.CallExpireType = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallRetryInterval(v int64) *CreateAiCallTaskShrinkRequest {
	s.CallRetryInterval = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallRetryReasonShrink(v string) *CreateAiCallTaskShrinkRequest {
	s.CallRetryReasonShrink = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallRetryTimes(v int64) *CreateAiCallTaskShrinkRequest {
	s.CallRetryTimes = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallTimeShrink(v string) *CreateAiCallTaskShrinkRequest {
	s.CallTimeShrink = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetCallableTimeShrink(v string) *CreateAiCallTaskShrinkRequest {
	s.CallableTimeShrink = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetLineEncoding(v string) *CreateAiCallTaskShrinkRequest {
	s.LineEncoding = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetLinePhoneNum(v string) *CreateAiCallTaskShrinkRequest {
	s.LinePhoneNum = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetMissCallRetry(v bool) *CreateAiCallTaskShrinkRequest {
	s.MissCallRetry = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetOwnerId(v int64) *CreateAiCallTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetPhoneType(v int64) *CreateAiCallTaskShrinkRequest {
	s.PhoneType = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetResourceOwnerAccount(v string) *CreateAiCallTaskShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetResourceOwnerId(v int64) *CreateAiCallTaskShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetSource(v int64) *CreateAiCallTaskShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetStartType(v string) *CreateAiCallTaskShrinkRequest {
	s.StartType = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetTaskCps(v int64) *CreateAiCallTaskShrinkRequest {
	s.TaskCps = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetTaskName(v string) *CreateAiCallTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetTaskStartTime(v int64) *CreateAiCallTaskShrinkRequest {
	s.TaskStartTime = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) SetVirtualNumber(v string) *CreateAiCallTaskShrinkRequest {
	s.VirtualNumber = &v
	return s
}

func (s *CreateAiCallTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

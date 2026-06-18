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
	SetCallRetryInterval(v int64) *CreateAiCallTaskShrinkRequest
	GetCallRetryInterval() *int64
	SetCallRetryReasonShrink(v string) *CreateAiCallTaskShrinkRequest
	GetCallRetryReasonShrink() *string
	SetCallRetryTimes(v int64) *CreateAiCallTaskShrinkRequest
	GetCallRetryTimes() *int64
	SetCallTimeShrink(v string) *CreateAiCallTaskShrinkRequest
	GetCallTimeShrink() *string
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
	// The ID of a published agent.
	//
	// example:
	//
	// 1180**************
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 025****C98
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// The list of callable days.
	//
	// This parameter is required.
	CallDayShrink *string `json:"CallDay,omitempty" xml:"CallDay,omitempty"`
	// The retry interval, in minutes. The maximum value is 720.
	//
	// example:
	//
	// 32
	CallRetryInterval *int64 `json:"CallRetryInterval,omitempty" xml:"CallRetryInterval,omitempty"`
	// The failure reasons that trigger a retry.
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
	CallTimeShrink *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// example:
	//
	// JILIANG_***_***_NET
	LineEncoding *string `json:"LineEncoding,omitempty" xml:"LineEncoding,omitempty"`
	// example:
	//
	// 152****3120
	LinePhoneNum *string `json:"LinePhoneNum,omitempty" xml:"LinePhoneNum,omitempty"`
	// Specifies whether to enable retry. Valid values:
	//
	// - `true`: Enables retry.
	//
	// - `false` (default): Disables retry.
	//
	// example:
	//
	// false
	MissCallRetry *bool  `json:"MissCallRetry,omitempty" xml:"MissCallRetry,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 0
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 0
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The startup mode. Valid values:
	//
	// - `IMMEDIATE`: Starts the task immediately.
	//
	// - `SCHEDULE`: Starts the task at a scheduled time.
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
	// The name of the task. It must be unique within an account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试任务
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The scheduled start time for the task, specified as a timestamp in milliseconds. This parameter is required and applies only when `StartType` is set to `SCHEDULE`.
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

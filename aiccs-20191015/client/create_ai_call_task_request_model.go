// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiCallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateAiCallTaskRequest
	GetAgentId() *string
	SetApplicationCode(v string) *CreateAiCallTaskRequest
	GetApplicationCode() *string
	SetCallDay(v []*string) *CreateAiCallTaskRequest
	GetCallDay() []*string
	SetCallExpireDate(v string) *CreateAiCallTaskRequest
	GetCallExpireDate() *string
	SetCallExpireMinutes(v int64) *CreateAiCallTaskRequest
	GetCallExpireMinutes() *int64
	SetCallExpireType(v int64) *CreateAiCallTaskRequest
	GetCallExpireType() *int64
	SetCallRetryInterval(v int64) *CreateAiCallTaskRequest
	GetCallRetryInterval() *int64
	SetCallRetryReason(v []*string) *CreateAiCallTaskRequest
	GetCallRetryReason() []*string
	SetCallRetryTimes(v int64) *CreateAiCallTaskRequest
	GetCallRetryTimes() *int64
	SetCallTime(v []*string) *CreateAiCallTaskRequest
	GetCallTime() []*string
	SetCallableTime(v []*string) *CreateAiCallTaskRequest
	GetCallableTime() []*string
	SetLineEncoding(v string) *CreateAiCallTaskRequest
	GetLineEncoding() *string
	SetLinePhoneNum(v string) *CreateAiCallTaskRequest
	GetLinePhoneNum() *string
	SetMissCallRetry(v bool) *CreateAiCallTaskRequest
	GetMissCallRetry() *bool
	SetOwnerId(v int64) *CreateAiCallTaskRequest
	GetOwnerId() *int64
	SetPhoneType(v int64) *CreateAiCallTaskRequest
	GetPhoneType() *int64
	SetResourceOwnerAccount(v string) *CreateAiCallTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAiCallTaskRequest
	GetResourceOwnerId() *int64
	SetSource(v int64) *CreateAiCallTaskRequest
	GetSource() *int64
	SetStartType(v string) *CreateAiCallTaskRequest
	GetStartType() *string
	SetTaskCps(v int64) *CreateAiCallTaskRequest
	GetTaskCps() *int64
	SetTaskName(v string) *CreateAiCallTaskRequest
	GetTaskName() *string
	SetTaskStartTime(v int64) *CreateAiCallTaskRequest
	GetTaskStartTime() *int64
	SetVirtualNumber(v string) *CreateAiCallTaskRequest
	GetVirtualNumber() *string
}

type CreateAiCallTaskRequest struct {
	// The code of the agent that has been published.
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
	CallDay []*string `json:"CallDay,omitempty" xml:"CallDay,omitempty" type:"Repeated"`
	// The expiration date of outbound call details (specific deadline). Format: YYYY-MM-DD HH:mm:ss.
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
	// 0: Permanently valid.
	//
	// 1: Valid for a specified duration after import.
	//
	// 2: Valid until a specified date.
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
	// The list of retry reasons for failed calls.
	CallRetryReason []*string `json:"CallRetryReason,omitempty" xml:"CallRetryReason,omitempty" type:"Repeated"`
	// The number of retries. The maximum value is 3.
	//
	// example:
	//
	// 2
	CallRetryTimes *int64 `json:"CallRetryTimes,omitempty" xml:"CallRetryTimes,omitempty"`
	// The list of callable time periods.
	//
	// This parameter is required.
	CallTime     []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
	CallableTime []*string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty" type:"Repeated"`
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
	// false
	MissCallRetry *bool  `json:"MissCallRetry,omitempty" xml:"MissCallRetry,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number type. This parameter is used when the creation source is engine. Valid values:
	//
	// - 0: Alibaba Cloud number.
	//
	// - 1: Customer-provided line.
	//
	// example:
	//
	// 0
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The creation source. Valid values:
	//
	// - 0: Created by an agent.
	//
	// - 1: Created by an engine.
	//
	// example:
	//
	// 0
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start type. Valid values:
	//
	// - IMMEDIATE: Start immediately.
	//
	// - SCHEDULE: Start at a scheduled time.
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
	// The preset start time of the task. The value is a UNIX timestamp in milliseconds. This parameter is valid and required when StartType is set to SCHEDULE. The task automatically starts at the time specified by this parameter.
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

func (s CreateAiCallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiCallTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAiCallTaskRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAiCallTaskRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *CreateAiCallTaskRequest) GetCallDay() []*string {
	return s.CallDay
}

func (s *CreateAiCallTaskRequest) GetCallExpireDate() *string {
	return s.CallExpireDate
}

func (s *CreateAiCallTaskRequest) GetCallExpireMinutes() *int64 {
	return s.CallExpireMinutes
}

func (s *CreateAiCallTaskRequest) GetCallExpireType() *int64 {
	return s.CallExpireType
}

func (s *CreateAiCallTaskRequest) GetCallRetryInterval() *int64 {
	return s.CallRetryInterval
}

func (s *CreateAiCallTaskRequest) GetCallRetryReason() []*string {
	return s.CallRetryReason
}

func (s *CreateAiCallTaskRequest) GetCallRetryTimes() *int64 {
	return s.CallRetryTimes
}

func (s *CreateAiCallTaskRequest) GetCallTime() []*string {
	return s.CallTime
}

func (s *CreateAiCallTaskRequest) GetCallableTime() []*string {
	return s.CallableTime
}

func (s *CreateAiCallTaskRequest) GetLineEncoding() *string {
	return s.LineEncoding
}

func (s *CreateAiCallTaskRequest) GetLinePhoneNum() *string {
	return s.LinePhoneNum
}

func (s *CreateAiCallTaskRequest) GetMissCallRetry() *bool {
	return s.MissCallRetry
}

func (s *CreateAiCallTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAiCallTaskRequest) GetPhoneType() *int64 {
	return s.PhoneType
}

func (s *CreateAiCallTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAiCallTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAiCallTaskRequest) GetSource() *int64 {
	return s.Source
}

func (s *CreateAiCallTaskRequest) GetStartType() *string {
	return s.StartType
}

func (s *CreateAiCallTaskRequest) GetTaskCps() *int64 {
	return s.TaskCps
}

func (s *CreateAiCallTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateAiCallTaskRequest) GetTaskStartTime() *int64 {
	return s.TaskStartTime
}

func (s *CreateAiCallTaskRequest) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *CreateAiCallTaskRequest) SetAgentId(v string) *CreateAiCallTaskRequest {
	s.AgentId = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetApplicationCode(v string) *CreateAiCallTaskRequest {
	s.ApplicationCode = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallDay(v []*string) *CreateAiCallTaskRequest {
	s.CallDay = v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallExpireDate(v string) *CreateAiCallTaskRequest {
	s.CallExpireDate = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallExpireMinutes(v int64) *CreateAiCallTaskRequest {
	s.CallExpireMinutes = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallExpireType(v int64) *CreateAiCallTaskRequest {
	s.CallExpireType = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallRetryInterval(v int64) *CreateAiCallTaskRequest {
	s.CallRetryInterval = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallRetryReason(v []*string) *CreateAiCallTaskRequest {
	s.CallRetryReason = v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallRetryTimes(v int64) *CreateAiCallTaskRequest {
	s.CallRetryTimes = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallTime(v []*string) *CreateAiCallTaskRequest {
	s.CallTime = v
	return s
}

func (s *CreateAiCallTaskRequest) SetCallableTime(v []*string) *CreateAiCallTaskRequest {
	s.CallableTime = v
	return s
}

func (s *CreateAiCallTaskRequest) SetLineEncoding(v string) *CreateAiCallTaskRequest {
	s.LineEncoding = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetLinePhoneNum(v string) *CreateAiCallTaskRequest {
	s.LinePhoneNum = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetMissCallRetry(v bool) *CreateAiCallTaskRequest {
	s.MissCallRetry = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetOwnerId(v int64) *CreateAiCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetPhoneType(v int64) *CreateAiCallTaskRequest {
	s.PhoneType = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetResourceOwnerAccount(v string) *CreateAiCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetResourceOwnerId(v int64) *CreateAiCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetSource(v int64) *CreateAiCallTaskRequest {
	s.Source = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetStartType(v string) *CreateAiCallTaskRequest {
	s.StartType = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetTaskCps(v int64) *CreateAiCallTaskRequest {
	s.TaskCps = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetTaskName(v string) *CreateAiCallTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetTaskStartTime(v int64) *CreateAiCallTaskRequest {
	s.TaskStartTime = &v
	return s
}

func (s *CreateAiCallTaskRequest) SetVirtualNumber(v string) *CreateAiCallTaskRequest {
	s.VirtualNumber = &v
	return s
}

func (s *CreateAiCallTaskRequest) Validate() error {
	return dara.Validate(s)
}

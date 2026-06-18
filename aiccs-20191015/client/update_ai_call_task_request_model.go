// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiCallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallDay(v []*string) *UpdateAiCallTaskRequest
	GetCallDay() []*string
	SetCallRetryInterval(v int64) *UpdateAiCallTaskRequest
	GetCallRetryInterval() *int64
	SetCallRetryReason(v []*string) *UpdateAiCallTaskRequest
	GetCallRetryReason() []*string
	SetCallRetryTimes(v int64) *UpdateAiCallTaskRequest
	GetCallRetryTimes() *int64
	SetCallTime(v []*string) *UpdateAiCallTaskRequest
	GetCallTime() []*string
	SetLineEncoding(v string) *UpdateAiCallTaskRequest
	GetLineEncoding() *string
	SetLinePhoneNum(v string) *UpdateAiCallTaskRequest
	GetLinePhoneNum() *string
	SetMissCallRetry(v bool) *UpdateAiCallTaskRequest
	GetMissCallRetry() *bool
	SetOwnerId(v int64) *UpdateAiCallTaskRequest
	GetOwnerId() *int64
	SetPhoneType(v int64) *UpdateAiCallTaskRequest
	GetPhoneType() *int64
	SetResourceOwnerAccount(v string) *UpdateAiCallTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAiCallTaskRequest
	GetResourceOwnerId() *int64
	SetSource(v int64) *UpdateAiCallTaskRequest
	GetSource() *int64
	SetStartType(v string) *UpdateAiCallTaskRequest
	GetStartType() *string
	SetTaskCps(v int64) *UpdateAiCallTaskRequest
	GetTaskCps() *int64
	SetTaskId(v string) *UpdateAiCallTaskRequest
	GetTaskId() *string
	SetTaskName(v string) *UpdateAiCallTaskRequest
	GetTaskName() *string
	SetTaskStartTime(v int64) *UpdateAiCallTaskRequest
	GetTaskStartTime() *int64
	SetVirtualNumber(v string) *UpdateAiCallTaskRequest
	GetVirtualNumber() *string
}

type UpdateAiCallTaskRequest struct {
	// The days of the week when calls can be made.
	//
	// This parameter is required.
	CallDay []*string `json:"CallDay,omitempty" xml:"CallDay,omitempty" type:"Repeated"`
	// The retry interval in minutes. The maximum value is 120.
	//
	// example:
	//
	// 25
	CallRetryInterval *int64 `json:"CallRetryInterval,omitempty" xml:"CallRetryInterval,omitempty"`
	// The call failure statuses that trigger a retry.
	CallRetryReason []*string `json:"CallRetryReason,omitempty" xml:"CallRetryReason,omitempty" type:"Repeated"`
	// The number of retries. The maximum value is 3.
	//
	// example:
	//
	// 2
	CallRetryTimes *int64 `json:"CallRetryTimes,omitempty" xml:"CallRetryTimes,omitempty"`
	// The callable time windows.
	//
	// This parameter is required.
	CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
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

func (s UpdateAiCallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiCallTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiCallTaskRequest) GetCallDay() []*string {
	return s.CallDay
}

func (s *UpdateAiCallTaskRequest) GetCallRetryInterval() *int64 {
	return s.CallRetryInterval
}

func (s *UpdateAiCallTaskRequest) GetCallRetryReason() []*string {
	return s.CallRetryReason
}

func (s *UpdateAiCallTaskRequest) GetCallRetryTimes() *int64 {
	return s.CallRetryTimes
}

func (s *UpdateAiCallTaskRequest) GetCallTime() []*string {
	return s.CallTime
}

func (s *UpdateAiCallTaskRequest) GetLineEncoding() *string {
	return s.LineEncoding
}

func (s *UpdateAiCallTaskRequest) GetLinePhoneNum() *string {
	return s.LinePhoneNum
}

func (s *UpdateAiCallTaskRequest) GetMissCallRetry() *bool {
	return s.MissCallRetry
}

func (s *UpdateAiCallTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAiCallTaskRequest) GetPhoneType() *int64 {
	return s.PhoneType
}

func (s *UpdateAiCallTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAiCallTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAiCallTaskRequest) GetSource() *int64 {
	return s.Source
}

func (s *UpdateAiCallTaskRequest) GetStartType() *string {
	return s.StartType
}

func (s *UpdateAiCallTaskRequest) GetTaskCps() *int64 {
	return s.TaskCps
}

func (s *UpdateAiCallTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAiCallTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateAiCallTaskRequest) GetTaskStartTime() *int64 {
	return s.TaskStartTime
}

func (s *UpdateAiCallTaskRequest) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *UpdateAiCallTaskRequest) SetCallDay(v []*string) *UpdateAiCallTaskRequest {
	s.CallDay = v
	return s
}

func (s *UpdateAiCallTaskRequest) SetCallRetryInterval(v int64) *UpdateAiCallTaskRequest {
	s.CallRetryInterval = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetCallRetryReason(v []*string) *UpdateAiCallTaskRequest {
	s.CallRetryReason = v
	return s
}

func (s *UpdateAiCallTaskRequest) SetCallRetryTimes(v int64) *UpdateAiCallTaskRequest {
	s.CallRetryTimes = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetCallTime(v []*string) *UpdateAiCallTaskRequest {
	s.CallTime = v
	return s
}

func (s *UpdateAiCallTaskRequest) SetLineEncoding(v string) *UpdateAiCallTaskRequest {
	s.LineEncoding = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetLinePhoneNum(v string) *UpdateAiCallTaskRequest {
	s.LinePhoneNum = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetMissCallRetry(v bool) *UpdateAiCallTaskRequest {
	s.MissCallRetry = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetOwnerId(v int64) *UpdateAiCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetPhoneType(v int64) *UpdateAiCallTaskRequest {
	s.PhoneType = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetResourceOwnerAccount(v string) *UpdateAiCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetResourceOwnerId(v int64) *UpdateAiCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetSource(v int64) *UpdateAiCallTaskRequest {
	s.Source = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetStartType(v string) *UpdateAiCallTaskRequest {
	s.StartType = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetTaskCps(v int64) *UpdateAiCallTaskRequest {
	s.TaskCps = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetTaskId(v string) *UpdateAiCallTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetTaskName(v string) *UpdateAiCallTaskRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetTaskStartTime(v int64) *UpdateAiCallTaskRequest {
	s.TaskStartTime = &v
	return s
}

func (s *UpdateAiCallTaskRequest) SetVirtualNumber(v string) *UpdateAiCallTaskRequest {
	s.VirtualNumber = &v
	return s
}

func (s *UpdateAiCallTaskRequest) Validate() error {
	return dara.Validate(s)
}

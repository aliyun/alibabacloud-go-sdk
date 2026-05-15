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
	SetCallRetryInterval(v int64) *CreateAiCallTaskRequest
	GetCallRetryInterval() *int64
	SetCallRetryReason(v []*string) *CreateAiCallTaskRequest
	GetCallRetryReason() []*string
	SetCallRetryTimes(v int64) *CreateAiCallTaskRequest
	GetCallRetryTimes() *int64
	SetCallTime(v []*string) *CreateAiCallTaskRequest
	GetCallTime() []*string
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
	// example:
	//
	// 121223123123****
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值示例值
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// This parameter is required.
	CallDay []*string `json:"CallDay,omitempty" xml:"CallDay,omitempty" type:"Repeated"`
	// example:
	//
	// 32
	CallRetryInterval *int64    `json:"CallRetryInterval,omitempty" xml:"CallRetryInterval,omitempty"`
	CallRetryReason   []*string `json:"CallRetryReason,omitempty" xml:"CallRetryReason,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CallRetryTimes *int64 `json:"CallRetryTimes,omitempty" xml:"CallRetryTimes,omitempty"`
	// This parameter is required.
	CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	LineEncoding *string `json:"LineEncoding,omitempty" xml:"LineEncoding,omitempty"`
	// example:
	//
	// 示例值示例值
	LinePhoneNum *string `json:"LinePhoneNum,omitempty" xml:"LinePhoneNum,omitempty"`
	// example:
	//
	// false
	MissCallRetry *bool  `json:"MissCallRetry,omitempty" xml:"MissCallRetry,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 54
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 92
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SCHEDULE
	StartType *string `json:"StartType,omitempty" xml:"StartType,omitempty"`
	// example:
	//
	// 75
	TaskCps *int64 `json:"TaskCps,omitempty" xml:"TaskCps,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 12313123133
	TaskStartTime *int64 `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
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

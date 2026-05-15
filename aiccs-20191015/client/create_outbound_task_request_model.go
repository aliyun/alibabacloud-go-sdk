// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAni(v string) *CreateOutboundTaskRequest
	GetAni() *string
	SetCallInfos(v string) *CreateOutboundTaskRequest
	GetCallInfos() *string
	SetDepartmentId(v int64) *CreateOutboundTaskRequest
	GetDepartmentId() *int64
	SetDescription(v string) *CreateOutboundTaskRequest
	GetDescription() *string
	SetEndDate(v string) *CreateOutboundTaskRequest
	GetEndDate() *string
	SetEndTime(v string) *CreateOutboundTaskRequest
	GetEndTime() *string
	SetExtAttrs(v string) *CreateOutboundTaskRequest
	GetExtAttrs() *string
	SetGroupName(v string) *CreateOutboundTaskRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateOutboundTaskRequest
	GetInstanceId() *string
	SetModel(v int32) *CreateOutboundTaskRequest
	GetModel() *int32
	SetRetryInterval(v int32) *CreateOutboundTaskRequest
	GetRetryInterval() *int32
	SetRetryTime(v int32) *CreateOutboundTaskRequest
	GetRetryTime() *int32
	SetSkillGroup(v int64) *CreateOutboundTaskRequest
	GetSkillGroup() *int64
	SetStartDate(v string) *CreateOutboundTaskRequest
	GetStartDate() *string
	SetStartTime(v string) *CreateOutboundTaskRequest
	GetStartTime() *string
	SetTaskName(v string) *CreateOutboundTaskRequest
	GetTaskName() *string
	SetTaskType(v int32) *CreateOutboundTaskRequest
	GetTaskType() *int32
}

type CreateOutboundTaskRequest struct {
	// This parameter is required.
	Ani          *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	CallInfos    *string `json:"CallInfos,omitempty" xml:"CallInfos,omitempty"`
	DepartmentId *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtAttrs  *string `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Model         *int32  `json:"Model,omitempty" xml:"Model,omitempty"`
	RetryInterval *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryTime     *int32  `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// This parameter is required.
	SkillGroup *int64 `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	// This parameter is required.
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOutboundTaskRequest) GetAni() *string {
	return s.Ani
}

func (s *CreateOutboundTaskRequest) GetCallInfos() *string {
	return s.CallInfos
}

func (s *CreateOutboundTaskRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *CreateOutboundTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOutboundTaskRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateOutboundTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateOutboundTaskRequest) GetExtAttrs() *string {
	return s.ExtAttrs
}

func (s *CreateOutboundTaskRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOutboundTaskRequest) GetModel() *int32 {
	return s.Model
}

func (s *CreateOutboundTaskRequest) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *CreateOutboundTaskRequest) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *CreateOutboundTaskRequest) GetSkillGroup() *int64 {
	return s.SkillGroup
}

func (s *CreateOutboundTaskRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateOutboundTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateOutboundTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateOutboundTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateOutboundTaskRequest) SetAni(v string) *CreateOutboundTaskRequest {
	s.Ani = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetCallInfos(v string) *CreateOutboundTaskRequest {
	s.CallInfos = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetDepartmentId(v int64) *CreateOutboundTaskRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetDescription(v string) *CreateOutboundTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetEndDate(v string) *CreateOutboundTaskRequest {
	s.EndDate = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetEndTime(v string) *CreateOutboundTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetExtAttrs(v string) *CreateOutboundTaskRequest {
	s.ExtAttrs = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetGroupName(v string) *CreateOutboundTaskRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetInstanceId(v string) *CreateOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetModel(v int32) *CreateOutboundTaskRequest {
	s.Model = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetRetryInterval(v int32) *CreateOutboundTaskRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetRetryTime(v int32) *CreateOutboundTaskRequest {
	s.RetryTime = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetSkillGroup(v int64) *CreateOutboundTaskRequest {
	s.SkillGroup = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetStartDate(v string) *CreateOutboundTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetStartTime(v string) *CreateOutboundTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetTaskName(v string) *CreateOutboundTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateOutboundTaskRequest) SetTaskType(v int32) *CreateOutboundTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}

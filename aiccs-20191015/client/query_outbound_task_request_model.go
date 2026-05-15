// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAni(v string) *QueryOutboundTaskRequest
	GetAni() *string
	SetCurrentPage(v int32) *QueryOutboundTaskRequest
	GetCurrentPage() *int32
	SetDepartmentId(v string) *QueryOutboundTaskRequest
	GetDepartmentId() *string
	SetEndDate(v string) *QueryOutboundTaskRequest
	GetEndDate() *string
	SetEndTime(v string) *QueryOutboundTaskRequest
	GetEndTime() *string
	SetGroupName(v string) *QueryOutboundTaskRequest
	GetGroupName() *string
	SetInstanceId(v string) *QueryOutboundTaskRequest
	GetInstanceId() *string
	SetPageSize(v int32) *QueryOutboundTaskRequest
	GetPageSize() *int32
	SetSkillGroup(v int64) *QueryOutboundTaskRequest
	GetSkillGroup() *int64
	SetStartDate(v string) *QueryOutboundTaskRequest
	GetStartDate() *string
	SetStartTime(v string) *QueryOutboundTaskRequest
	GetStartTime() *string
	SetStatus(v string) *QueryOutboundTaskRequest
	GetStatus() *string
	SetTaskId(v int64) *QueryOutboundTaskRequest
	GetTaskId() *int64
	SetTaskName(v string) *QueryOutboundTaskRequest
	GetTaskName() *string
	SetTaskType(v int32) *QueryOutboundTaskRequest
	GetTaskType() *int32
}

type QueryOutboundTaskRequest struct {
	Ani          *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroup *int64  `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskType   *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskRequest) GetAni() *string {
	return s.Ani
}

func (s *QueryOutboundTaskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryOutboundTaskRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *QueryOutboundTaskRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryOutboundTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryOutboundTaskRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryOutboundTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOutboundTaskRequest) GetSkillGroup() *int64 {
	return s.SkillGroup
}

func (s *QueryOutboundTaskRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryOutboundTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryOutboundTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryOutboundTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryOutboundTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryOutboundTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *QueryOutboundTaskRequest) SetAni(v string) *QueryOutboundTaskRequest {
	s.Ani = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetCurrentPage(v int32) *QueryOutboundTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetDepartmentId(v string) *QueryOutboundTaskRequest {
	s.DepartmentId = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetEndDate(v string) *QueryOutboundTaskRequest {
	s.EndDate = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetEndTime(v string) *QueryOutboundTaskRequest {
	s.EndTime = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetGroupName(v string) *QueryOutboundTaskRequest {
	s.GroupName = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetInstanceId(v string) *QueryOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetPageSize(v int32) *QueryOutboundTaskRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetSkillGroup(v int64) *QueryOutboundTaskRequest {
	s.SkillGroup = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetStartDate(v string) *QueryOutboundTaskRequest {
	s.StartDate = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetStartTime(v string) *QueryOutboundTaskRequest {
	s.StartTime = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetStatus(v string) *QueryOutboundTaskRequest {
	s.Status = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetTaskId(v int64) *QueryOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetTaskName(v string) *QueryOutboundTaskRequest {
	s.TaskName = &v
	return s
}

func (s *QueryOutboundTaskRequest) SetTaskType(v int32) *QueryOutboundTaskRequest {
	s.TaskType = &v
	return s
}

func (s *QueryOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskInstance interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreatedTime(v string) *TaskInstance
	GetGmtCreatedTime() *string
	SetGmtEndTime(v string) *TaskInstance
	GetGmtEndTime() *string
	SetStatus(v string) *TaskInstance
	GetStatus() *string
	SetTaskId(v string) *TaskInstance
	GetTaskId() *string
	SetTaskInstanceId(v string) *TaskInstance
	GetTaskInstanceId() *string
	SetTenantId(v string) *TaskInstance
	GetTenantId() *string
	SetUserId(v string) *TaskInstance
	GetUserId() *string
}

type TaskInstance struct {
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtEndTime     *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	TenantId       *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s TaskInstance) String() string {
	return dara.Prettify(s)
}

func (s TaskInstance) GoString() string {
	return s.String()
}

func (s *TaskInstance) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *TaskInstance) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *TaskInstance) GetStatus() *string {
	return s.Status
}

func (s *TaskInstance) GetTaskId() *string {
	return s.TaskId
}

func (s *TaskInstance) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *TaskInstance) GetTenantId() *string {
	return s.TenantId
}

func (s *TaskInstance) GetUserId() *string {
	return s.UserId
}

func (s *TaskInstance) SetGmtCreatedTime(v string) *TaskInstance {
	s.GmtCreatedTime = &v
	return s
}

func (s *TaskInstance) SetGmtEndTime(v string) *TaskInstance {
	s.GmtEndTime = &v
	return s
}

func (s *TaskInstance) SetStatus(v string) *TaskInstance {
	s.Status = &v
	return s
}

func (s *TaskInstance) SetTaskId(v string) *TaskInstance {
	s.TaskId = &v
	return s
}

func (s *TaskInstance) SetTaskInstanceId(v string) *TaskInstance {
	s.TaskInstanceId = &v
	return s
}

func (s *TaskInstance) SetTenantId(v string) *TaskInstance {
	s.TenantId = &v
	return s
}

func (s *TaskInstance) SetUserId(v string) *TaskInstance {
	s.UserId = &v
	return s
}

func (s *TaskInstance) Validate() error {
	return dara.Validate(s)
}

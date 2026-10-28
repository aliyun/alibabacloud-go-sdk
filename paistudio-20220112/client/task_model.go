// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTask interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*Action) *Task
	GetActions() []*Action
	SetDescription(v string) *Task
	GetDescription() *string
	SetGmtActivatedTime(v string) *Task
	GetGmtActivatedTime() *string
	SetGmtCreatedTime(v string) *Task
	GetGmtCreatedTime() *string
	SetGmtModifiedTime(v string) *Task
	GetGmtModifiedTime() *string
	SetGmtStoppedTime(v string) *Task
	GetGmtStoppedTime() *string
	SetQuotaId(v string) *Task
	GetQuotaId() *string
	SetRules(v *Rules) *Task
	GetRules() *Rules
	SetStatus(v string) *Task
	GetStatus() *string
	SetTaskId(v string) *Task
	GetTaskId() *string
	SetTaskName(v string) *Task
	GetTaskName() *string
	SetUserId(v string) *Task
	GetUserId() *string
	SetUserName(v string) *Task
	GetUserName() *string
}

type Task struct {
	// The actions performed by the task.
	Actions []*Action `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The description of the task.
	//
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the task was activated.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtActivatedTime *string `json:"GmtActivatedTime,omitempty" xml:"GmtActivatedTime,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The time when the task was stopped.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtStoppedTime *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// quotad2kd8ljpsno
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The execution rules of the task.
	Rules *Rules `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// taskxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Task) String() string {
	return dara.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) GetActions() []*Action {
	return s.Actions
}

func (s *Task) GetDescription() *string {
	return s.Description
}

func (s *Task) GetGmtActivatedTime() *string {
	return s.GmtActivatedTime
}

func (s *Task) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *Task) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Task) GetGmtStoppedTime() *string {
	return s.GmtStoppedTime
}

func (s *Task) GetQuotaId() *string {
	return s.QuotaId
}

func (s *Task) GetRules() *Rules {
	return s.Rules
}

func (s *Task) GetStatus() *string {
	return s.Status
}

func (s *Task) GetTaskId() *string {
	return s.TaskId
}

func (s *Task) GetTaskName() *string {
	return s.TaskName
}

func (s *Task) GetUserId() *string {
	return s.UserId
}

func (s *Task) GetUserName() *string {
	return s.UserName
}

func (s *Task) SetActions(v []*Action) *Task {
	s.Actions = v
	return s
}

func (s *Task) SetDescription(v string) *Task {
	s.Description = &v
	return s
}

func (s *Task) SetGmtActivatedTime(v string) *Task {
	s.GmtActivatedTime = &v
	return s
}

func (s *Task) SetGmtCreatedTime(v string) *Task {
	s.GmtCreatedTime = &v
	return s
}

func (s *Task) SetGmtModifiedTime(v string) *Task {
	s.GmtModifiedTime = &v
	return s
}

func (s *Task) SetGmtStoppedTime(v string) *Task {
	s.GmtStoppedTime = &v
	return s
}

func (s *Task) SetQuotaId(v string) *Task {
	s.QuotaId = &v
	return s
}

func (s *Task) SetRules(v *Rules) *Task {
	s.Rules = v
	return s
}

func (s *Task) SetStatus(v string) *Task {
	s.Status = &v
	return s
}

func (s *Task) SetTaskId(v string) *Task {
	s.TaskId = &v
	return s
}

func (s *Task) SetTaskName(v string) *Task {
	s.TaskName = &v
	return s
}

func (s *Task) SetUserId(v string) *Task {
	s.UserId = &v
	return s
}

func (s *Task) SetUserName(v string) *Task {
	s.UserName = &v
	return s
}

func (s *Task) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

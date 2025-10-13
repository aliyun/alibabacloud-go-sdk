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
	Actions          []*Action `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtActivatedTime *string   `json:"GmtActivatedTime,omitempty" xml:"GmtActivatedTime,omitempty"`
	GmtCreatedTime   *string   `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime  *string   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStoppedTime   *string   `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	QuotaId          *string   `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	Rules            *Rules    `json:"Rules,omitempty" xml:"Rules,omitempty"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId           *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName         *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	UserId           *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName         *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

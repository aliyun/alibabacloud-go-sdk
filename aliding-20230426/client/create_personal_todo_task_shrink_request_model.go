// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTodoTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalTodoTaskShrinkRequest
	GetDescription() *string
	SetDueTime(v int64) *CreatePersonalTodoTaskShrinkRequest
	GetDueTime() *int64
	SetExecutorIdsShrink(v string) *CreatePersonalTodoTaskShrinkRequest
	GetExecutorIdsShrink() *string
	SetNotifyConfigsShrink(v string) *CreatePersonalTodoTaskShrinkRequest
	GetNotifyConfigsShrink() *string
	SetParticipantIdsShrink(v string) *CreatePersonalTodoTaskShrinkRequest
	GetParticipantIdsShrink() *string
	SetReminderTimeStamp(v int64) *CreatePersonalTodoTaskShrinkRequest
	GetReminderTimeStamp() *int64
	SetSubject(v string) *CreatePersonalTodoTaskShrinkRequest
	GetSubject() *string
	SetTenantContextShrink(v string) *CreatePersonalTodoTaskShrinkRequest
	GetTenantContextShrink() *string
}

type CreatePersonalTodoTaskShrinkRequest struct {
	// example:
	//
	// 待办备注信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1703750708595
	DueTime *int64 `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [ "012345" ]
	ExecutorIdsShrink   *string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty"`
	NotifyConfigsShrink *string `json:"NotifyConfigs,omitempty" xml:"NotifyConfigs,omitempty"`
	// example:
	//
	// [ "012345" ]
	ParticipantIdsShrink *string `json:"ParticipantIds,omitempty" xml:"ParticipantIds,omitempty"`
	// example:
	//
	// 1703750708595
	ReminderTimeStamp *int64 `json:"ReminderTimeStamp,omitempty" xml:"ReminderTimeStamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待办标题
	Subject             *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CreatePersonalTodoTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetExecutorIdsShrink() *string {
	return s.ExecutorIdsShrink
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetNotifyConfigsShrink() *string {
	return s.NotifyConfigsShrink
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetParticipantIdsShrink() *string {
	return s.ParticipantIdsShrink
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetReminderTimeStamp() *int64 {
	return s.ReminderTimeStamp
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreatePersonalTodoTaskShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetDescription(v string) *CreatePersonalTodoTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetDueTime(v int64) *CreatePersonalTodoTaskShrinkRequest {
	s.DueTime = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetExecutorIdsShrink(v string) *CreatePersonalTodoTaskShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetNotifyConfigsShrink(v string) *CreatePersonalTodoTaskShrinkRequest {
	s.NotifyConfigsShrink = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetParticipantIdsShrink(v string) *CreatePersonalTodoTaskShrinkRequest {
	s.ParticipantIdsShrink = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetReminderTimeStamp(v int64) *CreatePersonalTodoTaskShrinkRequest {
	s.ReminderTimeStamp = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetSubject(v string) *CreatePersonalTodoTaskShrinkRequest {
	s.Subject = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) SetTenantContextShrink(v string) *CreatePersonalTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

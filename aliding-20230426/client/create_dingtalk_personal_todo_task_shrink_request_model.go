// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDingtalkPersonalTodoTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetDescription() *string
	SetDueTime(v int64) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetDueTime() *int64
	SetExecutorIdsShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetExecutorIdsShrink() *string
	SetNotifyConfigsShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetNotifyConfigsShrink() *string
	SetParticipantIdsShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetParticipantIdsShrink() *string
	SetSubject(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetSubject() *string
	SetTenantContextShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetTenantContextShrink() *string
	SetUserToken(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest
	GetUserToken() *string
}

type CreateDingtalkPersonalTodoTaskShrinkRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 待办标题
	Subject             *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 用户token
	UserToken *string `json:"UserToken,omitempty" xml:"UserToken,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetExecutorIdsShrink() *string {
	return s.ExecutorIdsShrink
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetNotifyConfigsShrink() *string {
	return s.NotifyConfigsShrink
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetParticipantIdsShrink() *string {
	return s.ParticipantIdsShrink
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) GetUserToken() *string {
	return s.UserToken
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetDescription(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetDueTime(v int64) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.DueTime = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetExecutorIdsShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetNotifyConfigsShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.NotifyConfigsShrink = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetParticipantIdsShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.ParticipantIdsShrink = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetSubject(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.Subject = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetTenantContextShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) SetUserToken(v string) *CreateDingtalkPersonalTodoTaskShrinkRequest {
	s.UserToken = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *UpdateTodoTaskShrinkRequest
	GetTenantContextShrink() *string
	SetDescription(v string) *UpdateTodoTaskShrinkRequest
	GetDescription() *string
	SetDone(v bool) *UpdateTodoTaskShrinkRequest
	GetDone() *bool
	SetDueTime(v int64) *UpdateTodoTaskShrinkRequest
	GetDueTime() *int64
	SetExecutorIdsShrink(v string) *UpdateTodoTaskShrinkRequest
	GetExecutorIdsShrink() *string
	SetParticipantIdsShrink(v string) *UpdateTodoTaskShrinkRequest
	GetParticipantIdsShrink() *string
	SetSubject(v string) *UpdateTodoTaskShrinkRequest
	GetSubject() *string
	SetTaskId(v string) *UpdateTodoTaskShrinkRequest
	GetTaskId() *string
}

type UpdateTodoTaskShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 应用可以调用该接口发起一个钉钉待办任务，该待办事项会出现在钉钉客户端“待办”页面，需要注意的是，通过开放接口发起的待办，目前仅支持直接跳转ISV应用详情页（ISV在调该接口时需传入自身应用详情页链接）。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// example:
	//
	// 1617675000000
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// example:
	//
	// []
	ExecutorIdsShrink *string `json:"executorIds,omitempty" xml:"executorIds,omitempty"`
	// example:
	//
	// []
	ParticipantIdsShrink *string `json:"participantIds,omitempty" xml:"participantIds,omitempty"`
	// example:
	//
	// 更新钉钉待办
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateTodoTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTodoTaskShrinkRequest) GetDone() *bool {
	return s.Done
}

func (s *UpdateTodoTaskShrinkRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *UpdateTodoTaskShrinkRequest) GetExecutorIdsShrink() *string {
	return s.ExecutorIdsShrink
}

func (s *UpdateTodoTaskShrinkRequest) GetParticipantIdsShrink() *string {
	return s.ParticipantIdsShrink
}

func (s *UpdateTodoTaskShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *UpdateTodoTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTodoTaskShrinkRequest) SetTenantContextShrink(v string) *UpdateTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetDescription(v string) *UpdateTodoTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetDone(v bool) *UpdateTodoTaskShrinkRequest {
	s.Done = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetDueTime(v int64) *UpdateTodoTaskShrinkRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetExecutorIdsShrink(v string) *UpdateTodoTaskShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetParticipantIdsShrink(v string) *UpdateTodoTaskShrinkRequest {
	s.ParticipantIdsShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetSubject(v string) *UpdateTodoTaskShrinkRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetTaskId(v string) *UpdateTodoTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskExecutorStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *UpdateTodoTaskExecutorStatusShrinkRequest
	GetTenantContextShrink() *string
	SetExecutorStatusListShrink(v string) *UpdateTodoTaskExecutorStatusShrinkRequest
	GetExecutorStatusListShrink() *string
	SetOperatorId(v string) *UpdateTodoTaskExecutorStatusShrinkRequest
	GetOperatorId() *string
	SetTaskId(v string) *UpdateTodoTaskExecutorStatusShrinkRequest
	GetTaskId() *string
}

type UpdateTodoTaskExecutorStatusShrinkRequest struct {
	TenantContextShrink      *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	ExecutorStatusListShrink *string `json:"executorStatusList,omitempty" xml:"executorStatusList,omitempty"`
	// example:
	//
	// xxxx
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15002141
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) GetExecutorStatusListShrink() *string {
	return s.ExecutorStatusListShrink
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetTenantContextShrink(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetExecutorStatusListShrink(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.ExecutorStatusListShrink = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetOperatorId(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetTaskId(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}

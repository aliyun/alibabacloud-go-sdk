// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTodoTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *DeleteTodoTaskShrinkRequest
	GetTenantContextShrink() *string
	SetOperatorId(v string) *DeleteTodoTaskShrinkRequest
	GetOperatorId() *string
	SetTaskId(v string) *DeleteTodoTaskShrinkRequest
	GetTaskId() *string
}

type DeleteTodoTaskShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	OperatorId          *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteTodoTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteTodoTaskShrinkRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *DeleteTodoTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteTodoTaskShrinkRequest) SetTenantContextShrink(v string) *DeleteTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteTodoTaskShrinkRequest) SetOperatorId(v string) *DeleteTodoTaskShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteTodoTaskShrinkRequest) SetTaskId(v string) *DeleteTodoTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteTodoTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

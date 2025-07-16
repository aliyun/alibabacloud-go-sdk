// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTodoTaskShrinkRequest
	GetTaskId() *string
	SetTenantContextShrink(v string) *GetTodoTaskShrinkRequest
	GetTenantContextShrink() *string
}

type GetTodoTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// OPJpwtxxxx
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetTodoTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTodoTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTodoTaskShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetTodoTaskShrinkRequest) SetTaskId(v string) *GetTodoTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *GetTodoTaskShrinkRequest) SetTenantContextShrink(v string) *GetTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetTodoTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

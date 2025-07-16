// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTodoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *DeleteTodoTaskRequestTenantContext) *DeleteTodoTaskRequest
	GetTenantContext() *DeleteTodoTaskRequestTenantContext
	SetOperatorId(v string) *DeleteTodoTaskRequest
	GetOperatorId() *string
	SetTaskId(v string) *DeleteTodoTaskRequest
	GetTaskId() *string
}

type DeleteTodoTaskRequest struct {
	TenantContext *DeleteTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// 12345
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 63edc8da7e917d6ecdaab11b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteTodoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskRequest) GetTenantContext() *DeleteTodoTaskRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteTodoTaskRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *DeleteTodoTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteTodoTaskRequest) SetTenantContext(v *DeleteTodoTaskRequestTenantContext) *DeleteTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteTodoTaskRequest) SetOperatorId(v string) *DeleteTodoTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteTodoTaskRequest) SetTaskId(v string) *DeleteTodoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteTodoTaskRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteTodoTaskRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteTodoTaskRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteTodoTaskRequestTenantContext) SetTenantId(v string) *DeleteTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteTodoTaskRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

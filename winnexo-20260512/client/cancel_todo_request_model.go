// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *CancelTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *CancelTodoRequest
	GetTodoId() *string
	SetWorkItemId(v string) *CancelTodoRequest
	GetWorkItemId() *string
}

type CancelTodoRequest struct {
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 待办ID（与 workItemId 二选一）
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
	// 工作项ID（按来源批量作废；与 todoId 二选一）
	//
	// example:
	//
	// wi_1
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s CancelTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTodoRequest) GoString() string {
	return s.String()
}

func (s *CancelTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CancelTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *CancelTodoRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *CancelTodoRequest) SetTenantId(v string) *CancelTodoRequest {
	s.TenantId = &v
	return s
}

func (s *CancelTodoRequest) SetTodoId(v string) *CancelTodoRequest {
	s.TodoId = &v
	return s
}

func (s *CancelTodoRequest) SetWorkItemId(v string) *CancelTodoRequest {
	s.WorkItemId = &v
	return s
}

func (s *CancelTodoRequest) Validate() error {
	return dara.Validate(s)
}

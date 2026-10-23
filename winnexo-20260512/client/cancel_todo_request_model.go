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
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do item. You must specify either this parameter or workItemId.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
	// The ID of the work item. You can use this parameter to cancel to-do items in batches by source. You must specify either this parameter or todoId.
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

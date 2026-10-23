// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetTodoDetailRequest
	GetTenantId() *string
	SetTodoId(v string) *GetTodoDetailRequest
	GetTodoId() *string
}

type GetTodoDetailRequest struct {
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do item.
	//
	// This parameter is required.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s GetTodoDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTodoDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTodoDetailRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *GetTodoDetailRequest) SetTenantId(v string) *GetTodoDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetTodoDetailRequest) SetTodoId(v string) *GetTodoDetailRequest {
	s.TodoId = &v
	return s
}

func (s *GetTodoDetailRequest) Validate() error {
	return dara.Validate(s)
}

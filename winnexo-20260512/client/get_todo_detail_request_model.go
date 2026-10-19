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
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 待办ID
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

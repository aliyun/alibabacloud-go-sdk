// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *BatchCancelTodosRequest
	GetTenantId() *string
	SetTodoIds(v []*string) *BatchCancelTodosRequest
	GetTodoIds() []*string
}

type BatchCancelTodosRequest struct {
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 待办ID列表（1~100）
	//
	// This parameter is required.
	//
	// example:
	//
	// todo_abc123
	TodoIds []*string `json:"todoIds,omitempty" xml:"todoIds,omitempty" type:"Repeated"`
}

func (s BatchCancelTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTodosRequest) GoString() string {
	return s.String()
}

func (s *BatchCancelTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchCancelTodosRequest) GetTodoIds() []*string {
	return s.TodoIds
}

func (s *BatchCancelTodosRequest) SetTenantId(v string) *BatchCancelTodosRequest {
	s.TenantId = &v
	return s
}

func (s *BatchCancelTodosRequest) SetTodoIds(v []*string) *BatchCancelTodosRequest {
	s.TodoIds = v
	return s
}

func (s *BatchCancelTodosRequest) Validate() error {
	return dara.Validate(s)
}

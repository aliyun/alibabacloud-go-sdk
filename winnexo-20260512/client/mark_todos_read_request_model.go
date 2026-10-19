// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkTodosReadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarkAllRead(v bool) *MarkTodosReadRequest
	GetMarkAllRead() *bool
	SetTenantId(v string) *MarkTodosReadRequest
	GetTenantId() *string
	SetTodoIds(v []*string) *MarkTodosReadRequest
	GetTodoIds() []*string
}

type MarkTodosReadRequest struct {
	// 全部标记已读（优先于 todoIds）
	//
	// example:
	//
	// false
	MarkAllRead *bool `json:"markAllRead,omitempty" xml:"markAllRead,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 待办ID列表（≤100；markAllRead=true 时忽略）
	//
	// example:
	//
	// todo_abc123
	TodoIds []*string `json:"todoIds,omitempty" xml:"todoIds,omitempty" type:"Repeated"`
}

func (s MarkTodosReadRequest) String() string {
	return dara.Prettify(s)
}

func (s MarkTodosReadRequest) GoString() string {
	return s.String()
}

func (s *MarkTodosReadRequest) GetMarkAllRead() *bool {
	return s.MarkAllRead
}

func (s *MarkTodosReadRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MarkTodosReadRequest) GetTodoIds() []*string {
	return s.TodoIds
}

func (s *MarkTodosReadRequest) SetMarkAllRead(v bool) *MarkTodosReadRequest {
	s.MarkAllRead = &v
	return s
}

func (s *MarkTodosReadRequest) SetTenantId(v string) *MarkTodosReadRequest {
	s.TenantId = &v
	return s
}

func (s *MarkTodosReadRequest) SetTodoIds(v []*string) *MarkTodosReadRequest {
	s.TodoIds = v
	return s
}

func (s *MarkTodosReadRequest) Validate() error {
	return dara.Validate(s)
}

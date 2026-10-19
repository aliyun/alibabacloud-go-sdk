// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkTodosReadShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarkAllRead(v bool) *MarkTodosReadShrinkRequest
	GetMarkAllRead() *bool
	SetTenantId(v string) *MarkTodosReadShrinkRequest
	GetTenantId() *string
	SetTodoIdsShrink(v string) *MarkTodosReadShrinkRequest
	GetTodoIdsShrink() *string
}

type MarkTodosReadShrinkRequest struct {
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
	TodoIdsShrink *string `json:"todoIds,omitempty" xml:"todoIds,omitempty"`
}

func (s MarkTodosReadShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MarkTodosReadShrinkRequest) GoString() string {
	return s.String()
}

func (s *MarkTodosReadShrinkRequest) GetMarkAllRead() *bool {
	return s.MarkAllRead
}

func (s *MarkTodosReadShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MarkTodosReadShrinkRequest) GetTodoIdsShrink() *string {
	return s.TodoIdsShrink
}

func (s *MarkTodosReadShrinkRequest) SetMarkAllRead(v bool) *MarkTodosReadShrinkRequest {
	s.MarkAllRead = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetTenantId(v string) *MarkTodosReadShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetTodoIdsShrink(v string) *MarkTodosReadShrinkRequest {
	s.TodoIdsShrink = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) Validate() error {
	return dara.Validate(s)
}

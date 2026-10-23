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
	// Specifies whether to mark all to-do items as read. This parameter takes precedence over the todoIds parameter.
	//
	// example:
	//
	// false
	MarkAllRead *bool `json:"markAllRead,omitempty" xml:"markAllRead,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of to-do item IDs. You can specify up to 100 IDs. This parameter is ignored if markAllRead is set to true.
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

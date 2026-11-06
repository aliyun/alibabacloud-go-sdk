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
	// Specifies whether to mark all to-do items as read. This parameter takes precedence over todoIds.
	//
	// example:
	//
	// false
	MarkAllRead *bool `json:"markAllRead,omitempty" xml:"markAllRead,omitempty"`
	// The list of to-do item IDs. A maximum of 100 IDs can be specified. This parameter is ignored when markAllRead is set to true.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of to-do item IDs. A maximum of 100 IDs can be specified. This parameter is ignored when markAllRead is set to true.
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

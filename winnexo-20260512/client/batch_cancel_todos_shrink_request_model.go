// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTodosShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *BatchCancelTodosShrinkRequest
	GetTenantId() *string
	SetTodoIdsShrink(v string) *BatchCancelTodosShrinkRequest
	GetTodoIdsShrink() *string
}

type BatchCancelTodosShrinkRequest struct {
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of to-do IDs. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// todo_abc123
	TodoIdsShrink *string `json:"todoIds,omitempty" xml:"todoIds,omitempty"`
}

func (s BatchCancelTodosShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTodosShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCancelTodosShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchCancelTodosShrinkRequest) GetTodoIdsShrink() *string {
	return s.TodoIdsShrink
}

func (s *BatchCancelTodosShrinkRequest) SetTenantId(v string) *BatchCancelTodosShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) SetTodoIdsShrink(v string) *BatchCancelTodosShrinkRequest {
	s.TodoIdsShrink = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) Validate() error {
	return dara.Validate(s)
}

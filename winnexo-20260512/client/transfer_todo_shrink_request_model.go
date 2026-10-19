// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTodoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *TransferTodoShrinkRequest
	GetComment() *string
	SetTargetUserIdsShrink(v string) *TransferTodoShrinkRequest
	GetTargetUserIdsShrink() *string
	SetTenantId(v string) *TransferTodoShrinkRequest
	GetTenantId() *string
	SetTodoId(v string) *TransferTodoShrinkRequest
	GetTodoId() *string
}

type TransferTodoShrinkRequest struct {
	// 交接说明
	//
	// example:
	//
	// 请继续处理
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 转办目标用户ID列表（1~100）
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	TargetUserIdsShrink *string `json:"targetUserIds,omitempty" xml:"targetUserIds,omitempty"`
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

func (s TransferTodoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTodoShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransferTodoShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *TransferTodoShrinkRequest) GetTargetUserIdsShrink() *string {
	return s.TargetUserIdsShrink
}

func (s *TransferTodoShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TransferTodoShrinkRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *TransferTodoShrinkRequest) SetComment(v string) *TransferTodoShrinkRequest {
	s.Comment = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetTargetUserIdsShrink(v string) *TransferTodoShrinkRequest {
	s.TargetUserIdsShrink = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetTenantId(v string) *TransferTodoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetTodoId(v string) *TransferTodoShrinkRequest {
	s.TodoId = &v
	return s
}

func (s *TransferTodoShrinkRequest) Validate() error {
	return dara.Validate(s)
}

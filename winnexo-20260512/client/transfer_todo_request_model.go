// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *TransferTodoRequest
	GetComment() *string
	SetTargetUserIds(v []*string) *TransferTodoRequest
	GetTargetUserIds() []*string
	SetTenantId(v string) *TransferTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *TransferTodoRequest
	GetTodoId() *string
}

type TransferTodoRequest struct {
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
	TargetUserIds []*string `json:"targetUserIds,omitempty" xml:"targetUserIds,omitempty" type:"Repeated"`
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

func (s TransferTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTodoRequest) GoString() string {
	return s.String()
}

func (s *TransferTodoRequest) GetComment() *string {
	return s.Comment
}

func (s *TransferTodoRequest) GetTargetUserIds() []*string {
	return s.TargetUserIds
}

func (s *TransferTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TransferTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *TransferTodoRequest) SetComment(v string) *TransferTodoRequest {
	s.Comment = &v
	return s
}

func (s *TransferTodoRequest) SetTargetUserIds(v []*string) *TransferTodoRequest {
	s.TargetUserIds = v
	return s
}

func (s *TransferTodoRequest) SetTenantId(v string) *TransferTodoRequest {
	s.TenantId = &v
	return s
}

func (s *TransferTodoRequest) SetTodoId(v string) *TransferTodoRequest {
	s.TodoId = &v
	return s
}

func (s *TransferTodoRequest) Validate() error {
	return dara.Validate(s)
}

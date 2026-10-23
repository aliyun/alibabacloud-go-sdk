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
	// The handover description.
	//
	// example:
	//
	// Please continue processing
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The list of target user IDs to whom the task is reassigned. You can specify 1 to 100 user IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	TargetUserIds []*string `json:"targetUserIds,omitempty" xml:"targetUserIds,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do task.
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

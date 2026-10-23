// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *GetTodoRequest
	GetMessageId() *string
	SetTenantId(v string) *GetTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *GetTodoRequest
	GetTodoId() *string
	SetWorkItemId(v string) *GetTodoRequest
	GetWorkItemId() *string
}

type GetTodoRequest struct {
	// The ID of the assistant message. This parameter is used to query multiple to-do items by message.
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do item (get).
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
	// The ID of the work item. This parameter is used to query multiple to-do items by source.
	//
	// example:
	//
	// wi_1
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTodoRequest) GoString() string {
	return s.String()
}

func (s *GetTodoRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *GetTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *GetTodoRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetTodoRequest) SetMessageId(v string) *GetTodoRequest {
	s.MessageId = &v
	return s
}

func (s *GetTodoRequest) SetTenantId(v string) *GetTodoRequest {
	s.TenantId = &v
	return s
}

func (s *GetTodoRequest) SetTodoId(v string) *GetTodoRequest {
	s.TodoId = &v
	return s
}

func (s *GetTodoRequest) SetWorkItemId(v string) *GetTodoRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetTodoRequest) Validate() error {
	return dara.Validate(s)
}

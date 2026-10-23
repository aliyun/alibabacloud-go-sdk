// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTodoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswersShrink(v string) *SubmitTodoShrinkRequest
	GetAnswersShrink() *string
	SetTenantId(v string) *SubmitTodoShrinkRequest
	GetTenantId() *string
	SetTodoId(v string) *SubmitTodoShrinkRequest
	GetTodoId() *string
}

type SubmitTodoShrinkRequest struct {
	// The array of answers. The original text of the question must match the to-do question.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AnswersShrink *string `json:"answers,omitempty" xml:"answers,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do item.
	//
	// This parameter is required.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s SubmitTodoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTodoShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTodoShrinkRequest) GetAnswersShrink() *string {
	return s.AnswersShrink
}

func (s *SubmitTodoShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SubmitTodoShrinkRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *SubmitTodoShrinkRequest) SetAnswersShrink(v string) *SubmitTodoShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *SubmitTodoShrinkRequest) SetTenantId(v string) *SubmitTodoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitTodoShrinkRequest) SetTodoId(v string) *SubmitTodoShrinkRequest {
	s.TodoId = &v
	return s
}

func (s *SubmitTodoShrinkRequest) Validate() error {
	return dara.Validate(s)
}

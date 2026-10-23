// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v []*SubmitTodoRequestAnswers) *SubmitTodoRequest
	GetAnswers() []*SubmitTodoRequestAnswers
	SetTenantId(v string) *SubmitTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *SubmitTodoRequest
	GetTodoId() *string
}

type SubmitTodoRequest struct {
	// The array of answers. The original text of the question must match the to-do question.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Answers []*SubmitTodoRequestAnswers `json:"answers,omitempty" xml:"answers,omitempty" type:"Repeated"`
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

func (s SubmitTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTodoRequest) GoString() string {
	return s.String()
}

func (s *SubmitTodoRequest) GetAnswers() []*SubmitTodoRequestAnswers {
	return s.Answers
}

func (s *SubmitTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SubmitTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *SubmitTodoRequest) SetAnswers(v []*SubmitTodoRequestAnswers) *SubmitTodoRequest {
	s.Answers = v
	return s
}

func (s *SubmitTodoRequest) SetTenantId(v string) *SubmitTodoRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitTodoRequest) SetTodoId(v string) *SubmitTodoRequest {
	s.TodoId = &v
	return s
}

func (s *SubmitTodoRequest) Validate() error {
	if s.Answers != nil {
		for _, item := range s.Answers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitTodoRequestAnswers struct {
	// The array of answer values.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Answer []*string `json:"answer,omitempty" xml:"answer,omitempty" type:"Repeated"`
	// The original text that corresponds to questions[].question.
	//
	// This parameter is required.
	//
	// example:
	//
	// Whether to approve
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
}

func (s SubmitTodoRequestAnswers) String() string {
	return dara.Prettify(s)
}

func (s SubmitTodoRequestAnswers) GoString() string {
	return s.String()
}

func (s *SubmitTodoRequestAnswers) GetAnswer() []*string {
	return s.Answer
}

func (s *SubmitTodoRequestAnswers) GetQuestion() *string {
	return s.Question
}

func (s *SubmitTodoRequestAnswers) SetAnswer(v []*string) *SubmitTodoRequestAnswers {
	s.Answer = v
	return s
}

func (s *SubmitTodoRequestAnswers) SetQuestion(v string) *SubmitTodoRequestAnswers {
	s.Question = &v
	return s
}

func (s *SubmitTodoRequestAnswers) Validate() error {
	return dara.Validate(s)
}

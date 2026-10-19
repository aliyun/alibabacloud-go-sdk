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
	// 作答数组（question 原文须与待办问题对齐）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Answers []*SubmitTodoRequestAnswers `json:"answers,omitempty" xml:"answers,omitempty" type:"Repeated"`
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
	// 作答值数组
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Answer []*string `json:"answer,omitempty" xml:"answer,omitempty" type:"Repeated"`
	// 对应 questions[].question 原文
	//
	// This parameter is required.
	//
	// example:
	//
	// 是否通过审批
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

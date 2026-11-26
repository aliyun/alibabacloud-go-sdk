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
	SetAppToken(v string) *SubmitTodoRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *SubmitTodoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SubmitTodoRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SubmitTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *SubmitTodoRequest
	GetTodoId() *string
}

type SubmitTodoRequest struct {
	// The answer array. The original text of the question must align with the to-do question.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Answers []*SubmitTodoRequestAnswers `json:"answers,omitempty" xml:"answers,omitempty" type:"Repeated"`
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The to-do task ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The answer array. The original text of the question must align with the to-do question.
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

func (s *SubmitTodoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SubmitTodoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SubmitTodoRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *SubmitTodoRequest) SetAppToken(v string) *SubmitTodoRequest {
	s.AppToken = &v
	return s
}

func (s *SubmitTodoRequest) SetOpsUserCode(v string) *SubmitTodoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SubmitTodoRequest) SetOpsUserId(v string) *SubmitTodoRequest {
	s.OpsUserId = &v
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
	// The original text corresponding to questions[].question.
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

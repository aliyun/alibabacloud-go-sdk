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
	SetAppToken(v string) *SubmitTodoShrinkRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *SubmitTodoShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SubmitTodoShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SubmitTodoShrinkRequest
	GetTenantId() *string
	SetTodoId(v string) *SubmitTodoShrinkRequest
	GetTodoId() *string
}

type SubmitTodoShrinkRequest struct {
	// The answer array. The original text of the question must align with the to-do question.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AnswersShrink *string `json:"answers,omitempty" xml:"answers,omitempty"`
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

func (s SubmitTodoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTodoShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTodoShrinkRequest) GetAnswersShrink() *string {
	return s.AnswersShrink
}

func (s *SubmitTodoShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SubmitTodoShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SubmitTodoShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *SubmitTodoShrinkRequest) SetAppToken(v string) *SubmitTodoShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SubmitTodoShrinkRequest) SetOpsUserCode(v string) *SubmitTodoShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SubmitTodoShrinkRequest) SetOpsUserId(v string) *SubmitTodoShrinkRequest {
	s.OpsUserId = &v
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

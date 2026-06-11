// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommitMsg(v string) *UpdatePromptVersionRequest
	GetCommitMsg() *string
	SetNamespaceId(v string) *UpdatePromptVersionRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *UpdatePromptVersionRequest
	GetPromptKey() *string
	SetTemplate(v string) *UpdatePromptVersionRequest
	GetTemplate() *string
	SetVariables(v string) *UpdatePromptVersionRequest
	GetVariables() *string
}

type UpdatePromptVersionRequest struct {
	// example:
	//
	// 优化回答语气
	CommitMsg *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-qa
	PromptKey *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 你是一个客服助手，请回答：{question}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// [{"name":"question","defaultValue":"Hello"}]
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s UpdatePromptVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePromptVersionRequest) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *UpdatePromptVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdatePromptVersionRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *UpdatePromptVersionRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdatePromptVersionRequest) GetVariables() *string {
	return s.Variables
}

func (s *UpdatePromptVersionRequest) SetCommitMsg(v string) *UpdatePromptVersionRequest {
	s.CommitMsg = &v
	return s
}

func (s *UpdatePromptVersionRequest) SetNamespaceId(v string) *UpdatePromptVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdatePromptVersionRequest) SetPromptKey(v string) *UpdatePromptVersionRequest {
	s.PromptKey = &v
	return s
}

func (s *UpdatePromptVersionRequest) SetTemplate(v string) *UpdatePromptVersionRequest {
	s.Template = &v
	return s
}

func (s *UpdatePromptVersionRequest) SetVariables(v string) *UpdatePromptVersionRequest {
	s.Variables = &v
	return s
}

func (s *UpdatePromptVersionRequest) Validate() error {
	return dara.Validate(s)
}

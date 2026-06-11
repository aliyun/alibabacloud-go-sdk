// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasedOnVersion(v string) *CreatePromptVersionRequest
	GetBasedOnVersion() *string
	SetCommitMsg(v string) *CreatePromptVersionRequest
	GetCommitMsg() *string
	SetNamespaceId(v string) *CreatePromptVersionRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *CreatePromptVersionRequest
	GetPromptKey() *string
	SetTargetVersion(v string) *CreatePromptVersionRequest
	GetTargetVersion() *string
	SetTemplate(v string) *CreatePromptVersionRequest
	GetTemplate() *string
	SetVariables(v string) *CreatePromptVersionRequest
	GetVariables() *string
}

type CreatePromptVersionRequest struct {
	// example:
	//
	// 0.0.1
	BasedOnVersion *string `json:"BasedOnVersion,omitempty" xml:"BasedOnVersion,omitempty"`
	// example:
	//
	// 初始版本
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
	// example:
	//
	// 0.0.2
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	// example:
	//
	// 你是一个客服助手，请回答：{question}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// example:
	//
	// [{"name":"question","defaultValue":"Hello"}]
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreatePromptVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptVersionRequest) GoString() string {
	return s.String()
}

func (s *CreatePromptVersionRequest) GetBasedOnVersion() *string {
	return s.BasedOnVersion
}

func (s *CreatePromptVersionRequest) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *CreatePromptVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreatePromptVersionRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *CreatePromptVersionRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreatePromptVersionRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreatePromptVersionRequest) GetVariables() *string {
	return s.Variables
}

func (s *CreatePromptVersionRequest) SetBasedOnVersion(v string) *CreatePromptVersionRequest {
	s.BasedOnVersion = &v
	return s
}

func (s *CreatePromptVersionRequest) SetCommitMsg(v string) *CreatePromptVersionRequest {
	s.CommitMsg = &v
	return s
}

func (s *CreatePromptVersionRequest) SetNamespaceId(v string) *CreatePromptVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreatePromptVersionRequest) SetPromptKey(v string) *CreatePromptVersionRequest {
	s.PromptKey = &v
	return s
}

func (s *CreatePromptVersionRequest) SetTargetVersion(v string) *CreatePromptVersionRequest {
	s.TargetVersion = &v
	return s
}

func (s *CreatePromptVersionRequest) SetTemplate(v string) *CreatePromptVersionRequest {
	s.Template = &v
	return s
}

func (s *CreatePromptVersionRequest) SetVariables(v string) *CreatePromptVersionRequest {
	s.Variables = &v
	return s
}

func (s *CreatePromptVersionRequest) Validate() error {
	return dara.Validate(s)
}

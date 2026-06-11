// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizTags(v string) *CreatePromptRequest
	GetBizTags() *string
	SetCommitMsg(v string) *CreatePromptRequest
	GetCommitMsg() *string
	SetDescription(v string) *CreatePromptRequest
	GetDescription() *string
	SetNamespaceId(v string) *CreatePromptRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *CreatePromptRequest
	GetPromptKey() *string
	SetTargetVersion(v string) *CreatePromptRequest
	GetTargetVersion() *string
	SetTemplate(v string) *CreatePromptRequest
	GetTemplate() *string
	SetVariables(v string) *CreatePromptRequest
	GetVariables() *string
}

type CreatePromptRequest struct {
	// example:
	//
	// cs,qa,support
	BizTags *string `json:"BizTags,omitempty" xml:"BizTags,omitempty"`
	// example:
	//
	// 初始版本
	CommitMsg *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	// example:
	//
	// 客服问答 Prompt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 0.0.1
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
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

func (s CreatePromptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptRequest) GoString() string {
	return s.String()
}

func (s *CreatePromptRequest) GetBizTags() *string {
	return s.BizTags
}

func (s *CreatePromptRequest) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *CreatePromptRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePromptRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreatePromptRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *CreatePromptRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreatePromptRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreatePromptRequest) GetVariables() *string {
	return s.Variables
}

func (s *CreatePromptRequest) SetBizTags(v string) *CreatePromptRequest {
	s.BizTags = &v
	return s
}

func (s *CreatePromptRequest) SetCommitMsg(v string) *CreatePromptRequest {
	s.CommitMsg = &v
	return s
}

func (s *CreatePromptRequest) SetDescription(v string) *CreatePromptRequest {
	s.Description = &v
	return s
}

func (s *CreatePromptRequest) SetNamespaceId(v string) *CreatePromptRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreatePromptRequest) SetPromptKey(v string) *CreatePromptRequest {
	s.PromptKey = &v
	return s
}

func (s *CreatePromptRequest) SetTargetVersion(v string) *CreatePromptRequest {
	s.TargetVersion = &v
	return s
}

func (s *CreatePromptRequest) SetTemplate(v string) *CreatePromptRequest {
	s.Template = &v
	return s
}

func (s *CreatePromptRequest) SetVariables(v string) *CreatePromptRequest {
	s.Variables = &v
	return s
}

func (s *CreatePromptRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateLlmTemplateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ValidateLlmTemplateConfigRequest
	GetConfig() *string
	SetProviderTemplateId(v string) *ValidateLlmTemplateConfigRequest
	GetProviderTemplateId() *string
}

type ValidateLlmTemplateConfigRequest struct {
	// The JSON string of the LLM configuration to be validated.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"id":"qwen3-max-2026-01-23","name":"Qwen3 Max","contextWindow":200000,"maxTokens":32768}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the provider template.
	//
	// This parameter is required.
	//
	// example:
	//
	// pt-xxx-001
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s ValidateLlmTemplateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateLlmTemplateConfigRequest) GoString() string {
	return s.String()
}

func (s *ValidateLlmTemplateConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ValidateLlmTemplateConfigRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ValidateLlmTemplateConfigRequest) SetConfig(v string) *ValidateLlmTemplateConfigRequest {
	s.Config = &v
	return s
}

func (s *ValidateLlmTemplateConfigRequest) SetProviderTemplateId(v string) *ValidateLlmTemplateConfigRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *ValidateLlmTemplateConfigRequest) Validate() error {
	return dara.Validate(s)
}

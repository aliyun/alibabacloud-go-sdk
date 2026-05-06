// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *CreateModelProviderTemplateRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *CreateModelProviderTemplateRequest
	GetAgentProvider() *string
	SetBizType(v int32) *CreateModelProviderTemplateRequest
	GetBizType() *int32
	SetConfig(v string) *CreateModelProviderTemplateRequest
	GetConfig() *string
	SetDescription(v string) *CreateModelProviderTemplateRequest
	GetDescription() *string
	SetEnableWuyingProxy(v bool) *CreateModelProviderTemplateRequest
	GetEnableWuyingProxy() *bool
	SetModelTemplateId(v string) *CreateModelProviderTemplateRequest
	GetModelTemplateId() *string
	SetName(v string) *CreateModelProviderTemplateRequest
	GetName() *string
	SetProviderName(v string) *CreateModelProviderTemplateRequest
	GetProviderName() *string
	SetProviderType(v string) *CreateModelProviderTemplateRequest
	GetProviderType() *string
}

type CreateModelProviderTemplateRequest struct {
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"api": "openai-completions",
	//
	// 	"apiKey": "sk-xxxx",
	//
	// 	"baseUrl": "https://dashscope.aliyuncs.com/compatible-mode/v1"
	//
	// }
	Config      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	EnableWuyingProxy *bool `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Managed
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
}

func (s CreateModelProviderTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateModelProviderTemplateRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *CreateModelProviderTemplateRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *CreateModelProviderTemplateRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *CreateModelProviderTemplateRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateModelProviderTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelProviderTemplateRequest) GetEnableWuyingProxy() *bool {
	return s.EnableWuyingProxy
}

func (s *CreateModelProviderTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *CreateModelProviderTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelProviderTemplateRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreateModelProviderTemplateRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateModelProviderTemplateRequest) SetAgentPlatform(v string) *CreateModelProviderTemplateRequest {
	s.AgentPlatform = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetAgentProvider(v string) *CreateModelProviderTemplateRequest {
	s.AgentProvider = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetBizType(v int32) *CreateModelProviderTemplateRequest {
	s.BizType = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetConfig(v string) *CreateModelProviderTemplateRequest {
	s.Config = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetDescription(v string) *CreateModelProviderTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetEnableWuyingProxy(v bool) *CreateModelProviderTemplateRequest {
	s.EnableWuyingProxy = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetModelTemplateId(v string) *CreateModelProviderTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetName(v string) *CreateModelProviderTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetProviderName(v string) *CreateModelProviderTemplateRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) SetProviderType(v string) *CreateModelProviderTemplateRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateModelProviderTemplateRequest) Validate() error {
	return dara.Validate(s)
}

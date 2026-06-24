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
	// Agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// Agent provider name.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// Business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Model provider configuration JSON, containing connection information such as baseUrl, apiKey, and api. The apiKey is encrypted after creation. Not required when ProviderType is WuyingCredit, as it is copied from the system template.
	//
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
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Model provider template description.
	//
	// example:
	//
	// 阿里云百炼服务商
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to enable Wuying security proxy. Must be true when ProviderType is WuyingCredit.
	//
	// example:
	//
	// true
	EnableWuyingProxy *bool `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	// Associated model group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// Model provider template name.
	//
	// example:
	//
	// 阿里云百炼
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Model provider name. Must be unique within the same model template. Naming rules vary by ProviderType. For details, see the ProviderType description.
	//
	// This parameter is required.
	//
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// Model provider type. Different types impose different constraints on ProviderName and Config:
	//
	// - WuyingCredit: Wuying credit package. ProviderName must be wuying-credit. Created by copying from the system template. Config is not required.
	//
	// - Managed: Managed provider. System-reserved names such as wuying-credit cannot be used. Config is required.
	//
	// - Custom: User-defined provider. ProviderName must start with the provider- prefix. Config is required.
	//
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

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
	// The Agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The Agent provider name.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The model provider configuration in JSON format, which contains connection information such as baseUrl, apiKey, and api. The apiKey is encrypted after creation. When ProviderType is set to WuyingCredit, this parameter is not required because the configuration is copied from the system template.
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
	// The description of the model provider template.
	//
	// example:
	//
	// 阿里云百炼服务商
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the WUYING secure proxy. This parameter must be set to true when ProviderType is set to WuyingCredit.
	//
	// example:
	//
	// true
	EnableWuyingProxy *bool `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	// The ID of the associated model template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The name of the model provider template.
	//
	// example:
	//
	// 阿里云百炼
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The model provider name. The name must be unique within the same model template. The naming rules vary based on the value of ProviderType. For more information, see the description of ProviderType.
	//
	// This parameter is required.
	//
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The model provider type. Different types impose different constraints on ProviderName and Config. Valid values:
	//
	// - WuyingCredit: WUYING credit plan. ProviderName must be set to wuying-credit. The template is created by copying from a system template, and Config is not required.
	//
	// - Managed: managed provider. System-reserved names such as wuying-credit cannot be used. Config is required.
	//
	// - Custom: user-defined provider. ProviderName must start with the prefix provider-. Config is required.
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

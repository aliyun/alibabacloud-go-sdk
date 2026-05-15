// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelConnectionInput interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerApiKeys(v []*CreateModelConnectionInputConsumerApiKeys) *CreateModelConnectionInput
	GetConsumerApiKeys() []*CreateModelConnectionInputConsumerApiKeys
	SetDescription(v string) *CreateModelConnectionInput
	GetDescription() *string
	SetModelConnectionName(v string) *CreateModelConnectionInput
	GetModelConnectionName() *string
	SetModelInfoConfigs(v []*ModelInfoConfig) *CreateModelConnectionInput
	GetModelInfoConfigs() []*ModelInfoConfig
	SetProvider(v string) *CreateModelConnectionInput
	GetProvider() *string
	SetProviderSettings(v *ModelConnectionProviderSettings) *CreateModelConnectionInput
	GetProviderSettings() *ModelConnectionProviderSettings
	SetWorkspaceId(v string) *CreateModelConnectionInput
	GetWorkspaceId() *string
}

type CreateModelConnectionInput struct {
	// 要绑定的消费者API密钥列表；空表示开放模式
	ConsumerApiKeys []*CreateModelConnectionInputConsumerApiKeys `json:"consumerApiKeys" xml:"consumerApiKeys" type:"Repeated"`
	// 模型连接的描述信息
	//
	// example:
	//
	// OpenAI GPT-4 connection for production
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 模型连接的唯一名称标识
	//
	// This parameter is required.
	//
	// example:
	//
	// my-openai-connection
	ModelConnectionName *string `json:"modelConnectionName,omitempty" xml:"modelConnectionName,omitempty"`
	// 模型元数据配置列表
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs" xml:"modelInfoConfigs" type:"Repeated"`
	// 模型提供商名称
	//
	// This parameter is required.
	//
	// example:
	//
	// openai
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 模型提供商的配置信息，包括基础URL、模型列表等
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	ProviderSettings *ModelConnectionProviderSettings `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	// 模型连接所属的工作空间标识符
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateModelConnectionInput) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionInput) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionInput) GetConsumerApiKeys() []*CreateModelConnectionInputConsumerApiKeys {
	return s.ConsumerApiKeys
}

func (s *CreateModelConnectionInput) GetDescription() *string {
	return s.Description
}

func (s *CreateModelConnectionInput) GetModelConnectionName() *string {
	return s.ModelConnectionName
}

func (s *CreateModelConnectionInput) GetModelInfoConfigs() []*ModelInfoConfig {
	return s.ModelInfoConfigs
}

func (s *CreateModelConnectionInput) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelConnectionInput) GetProviderSettings() *ModelConnectionProviderSettings {
	return s.ProviderSettings
}

func (s *CreateModelConnectionInput) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateModelConnectionInput) SetConsumerApiKeys(v []*CreateModelConnectionInputConsumerApiKeys) *CreateModelConnectionInput {
	s.ConsumerApiKeys = v
	return s
}

func (s *CreateModelConnectionInput) SetDescription(v string) *CreateModelConnectionInput {
	s.Description = &v
	return s
}

func (s *CreateModelConnectionInput) SetModelConnectionName(v string) *CreateModelConnectionInput {
	s.ModelConnectionName = &v
	return s
}

func (s *CreateModelConnectionInput) SetModelInfoConfigs(v []*ModelInfoConfig) *CreateModelConnectionInput {
	s.ModelInfoConfigs = v
	return s
}

func (s *CreateModelConnectionInput) SetProvider(v string) *CreateModelConnectionInput {
	s.Provider = &v
	return s
}

func (s *CreateModelConnectionInput) SetProviderSettings(v *ModelConnectionProviderSettings) *CreateModelConnectionInput {
	s.ProviderSettings = v
	return s
}

func (s *CreateModelConnectionInput) SetWorkspaceId(v string) *CreateModelConnectionInput {
	s.WorkspaceId = &v
	return s
}

func (s *CreateModelConnectionInput) Validate() error {
	if s.ConsumerApiKeys != nil {
		for _, item := range s.ConsumerApiKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelInfoConfigs != nil {
		for _, item := range s.ModelInfoConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProviderSettings != nil {
		if err := s.ProviderSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelConnectionInputConsumerApiKeys struct {
	ApiKeyId *string `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateModelConnectionInputConsumerApiKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionInputConsumerApiKeys) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionInputConsumerApiKeys) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *CreateModelConnectionInputConsumerApiKeys) GetValue() *string {
	return s.Value
}

func (s *CreateModelConnectionInputConsumerApiKeys) SetApiKeyId(v string) *CreateModelConnectionInputConsumerApiKeys {
	s.ApiKeyId = &v
	return s
}

func (s *CreateModelConnectionInputConsumerApiKeys) SetValue(v string) *CreateModelConnectionInputConsumerApiKeys {
	s.Value = &v
	return s
}

func (s *CreateModelConnectionInputConsumerApiKeys) Validate() error {
	return dara.Validate(s)
}

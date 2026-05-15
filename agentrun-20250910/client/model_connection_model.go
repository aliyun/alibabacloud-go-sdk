// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConnection interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerApiKeys(v []*ModelConnectionConsumerAPIKey) *ModelConnection
	GetConsumerApiKeys() []*ModelConnectionConsumerAPIKey
	SetCreatedAt(v string) *ModelConnection
	GetCreatedAt() *string
	SetDescription(v string) *ModelConnection
	GetDescription() *string
	SetLastUpdatedAt(v string) *ModelConnection
	GetLastUpdatedAt() *string
	SetModelConnectionId(v string) *ModelConnection
	GetModelConnectionId() *string
	SetModelConnectionName(v string) *ModelConnection
	GetModelConnectionName() *string
	SetModelInfoConfigs(v []*ModelInfoConfig) *ModelConnection
	GetModelInfoConfigs() []*ModelInfoConfig
	SetProvider(v string) *ModelConnection
	GetProvider() *string
	SetProviderSettings(v *ModelConnectionProviderSettings) *ModelConnection
	GetProviderSettings() *ModelConnectionProviderSettings
	SetWorkspaceId(v string) *ModelConnection
	GetWorkspaceId() *string
}

type ModelConnection struct {
	// 绑定的消费者API密钥列表
	ConsumerApiKeys []*ModelConnectionConsumerAPIKey `json:"consumerApiKeys" xml:"consumerApiKeys" type:"Repeated"`
	// 模型连接的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 模型连接的描述信息
	//
	// example:
	//
	// OpenAI GPT-4 connection for production
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 模型连接最后一次更新的时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// 模型连接的唯一标识符
	//
	// example:
	//
	// mc-1234567890abcdef
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// 模型连接的唯一名称标识
	//
	// example:
	//
	// my-openai-connection
	ModelConnectionName *string `json:"modelConnectionName,omitempty" xml:"modelConnectionName,omitempty"`
	// 模型元数据配置列表，包含各个模型的功能特性和参数规则
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs" xml:"modelInfoConfigs" type:"Repeated"`
	// 模型提供商名称
	//
	// example:
	//
	// openai
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 模型提供商的配置信息
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

func (s ModelConnection) String() string {
	return dara.Prettify(s)
}

func (s ModelConnection) GoString() string {
	return s.String()
}

func (s *ModelConnection) GetConsumerApiKeys() []*ModelConnectionConsumerAPIKey {
	return s.ConsumerApiKeys
}

func (s *ModelConnection) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ModelConnection) GetDescription() *string {
	return s.Description
}

func (s *ModelConnection) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *ModelConnection) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *ModelConnection) GetModelConnectionName() *string {
	return s.ModelConnectionName
}

func (s *ModelConnection) GetModelInfoConfigs() []*ModelInfoConfig {
	return s.ModelInfoConfigs
}

func (s *ModelConnection) GetProvider() *string {
	return s.Provider
}

func (s *ModelConnection) GetProviderSettings() *ModelConnectionProviderSettings {
	return s.ProviderSettings
}

func (s *ModelConnection) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModelConnection) SetConsumerApiKeys(v []*ModelConnectionConsumerAPIKey) *ModelConnection {
	s.ConsumerApiKeys = v
	return s
}

func (s *ModelConnection) SetCreatedAt(v string) *ModelConnection {
	s.CreatedAt = &v
	return s
}

func (s *ModelConnection) SetDescription(v string) *ModelConnection {
	s.Description = &v
	return s
}

func (s *ModelConnection) SetLastUpdatedAt(v string) *ModelConnection {
	s.LastUpdatedAt = &v
	return s
}

func (s *ModelConnection) SetModelConnectionId(v string) *ModelConnection {
	s.ModelConnectionId = &v
	return s
}

func (s *ModelConnection) SetModelConnectionName(v string) *ModelConnection {
	s.ModelConnectionName = &v
	return s
}

func (s *ModelConnection) SetModelInfoConfigs(v []*ModelInfoConfig) *ModelConnection {
	s.ModelInfoConfigs = v
	return s
}

func (s *ModelConnection) SetProvider(v string) *ModelConnection {
	s.Provider = &v
	return s
}

func (s *ModelConnection) SetProviderSettings(v *ModelConnectionProviderSettings) *ModelConnection {
	s.ProviderSettings = v
	return s
}

func (s *ModelConnection) SetWorkspaceId(v string) *ModelConnection {
	s.WorkspaceId = &v
	return s
}

func (s *ModelConnection) Validate() error {
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

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
	// A list of consumer API keys bound to this model connection.
	ConsumerApiKeys []*ModelConnectionConsumerAPIKey `json:"consumerApiKeys" xml:"consumerApiKeys" type:"Repeated"`
	// The time the model connection was created, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// A description of the model connection.
	//
	// example:
	//
	// OpenAI GPT-4 connection for production
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time the model connection was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The model connection\\"s unique identifier.
	//
	// example:
	//
	// mc-1234567890abcdef
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// A unique name identifying the model connection.
	//
	// example:
	//
	// my-openai-connection
	ModelConnectionName *string `json:"modelConnectionName,omitempty" xml:"modelConnectionName,omitempty"`
	// A list of model metadata configurations. Each configuration includes the features and parameter rules for a specific model.
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs" xml:"modelInfoConfigs" type:"Repeated"`
	// The model provider\\"s name.
	//
	// example:
	//
	// openai
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The settings for the model provider.
	//
	// example:
	//
	// {}
	ProviderSettings *ModelConnectionProviderSettings `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	// The identifier of the workspace containing the model connection.
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

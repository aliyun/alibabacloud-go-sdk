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
	// A list of consumer API keys for the model connection. If this list is empty, the connection enters open mode.
	ConsumerApiKeys []*CreateModelConnectionInputConsumerApiKeys `json:"consumerApiKeys" xml:"consumerApiKeys" type:"Repeated"`
	// A description of the model connection.
	//
	// example:
	//
	// OpenAI GPT-4 connection for production
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A unique name for the model connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-openai-connection
	ModelConnectionName *string `json:"modelConnectionName,omitempty" xml:"modelConnectionName,omitempty"`
	// A list of model metadata configurations.
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs" xml:"modelInfoConfigs" type:"Repeated"`
	// The model provider name.
	//
	// This parameter is required.
	//
	// example:
	//
	// openai
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// Configuration settings for the model provider, such as the base URL and a list of models.
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	ProviderSettings *ModelConnectionProviderSettings `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	// The ID of the workspace containing the model connection.
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

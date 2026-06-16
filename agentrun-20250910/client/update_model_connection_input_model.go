// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelConnectionInput interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerApiKeys(v []*UpdateModelConnectionInputConsumerApiKeys) *UpdateModelConnectionInput
	GetConsumerApiKeys() []*UpdateModelConnectionInputConsumerApiKeys
	SetDescription(v string) *UpdateModelConnectionInput
	GetDescription() *string
	SetModelInfoConfigs(v []*ModelInfoConfig) *UpdateModelConnectionInput
	GetModelInfoConfigs() []*ModelInfoConfig
	SetProviderSettings(v *ModelConnectionProviderSettings) *UpdateModelConnectionInput
	GetProviderSettings() *ModelConnectionProviderSettings
}

type UpdateModelConnectionInput struct {
	// A list of consumer API keys to associate with the model connection.
	ConsumerApiKeys []*UpdateModelConnectionInputConsumerApiKeys `json:"consumerApiKeys" xml:"consumerApiKeys" type:"Repeated"`
	// A new description for the model connection.
	//
	// example:
	//
	// Updated connection description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A new list of model metadata configurations.
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs" xml:"modelInfoConfigs" type:"Repeated"`
	// A new configuration for the model provider.
	//
	// example:
	//
	// {}
	ProviderSettings *ModelConnectionProviderSettings `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
}

func (s UpdateModelConnectionInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionInput) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionInput) GetConsumerApiKeys() []*UpdateModelConnectionInputConsumerApiKeys {
	return s.ConsumerApiKeys
}

func (s *UpdateModelConnectionInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelConnectionInput) GetModelInfoConfigs() []*ModelInfoConfig {
	return s.ModelInfoConfigs
}

func (s *UpdateModelConnectionInput) GetProviderSettings() *ModelConnectionProviderSettings {
	return s.ProviderSettings
}

func (s *UpdateModelConnectionInput) SetConsumerApiKeys(v []*UpdateModelConnectionInputConsumerApiKeys) *UpdateModelConnectionInput {
	s.ConsumerApiKeys = v
	return s
}

func (s *UpdateModelConnectionInput) SetDescription(v string) *UpdateModelConnectionInput {
	s.Description = &v
	return s
}

func (s *UpdateModelConnectionInput) SetModelInfoConfigs(v []*ModelInfoConfig) *UpdateModelConnectionInput {
	s.ModelInfoConfigs = v
	return s
}

func (s *UpdateModelConnectionInput) SetProviderSettings(v *ModelConnectionProviderSettings) *UpdateModelConnectionInput {
	s.ProviderSettings = v
	return s
}

func (s *UpdateModelConnectionInput) Validate() error {
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

type UpdateModelConnectionInputConsumerApiKeys struct {
	ApiKeyId *string `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateModelConnectionInputConsumerApiKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionInputConsumerApiKeys) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionInputConsumerApiKeys) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UpdateModelConnectionInputConsumerApiKeys) GetValue() *string {
	return s.Value
}

func (s *UpdateModelConnectionInputConsumerApiKeys) SetApiKeyId(v string) *UpdateModelConnectionInputConsumerApiKeys {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateModelConnectionInputConsumerApiKeys) SetValue(v string) *UpdateModelConnectionInputConsumerApiKeys {
	s.Value = &v
	return s
}

func (s *UpdateModelConnectionInputConsumerApiKeys) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *CreateModelServiceInput
	GetCredentialName() *string
	SetDescription(v string) *CreateModelServiceInput
	GetDescription() *string
	SetModelInfoConfigs(v []*ModelInfoConfig) *CreateModelServiceInput
	GetModelInfoConfigs() []*ModelInfoConfig
	SetModelServiceName(v string) *CreateModelServiceInput
	GetModelServiceName() *string
	SetModelType(v string) *CreateModelServiceInput
	GetModelType() *string
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateModelServiceInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetProvider(v string) *CreateModelServiceInput
	GetProvider() *string
	SetProviderSettings(v *ProviderSettings) *CreateModelServiceInput
	GetProviderSettings() *ProviderSettings
}

type CreateModelServiceInput struct {
	CredentialName   *string            `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description      *string            `json:"description,omitempty" xml:"description,omitempty"`
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs,omitempty" xml:"modelInfoConfigs,omitempty" type:"Repeated"`
	// This parameter is required.
	ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
	// This parameter is required.
	ModelType            *string               `json:"modelType,omitempty" xml:"modelType,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// This parameter is required.
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// This parameter is required.
	ProviderSettings *ProviderSettings `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
}

func (s CreateModelServiceInput) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceInput) GoString() string {
	return s.String()
}

func (s *CreateModelServiceInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateModelServiceInput) GetDescription() *string {
	return s.Description
}

func (s *CreateModelServiceInput) GetModelInfoConfigs() []*ModelInfoConfig {
	return s.ModelInfoConfigs
}

func (s *CreateModelServiceInput) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *CreateModelServiceInput) GetModelType() *string {
	return s.ModelType
}

func (s *CreateModelServiceInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateModelServiceInput) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelServiceInput) GetProviderSettings() *ProviderSettings {
	return s.ProviderSettings
}

func (s *CreateModelServiceInput) SetCredentialName(v string) *CreateModelServiceInput {
	s.CredentialName = &v
	return s
}

func (s *CreateModelServiceInput) SetDescription(v string) *CreateModelServiceInput {
	s.Description = &v
	return s
}

func (s *CreateModelServiceInput) SetModelInfoConfigs(v []*ModelInfoConfig) *CreateModelServiceInput {
	s.ModelInfoConfigs = v
	return s
}

func (s *CreateModelServiceInput) SetModelServiceName(v string) *CreateModelServiceInput {
	s.ModelServiceName = &v
	return s
}

func (s *CreateModelServiceInput) SetModelType(v string) *CreateModelServiceInput {
	s.ModelType = &v
	return s
}

func (s *CreateModelServiceInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateModelServiceInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateModelServiceInput) SetProvider(v string) *CreateModelServiceInput {
	s.Provider = &v
	return s
}

func (s *CreateModelServiceInput) SetProviderSettings(v *ProviderSettings) *CreateModelServiceInput {
	s.ProviderSettings = v
	return s
}

func (s *CreateModelServiceInput) Validate() error {
	if s.ModelInfoConfigs != nil {
		for _, item := range s.ModelInfoConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ProviderSettings != nil {
		if err := s.ProviderSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

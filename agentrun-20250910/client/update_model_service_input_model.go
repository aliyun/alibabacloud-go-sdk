// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelServiceInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *UpdateModelServiceInput
	GetCredentialName() *string
	SetDescription(v string) *UpdateModelServiceInput
	GetDescription() *string
	SetModelInfoConfigs(v []*ModelInfoConfig) *UpdateModelServiceInput
	GetModelInfoConfigs() []*ModelInfoConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateModelServiceInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetProviderSettings(v *ProviderSettings) *UpdateModelServiceInput
	GetProviderSettings() *ProviderSettings
	SetStatus(v string) *UpdateModelServiceInput
	GetStatus() *string
	SetStatusReason(v string) *UpdateModelServiceInput
	GetStatusReason() *string
}

type UpdateModelServiceInput struct {
	CredentialName       *string               `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	ModelInfoConfigs     []*ModelInfoConfig    `json:"modelInfoConfigs,omitempty" xml:"modelInfoConfigs,omitempty" type:"Repeated"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	ProviderSettings     *ProviderSettings     `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	Status               *string               `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason         *string               `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
}

func (s UpdateModelServiceInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelServiceInput) GoString() string {
	return s.String()
}

func (s *UpdateModelServiceInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *UpdateModelServiceInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelServiceInput) GetModelInfoConfigs() []*ModelInfoConfig {
	return s.ModelInfoConfigs
}

func (s *UpdateModelServiceInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateModelServiceInput) GetProviderSettings() *ProviderSettings {
	return s.ProviderSettings
}

func (s *UpdateModelServiceInput) GetStatus() *string {
	return s.Status
}

func (s *UpdateModelServiceInput) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateModelServiceInput) SetCredentialName(v string) *UpdateModelServiceInput {
	s.CredentialName = &v
	return s
}

func (s *UpdateModelServiceInput) SetDescription(v string) *UpdateModelServiceInput {
	s.Description = &v
	return s
}

func (s *UpdateModelServiceInput) SetModelInfoConfigs(v []*ModelInfoConfig) *UpdateModelServiceInput {
	s.ModelInfoConfigs = v
	return s
}

func (s *UpdateModelServiceInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateModelServiceInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateModelServiceInput) SetProviderSettings(v *ProviderSettings) *UpdateModelServiceInput {
	s.ProviderSettings = v
	return s
}

func (s *UpdateModelServiceInput) SetStatus(v string) *UpdateModelServiceInput {
	s.Status = &v
	return s
}

func (s *UpdateModelServiceInput) SetStatusReason(v string) *UpdateModelServiceInput {
	s.StatusReason = &v
	return s
}

func (s *UpdateModelServiceInput) Validate() error {
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

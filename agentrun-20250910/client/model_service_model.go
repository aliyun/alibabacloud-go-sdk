// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelService interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *ModelService
	GetCreatedAt() *string
	SetCredentialName(v string) *ModelService
	GetCredentialName() *string
	SetDescription(v string) *ModelService
	GetDescription() *string
	SetLastUpdatedAt(v string) *ModelService
	GetLastUpdatedAt() *string
	SetModelInfoConfigs(v []*ModelInfoConfig) *ModelService
	GetModelInfoConfigs() []*ModelInfoConfig
	SetModelServiceId(v string) *ModelService
	GetModelServiceId() *string
	SetModelServiceName(v string) *ModelService
	GetModelServiceName() *string
	SetModelType(v string) *ModelService
	GetModelType() *string
	SetNetworkConfiguration(v *NetworkConfiguration) *ModelService
	GetNetworkConfiguration() *NetworkConfiguration
	SetProvider(v string) *ModelService
	GetProvider() *string
	SetProviderSettings(v *ProviderSettings) *ModelService
	GetProviderSettings() *ProviderSettings
	SetStatus(v string) *ModelService
	GetStatus() *string
	SetStatusReason(v string) *ModelService
	GetStatusReason() *string
}

type ModelService struct {
	CreatedAt            *string               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialName       *string               `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	LastUpdatedAt        *string               `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	ModelInfoConfigs     []*ModelInfoConfig    `json:"modelInfoConfigs,omitempty" xml:"modelInfoConfigs,omitempty" type:"Repeated"`
	ModelServiceId       *string               `json:"modelServiceId,omitempty" xml:"modelServiceId,omitempty"`
	ModelServiceName     *string               `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
	ModelType            *string               `json:"modelType,omitempty" xml:"modelType,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	Provider             *string               `json:"provider,omitempty" xml:"provider,omitempty"`
	ProviderSettings     *ProviderSettings     `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	Status               *string               `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason         *string               `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
}

func (s ModelService) String() string {
	return dara.Prettify(s)
}

func (s ModelService) GoString() string {
	return s.String()
}

func (s *ModelService) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ModelService) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ModelService) GetDescription() *string {
	return s.Description
}

func (s *ModelService) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *ModelService) GetModelInfoConfigs() []*ModelInfoConfig {
	return s.ModelInfoConfigs
}

func (s *ModelService) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *ModelService) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *ModelService) GetModelType() *string {
	return s.ModelType
}

func (s *ModelService) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *ModelService) GetProvider() *string {
	return s.Provider
}

func (s *ModelService) GetProviderSettings() *ProviderSettings {
	return s.ProviderSettings
}

func (s *ModelService) GetStatus() *string {
	return s.Status
}

func (s *ModelService) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ModelService) SetCreatedAt(v string) *ModelService {
	s.CreatedAt = &v
	return s
}

func (s *ModelService) SetCredentialName(v string) *ModelService {
	s.CredentialName = &v
	return s
}

func (s *ModelService) SetDescription(v string) *ModelService {
	s.Description = &v
	return s
}

func (s *ModelService) SetLastUpdatedAt(v string) *ModelService {
	s.LastUpdatedAt = &v
	return s
}

func (s *ModelService) SetModelInfoConfigs(v []*ModelInfoConfig) *ModelService {
	s.ModelInfoConfigs = v
	return s
}

func (s *ModelService) SetModelServiceId(v string) *ModelService {
	s.ModelServiceId = &v
	return s
}

func (s *ModelService) SetModelServiceName(v string) *ModelService {
	s.ModelServiceName = &v
	return s
}

func (s *ModelService) SetModelType(v string) *ModelService {
	s.ModelType = &v
	return s
}

func (s *ModelService) SetNetworkConfiguration(v *NetworkConfiguration) *ModelService {
	s.NetworkConfiguration = v
	return s
}

func (s *ModelService) SetProvider(v string) *ModelService {
	s.Provider = &v
	return s
}

func (s *ModelService) SetProviderSettings(v *ProviderSettings) *ModelService {
	s.ProviderSettings = v
	return s
}

func (s *ModelService) SetStatus(v string) *ModelService {
	s.Status = &v
	return s
}

func (s *ModelService) SetStatusReason(v string) *ModelService {
	s.StatusReason = &v
	return s
}

func (s *ModelService) Validate() error {
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

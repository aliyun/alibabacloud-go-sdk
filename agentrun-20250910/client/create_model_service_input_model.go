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
	SetWorkspaceId(v string) *CreateModelServiceInput
	GetWorkspaceId() *string
}

type CreateModelServiceInput struct {
	// The credential name for authenticating with the cloud provider.
	//
	// example:
	//
	// credentialName
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// An optional description of the model service.
	//
	// example:
	//
	// Auto generate task: Pipeline[pipeline-run-1742178254775] pipelineTemplate[data-export-service-online-iVnQB5] taskTemplate[serverless-runner-task], time[2025-03-17T02:24:36Z]
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A list of configurations for models in the service.
	ModelInfoConfigs []*ModelInfoConfig `json:"modelInfoConfigs" xml:"modelInfoConfigs" type:"Repeated"`
	// The name of the model service.
	//
	// This parameter is required.
	//
	// example:
	//
	// modelServiceName
	ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
	// The model type. Valid values include:
	//
	// - `system`: A built-in model that the service provides.
	//
	// - `deployment`: A custom model that a user deploys.
	//
	// This parameter is required.
	//
	// example:
	//
	// system
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The network configuration for the model service. See `NetworkConfiguration` for details.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// The cloud provider for the model service. Currently, only Alibaba Cloud is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// Provider-specific configuration settings. See `ProviderSettings` for details.
	//
	// This parameter is required.
	ProviderSettings *ProviderSettings `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	// The ID of the workspace in which to create the model service.
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *CreateModelServiceInput) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *CreateModelServiceInput) SetWorkspaceId(v string) *CreateModelServiceInput {
	s.WorkspaceId = &v
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

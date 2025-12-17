// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProxyInput interface {
	dara.Model
	String() string
	GoString() string
	SetArmsConfiguration(v *ArmsConfiguration) *CreateModelProxyInput
	GetArmsConfiguration() *ArmsConfiguration
	SetCpu(v float32) *CreateModelProxyInput
	GetCpu() *float32
	SetCredentialName(v string) *CreateModelProxyInput
	GetCredentialName() *string
	SetDescription(v string) *CreateModelProxyInput
	GetDescription() *string
	SetExecutionRoleArn(v string) *CreateModelProxyInput
	GetExecutionRoleArn() *string
	SetLitellmVersion(v string) *CreateModelProxyInput
	GetLitellmVersion() *string
	SetLogConfiguration(v *LogConfiguration) *CreateModelProxyInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *CreateModelProxyInput
	GetMemory() *int32
	SetModelProxyName(v string) *CreateModelProxyInput
	GetModelProxyName() *string
	SetModelType(v string) *CreateModelProxyInput
	GetModelType() *string
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateModelProxyInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetProxyConfig(v *ProxyConfig) *CreateModelProxyInput
	GetProxyConfig() *ProxyConfig
	SetProxyMode(v string) *CreateModelProxyInput
	GetProxyMode() *string
	SetServiceRegionId(v string) *CreateModelProxyInput
	GetServiceRegionId() *string
}

type CreateModelProxyInput struct {
	ArmsConfiguration *ArmsConfiguration `json:"armsConfiguration,omitempty" xml:"armsConfiguration,omitempty"`
	// This parameter is required.
	Cpu              *float32          `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CredentialName   *string           `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description      *string           `json:"description,omitempty" xml:"description,omitempty"`
	ExecutionRoleArn *string           `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LitellmVersion   *string           `json:"litellmVersion,omitempty" xml:"litellmVersion,omitempty"`
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// This parameter is required.
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// This parameter is required.
	ModelProxyName       *string               `json:"modelProxyName,omitempty" xml:"modelProxyName,omitempty"`
	ModelType            *string               `json:"modelType,omitempty" xml:"modelType,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// This parameter is required.
	ProxyConfig *ProxyConfig `json:"proxyConfig,omitempty" xml:"proxyConfig,omitempty"`
	// This parameter is required.
	ProxyMode       *string `json:"proxyMode,omitempty" xml:"proxyMode,omitempty"`
	ServiceRegionId *string `json:"serviceRegionId,omitempty" xml:"serviceRegionId,omitempty"`
}

func (s CreateModelProxyInput) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProxyInput) GoString() string {
	return s.String()
}

func (s *CreateModelProxyInput) GetArmsConfiguration() *ArmsConfiguration {
	return s.ArmsConfiguration
}

func (s *CreateModelProxyInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateModelProxyInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateModelProxyInput) GetDescription() *string {
	return s.Description
}

func (s *CreateModelProxyInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateModelProxyInput) GetLitellmVersion() *string {
	return s.LitellmVersion
}

func (s *CreateModelProxyInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *CreateModelProxyInput) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateModelProxyInput) GetModelProxyName() *string {
	return s.ModelProxyName
}

func (s *CreateModelProxyInput) GetModelType() *string {
	return s.ModelType
}

func (s *CreateModelProxyInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateModelProxyInput) GetProxyConfig() *ProxyConfig {
	return s.ProxyConfig
}

func (s *CreateModelProxyInput) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *CreateModelProxyInput) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *CreateModelProxyInput) SetArmsConfiguration(v *ArmsConfiguration) *CreateModelProxyInput {
	s.ArmsConfiguration = v
	return s
}

func (s *CreateModelProxyInput) SetCpu(v float32) *CreateModelProxyInput {
	s.Cpu = &v
	return s
}

func (s *CreateModelProxyInput) SetCredentialName(v string) *CreateModelProxyInput {
	s.CredentialName = &v
	return s
}

func (s *CreateModelProxyInput) SetDescription(v string) *CreateModelProxyInput {
	s.Description = &v
	return s
}

func (s *CreateModelProxyInput) SetExecutionRoleArn(v string) *CreateModelProxyInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateModelProxyInput) SetLitellmVersion(v string) *CreateModelProxyInput {
	s.LitellmVersion = &v
	return s
}

func (s *CreateModelProxyInput) SetLogConfiguration(v *LogConfiguration) *CreateModelProxyInput {
	s.LogConfiguration = v
	return s
}

func (s *CreateModelProxyInput) SetMemory(v int32) *CreateModelProxyInput {
	s.Memory = &v
	return s
}

func (s *CreateModelProxyInput) SetModelProxyName(v string) *CreateModelProxyInput {
	s.ModelProxyName = &v
	return s
}

func (s *CreateModelProxyInput) SetModelType(v string) *CreateModelProxyInput {
	s.ModelType = &v
	return s
}

func (s *CreateModelProxyInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateModelProxyInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateModelProxyInput) SetProxyConfig(v *ProxyConfig) *CreateModelProxyInput {
	s.ProxyConfig = v
	return s
}

func (s *CreateModelProxyInput) SetProxyMode(v string) *CreateModelProxyInput {
	s.ProxyMode = &v
	return s
}

func (s *CreateModelProxyInput) SetServiceRegionId(v string) *CreateModelProxyInput {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateModelProxyInput) Validate() error {
	if s.ArmsConfiguration != nil {
		if err := s.ArmsConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ProxyConfig != nil {
		if err := s.ProxyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

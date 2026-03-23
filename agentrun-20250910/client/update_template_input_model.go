// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
	SetAllowAnonymousManage(v bool) *UpdateTemplateInput
	GetAllowAnonymousManage() *bool
	SetArmsConfiguration(v *ArmsConfiguration) *UpdateTemplateInput
	GetArmsConfiguration() *ArmsConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *UpdateTemplateInput
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *UpdateTemplateInput
	GetCpu() *float32
	SetCredentialConfiguration(v *CredentialConfiguration) *UpdateTemplateInput
	GetCredentialConfiguration() *CredentialConfiguration
	SetDescription(v string) *UpdateTemplateInput
	GetDescription() *string
	SetEnableAgent(v bool) *UpdateTemplateInput
	GetEnableAgent() *bool
	SetEnvironmentVariables(v map[string]*string) *UpdateTemplateInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *UpdateTemplateInput
	GetExecutionRoleArn() *string
	SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *UpdateTemplateInput
	GetMemory() *int32
	SetNasConfig(v *NASConfig) *UpdateTemplateInput
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssConfiguration(v []*OssConfiguration) *UpdateTemplateInput
	GetOssConfiguration() []*OssConfiguration
	SetSandboxIdleTimeoutInSeconds(v int32) *UpdateTemplateInput
	GetSandboxIdleTimeoutInSeconds() *int32
	SetSandboxTTLInSeconds(v int32) *UpdateTemplateInput
	GetSandboxTTLInSeconds() *int32
	SetScalingConfig(v *ScalingConfig) *UpdateTemplateInput
	GetScalingConfig() *ScalingConfig
	SetTemplateConfiguration(v map[string]interface{}) *UpdateTemplateInput
	GetTemplateConfiguration() map[string]interface{}
	SetWorkspaceId(v string) *UpdateTemplateInput
	GetWorkspaceId() *string
}

type UpdateTemplateInput struct {
	// if can be null:
	// true
	//
	// example:
	//
	// true
	AllowAnonymousManage *bool              `json:"allowAnonymousManage,omitempty" xml:"allowAnonymousManage,omitempty"`
	ArmsConfiguration    *ArmsConfiguration `json:"armsConfiguration,omitempty" xml:"armsConfiguration,omitempty"`
	// 容器配置（内置的不可改）
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// CPU资源配置（单位：核心）
	//
	// example:
	//
	// 2
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CredentialConfiguration *CredentialConfiguration `json:"credentialConfiguration,omitempty" xml:"credentialConfiguration,omitempty"`
	// example:
	//
	// demo description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	EnableAgent          *bool              `json:"enableAgent,omitempty" xml:"enableAgent,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// example:
	//
	// arn:acs:agentrun:cn-hangzhou:123456789:xxx/test
	ExecutionRoleArn *string           `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// 内存资源配置（单位：MB）
	//
	// example:
	//
	// 4
	Memory               *int32                `json:"memory,omitempty" xml:"memory,omitempty"`
	NasConfig            *NASConfig            `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	OssConfiguration     []*OssConfiguration   `json:"ossConfiguration,omitempty" xml:"ossConfiguration,omitempty" type:"Repeated"`
	// 沙箱空闲超时时间（秒）
	//
	// example:
	//
	// 21600
	SandboxIdleTimeoutInSeconds *int32 `json:"sandboxIdleTimeoutInSeconds,omitempty" xml:"sandboxIdleTimeoutInSeconds,omitempty"`
	// Deprecated
	//
	// 沙箱存活时间（秒）
	//
	// example:
	//
	// 86400
	SandboxTTLInSeconds *int32         `json:"sandboxTTLInSeconds,omitempty" xml:"sandboxTTLInSeconds,omitempty"`
	ScalingConfig       *ScalingConfig `json:"scalingConfig,omitempty" xml:"scalingConfig,omitempty"`
	// 模板配置（灵活的对象结构，根据 templateType 不同而不同）
	TemplateConfiguration map[string]interface{} `json:"templateConfiguration,omitempty" xml:"templateConfiguration,omitempty"`
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateInput) GoString() string {
	return s.String()
}

func (s *UpdateTemplateInput) GetAllowAnonymousManage() *bool {
	return s.AllowAnonymousManage
}

func (s *UpdateTemplateInput) GetArmsConfiguration() *ArmsConfiguration {
	return s.ArmsConfiguration
}

func (s *UpdateTemplateInput) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateTemplateInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *UpdateTemplateInput) GetCredentialConfiguration() *CredentialConfiguration {
	return s.CredentialConfiguration
}

func (s *UpdateTemplateInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateTemplateInput) GetEnableAgent() *bool {
	return s.EnableAgent
}

func (s *UpdateTemplateInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateTemplateInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateTemplateInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateTemplateInput) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateTemplateInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateTemplateInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateTemplateInput) GetOssConfiguration() []*OssConfiguration {
	return s.OssConfiguration
}

func (s *UpdateTemplateInput) GetSandboxIdleTimeoutInSeconds() *int32 {
	return s.SandboxIdleTimeoutInSeconds
}

func (s *UpdateTemplateInput) GetSandboxTTLInSeconds() *int32 {
	return s.SandboxTTLInSeconds
}

func (s *UpdateTemplateInput) GetScalingConfig() *ScalingConfig {
	return s.ScalingConfig
}

func (s *UpdateTemplateInput) GetTemplateConfiguration() map[string]interface{} {
	return s.TemplateConfiguration
}

func (s *UpdateTemplateInput) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateTemplateInput) SetAllowAnonymousManage(v bool) *UpdateTemplateInput {
	s.AllowAnonymousManage = &v
	return s
}

func (s *UpdateTemplateInput) SetArmsConfiguration(v *ArmsConfiguration) *UpdateTemplateInput {
	s.ArmsConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetContainerConfiguration(v *ContainerConfiguration) *UpdateTemplateInput {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetCpu(v float32) *UpdateTemplateInput {
	s.Cpu = &v
	return s
}

func (s *UpdateTemplateInput) SetCredentialConfiguration(v *CredentialConfiguration) *UpdateTemplateInput {
	s.CredentialConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetDescription(v string) *UpdateTemplateInput {
	s.Description = &v
	return s
}

func (s *UpdateTemplateInput) SetEnableAgent(v bool) *UpdateTemplateInput {
	s.EnableAgent = &v
	return s
}

func (s *UpdateTemplateInput) SetEnvironmentVariables(v map[string]*string) *UpdateTemplateInput {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateTemplateInput) SetExecutionRoleArn(v string) *UpdateTemplateInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateTemplateInput) SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput {
	s.LogConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetMemory(v int32) *UpdateTemplateInput {
	s.Memory = &v
	return s
}

func (s *UpdateTemplateInput) SetNasConfig(v *NASConfig) *UpdateTemplateInput {
	s.NasConfig = v
	return s
}

func (s *UpdateTemplateInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetOssConfiguration(v []*OssConfiguration) *UpdateTemplateInput {
	s.OssConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetSandboxIdleTimeoutInSeconds(v int32) *UpdateTemplateInput {
	s.SandboxIdleTimeoutInSeconds = &v
	return s
}

func (s *UpdateTemplateInput) SetSandboxTTLInSeconds(v int32) *UpdateTemplateInput {
	s.SandboxTTLInSeconds = &v
	return s
}

func (s *UpdateTemplateInput) SetScalingConfig(v *ScalingConfig) *UpdateTemplateInput {
	s.ScalingConfig = v
	return s
}

func (s *UpdateTemplateInput) SetTemplateConfiguration(v map[string]interface{}) *UpdateTemplateInput {
	s.TemplateConfiguration = v
	return s
}

func (s *UpdateTemplateInput) SetWorkspaceId(v string) *UpdateTemplateInput {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateTemplateInput) Validate() error {
	if s.ArmsConfiguration != nil {
		if err := s.ArmsConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.CredentialConfiguration != nil {
		if err := s.CredentialConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NasConfig != nil {
		if err := s.NasConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.OssConfiguration != nil {
		for _, item := range s.OssConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScalingConfig != nil {
		if err := s.ScalingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

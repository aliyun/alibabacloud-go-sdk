// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
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
	SetEnvironmentVariables(v map[string]*string) *UpdateTemplateInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *UpdateTemplateInput
	GetExecutionRoleArn() *string
	SetLogConfiguration(v *LogConfiguration) *UpdateTemplateInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *UpdateTemplateInput
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateTemplateInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssConfiguration(v []*OssConfiguration) *UpdateTemplateInput
	GetOssConfiguration() []*OssConfiguration
	SetSandboxIdleTimeoutInSeconds(v int32) *UpdateTemplateInput
	GetSandboxIdleTimeoutInSeconds() *int32
	SetSandboxTTLInSeconds(v int32) *UpdateTemplateInput
	GetSandboxTTLInSeconds() *int32
	SetTemplateConfiguration(v map[string]interface{}) *UpdateTemplateInput
	GetTemplateConfiguration() map[string]interface{}
}

type UpdateTemplateInput struct {
	ArmsConfiguration *ArmsConfiguration `json:"armsConfiguration,omitempty" xml:"armsConfiguration,omitempty"`
	// 容器配置（内置的不可改）
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// CPU资源配置（单位：核心）
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CredentialConfiguration *CredentialConfiguration `json:"credentialConfiguration,omitempty" xml:"credentialConfiguration,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	EnvironmentVariables    map[string]*string       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	ExecutionRoleArn        *string                  `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LogConfiguration        *LogConfiguration        `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// 内存资源配置（单位：MB）
	Memory               *int32                `json:"memory,omitempty" xml:"memory,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	OssConfiguration     []*OssConfiguration   `json:"ossConfiguration,omitempty" xml:"ossConfiguration,omitempty" type:"Repeated"`
	// 沙箱空闲超时时间（秒）
	SandboxIdleTimeoutInSeconds *int32 `json:"sandboxIdleTimeoutInSeconds,omitempty" xml:"sandboxIdleTimeoutInSeconds,omitempty"`
	// Deprecated
	//
	// 沙箱存活时间（秒）
	SandboxTTLInSeconds *int32 `json:"sandboxTTLInSeconds,omitempty" xml:"sandboxTTLInSeconds,omitempty"`
	// 模板配置（灵活的对象结构，根据 templateType 不同而不同）
	TemplateConfiguration map[string]interface{} `json:"templateConfiguration,omitempty" xml:"templateConfiguration,omitempty"`
}

func (s UpdateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateInput) GoString() string {
	return s.String()
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

func (s *UpdateTemplateInput) GetTemplateConfiguration() map[string]interface{} {
	return s.TemplateConfiguration
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

func (s *UpdateTemplateInput) SetTemplateConfiguration(v map[string]interface{}) *UpdateTemplateInput {
	s.TemplateConfiguration = v
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
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
	SetArmsConfiguration(v *ArmsConfiguration) *CreateTemplateInput
	GetArmsConfiguration() *ArmsConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *CreateTemplateInput
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *CreateTemplateInput
	GetCpu() *float32
	SetCredentialConfiguration(v *CredentialConfiguration) *CreateTemplateInput
	GetCredentialConfiguration() *CredentialConfiguration
	SetDescription(v string) *CreateTemplateInput
	GetDescription() *string
	SetDiskSize(v int32) *CreateTemplateInput
	GetDiskSize() *int32
	SetEnvironmentVariables(v map[string]*string) *CreateTemplateInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *CreateTemplateInput
	GetExecutionRoleArn() *string
	SetLogConfiguration(v *LogConfiguration) *CreateTemplateInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *CreateTemplateInput
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateTemplateInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssConfiguration(v []*OssConfiguration) *CreateTemplateInput
	GetOssConfiguration() []*OssConfiguration
	SetSandboxIdleTimeoutInSeconds(v int32) *CreateTemplateInput
	GetSandboxIdleTimeoutInSeconds() *int32
	SetSandboxTTLInSeconds(v int32) *CreateTemplateInput
	GetSandboxTTLInSeconds() *int32
	SetTemplateConfiguration(v map[string]interface{}) *CreateTemplateInput
	GetTemplateConfiguration() map[string]interface{}
	SetTemplateName(v string) *CreateTemplateInput
	GetTemplateName() *string
	SetTemplateType(v string) *CreateTemplateInput
	GetTemplateType() *string
}

type CreateTemplateInput struct {
	ArmsConfiguration *ArmsConfiguration `json:"armsConfiguration,omitempty" xml:"armsConfiguration,omitempty"`
	// 容器配置，只允许基于 Browser/Code Interpreter 基础镜像的 image
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// CPU资源配置（单位：核心）
	//
	// This parameter is required.
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CredentialConfiguration *CredentialConfiguration `json:"credentialConfiguration,omitempty" xml:"credentialConfiguration,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables    map[string]*string       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	ExecutionRoleArn        *string                  `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LogConfiguration        *LogConfiguration        `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// 内存资源配置（单位：MB）
	//
	// This parameter is required.
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// This parameter is required.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	OssConfiguration     []*OssConfiguration   `json:"ossConfiguration,omitempty" xml:"ossConfiguration,omitempty" type:"Repeated"`
	// 沙箱空闲超时时间（秒）
	SandboxIdleTimeoutInSeconds *int32 `json:"sandboxIdleTimeoutInSeconds,omitempty" xml:"sandboxIdleTimeoutInSeconds,omitempty"`
	// 沙箱存活时间（秒）
	SandboxTTLInSeconds *int32 `json:"sandboxTTLInSeconds,omitempty" xml:"sandboxTTLInSeconds,omitempty"`
	// 模板配置（灵活的对象结构，根据 templateType 不同而不同）
	TemplateConfiguration map[string]interface{} `json:"templateConfiguration,omitempty" xml:"templateConfiguration,omitempty"`
	// 模板名称（要求账号唯一的）
	//
	// This parameter is required.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// This parameter is required.
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s CreateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateInput) GoString() string {
	return s.String()
}

func (s *CreateTemplateInput) GetArmsConfiguration() *ArmsConfiguration {
	return s.ArmsConfiguration
}

func (s *CreateTemplateInput) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *CreateTemplateInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateTemplateInput) GetCredentialConfiguration() *CredentialConfiguration {
	return s.CredentialConfiguration
}

func (s *CreateTemplateInput) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateTemplateInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateTemplateInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateTemplateInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *CreateTemplateInput) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateTemplateInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateTemplateInput) GetOssConfiguration() []*OssConfiguration {
	return s.OssConfiguration
}

func (s *CreateTemplateInput) GetSandboxIdleTimeoutInSeconds() *int32 {
	return s.SandboxIdleTimeoutInSeconds
}

func (s *CreateTemplateInput) GetSandboxTTLInSeconds() *int32 {
	return s.SandboxTTLInSeconds
}

func (s *CreateTemplateInput) GetTemplateConfiguration() map[string]interface{} {
	return s.TemplateConfiguration
}

func (s *CreateTemplateInput) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateInput) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateTemplateInput) SetArmsConfiguration(v *ArmsConfiguration) *CreateTemplateInput {
	s.ArmsConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetContainerConfiguration(v *ContainerConfiguration) *CreateTemplateInput {
	s.ContainerConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetCpu(v float32) *CreateTemplateInput {
	s.Cpu = &v
	return s
}

func (s *CreateTemplateInput) SetCredentialConfiguration(v *CredentialConfiguration) *CreateTemplateInput {
	s.CredentialConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetDescription(v string) *CreateTemplateInput {
	s.Description = &v
	return s
}

func (s *CreateTemplateInput) SetDiskSize(v int32) *CreateTemplateInput {
	s.DiskSize = &v
	return s
}

func (s *CreateTemplateInput) SetEnvironmentVariables(v map[string]*string) *CreateTemplateInput {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateTemplateInput) SetExecutionRoleArn(v string) *CreateTemplateInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateTemplateInput) SetLogConfiguration(v *LogConfiguration) *CreateTemplateInput {
	s.LogConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetMemory(v int32) *CreateTemplateInput {
	s.Memory = &v
	return s
}

func (s *CreateTemplateInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateTemplateInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetOssConfiguration(v []*OssConfiguration) *CreateTemplateInput {
	s.OssConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetSandboxIdleTimeoutInSeconds(v int32) *CreateTemplateInput {
	s.SandboxIdleTimeoutInSeconds = &v
	return s
}

func (s *CreateTemplateInput) SetSandboxTTLInSeconds(v int32) *CreateTemplateInput {
	s.SandboxTTLInSeconds = &v
	return s
}

func (s *CreateTemplateInput) SetTemplateConfiguration(v map[string]interface{}) *CreateTemplateInput {
	s.TemplateConfiguration = v
	return s
}

func (s *CreateTemplateInput) SetTemplateName(v string) *CreateTemplateInput {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateInput) SetTemplateType(v string) *CreateTemplateInput {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateInput) Validate() error {
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

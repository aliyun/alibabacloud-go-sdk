// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRuntimeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v float32) *CreateTemplateRuntimeConfig
	GetCpu() *float32
	SetDiskSize(v int32) *CreateTemplateRuntimeConfig
	GetDiskSize() *int32
	SetInternetAccess(v bool) *CreateTemplateRuntimeConfig
	GetInternetAccess() *bool
	SetLogConfig(v *CreateTemplateLogConfig) *CreateTemplateRuntimeConfig
	GetLogConfig() *CreateTemplateLogConfig
	SetMemorySize(v int32) *CreateTemplateRuntimeConfig
	GetMemorySize() *int32
	SetSandboxConfig(v *CreateTemplateSandboxConfig) *CreateTemplateRuntimeConfig
	GetSandboxConfig() *CreateTemplateSandboxConfig
	SetVpcBindings(v []*string) *CreateTemplateRuntimeConfig
	GetVpcBindings() []*string
	SetVpcConfig(v *CreateTemplateVPCConfig) *CreateTemplateRuntimeConfig
	GetVpcConfig() *CreateTemplateVPCConfig
}

type CreateTemplateRuntimeConfig struct {
	// The CPU specifications of the sandbox instance, in cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The disk specifications of the sandbox instance, in MB.
	//
	// example:
	//
	// 10240
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Specifies whether the sandbox instance is allowed to access the Internet.
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// The log delivery configuration of the sandbox instance.
	LogConfig *CreateTemplateLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory specifications of the sandbox instance, in MB.
	//
	// example:
	//
	// 2048
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The image and startup configuration of the sandbox instance.
	SandboxConfig *CreateTemplateSandboxConfig `json:"sandboxConfig,omitempty" xml:"sandboxConfig,omitempty"`
	// A non-empty array configures the set of source VPCs that are allowed to access the sandbox data endpoint. Omitting the field, setting it to null, or providing an empty array indicates that no inbound bindings are configured. This configuration is independent of the egress network vpcConfig and does not modify the API key source policy. A null value is treated as omitted to maintain compatibility with SDK requests where the field is not set. You can update or clear the list by calling UpdateTemplate. The changes take effect after a new version is successfully built, and existing sandboxes continue to use the original configuration.
	//
	// example:
	//
	// ["vpc-examplea","vpc-exampleb"]
	VpcBindings []*string `json:"vpcBindings" xml:"vpcBindings" type:"Repeated"`
	// The VPC configuration of the sandbox instance.
	VpcConfig *CreateTemplateVPCConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s CreateTemplateRuntimeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRuntimeConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateRuntimeConfig) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateTemplateRuntimeConfig) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateTemplateRuntimeConfig) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *CreateTemplateRuntimeConfig) GetLogConfig() *CreateTemplateLogConfig {
	return s.LogConfig
}

func (s *CreateTemplateRuntimeConfig) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *CreateTemplateRuntimeConfig) GetSandboxConfig() *CreateTemplateSandboxConfig {
	return s.SandboxConfig
}

func (s *CreateTemplateRuntimeConfig) GetVpcBindings() []*string {
	return s.VpcBindings
}

func (s *CreateTemplateRuntimeConfig) GetVpcConfig() *CreateTemplateVPCConfig {
	return s.VpcConfig
}

func (s *CreateTemplateRuntimeConfig) SetCpu(v float32) *CreateTemplateRuntimeConfig {
	s.Cpu = &v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetDiskSize(v int32) *CreateTemplateRuntimeConfig {
	s.DiskSize = &v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetInternetAccess(v bool) *CreateTemplateRuntimeConfig {
	s.InternetAccess = &v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetLogConfig(v *CreateTemplateLogConfig) *CreateTemplateRuntimeConfig {
	s.LogConfig = v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetMemorySize(v int32) *CreateTemplateRuntimeConfig {
	s.MemorySize = &v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetSandboxConfig(v *CreateTemplateSandboxConfig) *CreateTemplateRuntimeConfig {
	s.SandboxConfig = v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetVpcBindings(v []*string) *CreateTemplateRuntimeConfig {
	s.VpcBindings = v
	return s
}

func (s *CreateTemplateRuntimeConfig) SetVpcConfig(v *CreateTemplateVPCConfig) *CreateTemplateRuntimeConfig {
	s.VpcConfig = v
	return s
}

func (s *CreateTemplateRuntimeConfig) Validate() error {
	if s.LogConfig != nil {
		if err := s.LogConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SandboxConfig != nil {
		if err := s.SandboxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

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
	SetVpcConfig(v *CreateTemplateVPCConfig) *CreateTemplateRuntimeConfig
	GetVpcConfig() *CreateTemplateVPCConfig
}

type CreateTemplateRuntimeConfig struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 10
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Specifies whether to allow access to the Internet.
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// The log configuration.
	LogConfig *CreateTemplateLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 2048
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The sandbox configuration.
	SandboxConfig *CreateTemplateSandboxConfig `json:"sandboxConfig,omitempty" xml:"sandboxConfig,omitempty"`
	// The VPC configuration.
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

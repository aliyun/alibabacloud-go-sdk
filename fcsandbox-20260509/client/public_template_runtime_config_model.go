// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateRuntimeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v float32) *PublicTemplateRuntimeConfig
	GetCpu() *float32
	SetDiskSize(v int32) *PublicTemplateRuntimeConfig
	GetDiskSize() *int32
	SetInternetAccess(v bool) *PublicTemplateRuntimeConfig
	GetInternetAccess() *bool
	SetLogConfig(v *PublicTemplateLogConfig) *PublicTemplateRuntimeConfig
	GetLogConfig() *PublicTemplateLogConfig
	SetMemorySize(v int32) *PublicTemplateRuntimeConfig
	GetMemorySize() *int32
	SetSandboxConfig(v *PublicTemplateSandboxConfig) *PublicTemplateRuntimeConfig
	GetSandboxConfig() *PublicTemplateSandboxConfig
	SetVpcConfig(v *PublicTemplateVPCConfig) *PublicTemplateRuntimeConfig
	GetVpcConfig() *PublicTemplateVPCConfig
}

type PublicTemplateRuntimeConfig struct {
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
	LogConfig *PublicTemplateLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 2048
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The sandbox configuration.
	SandboxConfig *PublicTemplateSandboxConfig `json:"sandboxConfig,omitempty" xml:"sandboxConfig,omitempty"`
	// The VPC configuration.
	VpcConfig *PublicTemplateVPCConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s PublicTemplateRuntimeConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateRuntimeConfig) GoString() string {
	return s.String()
}

func (s *PublicTemplateRuntimeConfig) GetCpu() *float32 {
	return s.Cpu
}

func (s *PublicTemplateRuntimeConfig) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *PublicTemplateRuntimeConfig) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *PublicTemplateRuntimeConfig) GetLogConfig() *PublicTemplateLogConfig {
	return s.LogConfig
}

func (s *PublicTemplateRuntimeConfig) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *PublicTemplateRuntimeConfig) GetSandboxConfig() *PublicTemplateSandboxConfig {
	return s.SandboxConfig
}

func (s *PublicTemplateRuntimeConfig) GetVpcConfig() *PublicTemplateVPCConfig {
	return s.VpcConfig
}

func (s *PublicTemplateRuntimeConfig) SetCpu(v float32) *PublicTemplateRuntimeConfig {
	s.Cpu = &v
	return s
}

func (s *PublicTemplateRuntimeConfig) SetDiskSize(v int32) *PublicTemplateRuntimeConfig {
	s.DiskSize = &v
	return s
}

func (s *PublicTemplateRuntimeConfig) SetInternetAccess(v bool) *PublicTemplateRuntimeConfig {
	s.InternetAccess = &v
	return s
}

func (s *PublicTemplateRuntimeConfig) SetLogConfig(v *PublicTemplateLogConfig) *PublicTemplateRuntimeConfig {
	s.LogConfig = v
	return s
}

func (s *PublicTemplateRuntimeConfig) SetMemorySize(v int32) *PublicTemplateRuntimeConfig {
	s.MemorySize = &v
	return s
}

func (s *PublicTemplateRuntimeConfig) SetSandboxConfig(v *PublicTemplateSandboxConfig) *PublicTemplateRuntimeConfig {
	s.SandboxConfig = v
	return s
}

func (s *PublicTemplateRuntimeConfig) SetVpcConfig(v *PublicTemplateVPCConfig) *PublicTemplateRuntimeConfig {
	s.VpcConfig = v
	return s
}

func (s *PublicTemplateRuntimeConfig) Validate() error {
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

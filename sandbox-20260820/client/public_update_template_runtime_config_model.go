// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateRuntimeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v float32) *PublicUpdateTemplateRuntimeConfig
	GetCpu() *float32
	SetDiskSize(v int32) *PublicUpdateTemplateRuntimeConfig
	GetDiskSize() *int32
	SetInternetAccess(v bool) *PublicUpdateTemplateRuntimeConfig
	GetInternetAccess() *bool
	SetLogConfig(v *PublicUpdateTemplateLogConfig) *PublicUpdateTemplateRuntimeConfig
	GetLogConfig() *PublicUpdateTemplateLogConfig
	SetMemorySize(v int32) *PublicUpdateTemplateRuntimeConfig
	GetMemorySize() *int32
	SetSandboxConfig(v *PublicUpdateTemplateSandboxConfig) *PublicUpdateTemplateRuntimeConfig
	GetSandboxConfig() *PublicUpdateTemplateSandboxConfig
	SetVpcBindings(v []*string) *PublicUpdateTemplateRuntimeConfig
	GetVpcBindings() []*string
	SetVpcConfig(v *PublicUpdateTemplateVPCConfig) *PublicUpdateTemplateRuntimeConfig
	GetVpcConfig() *PublicUpdateTemplateVPCConfig
}

type PublicUpdateTemplateRuntimeConfig struct {
	// The number of vCPU cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The system cloud disk size in MB.
	//
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Specifies whether the sandbox is allowed to access the Internet.
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// The log configuration.
	LogConfig *PublicUpdateTemplateLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory size in MB.
	//
	// example:
	//
	// 2048
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The sandbox configuration.
	SandboxConfig *PublicUpdateTemplateSandboxConfig `json:"sandboxConfig,omitempty" xml:"sandboxConfig,omitempty"`
	// If omitted or set to null, the original list is retained. A non-empty array replaces the entire allowed set. An explicit empty array [] clears the inbound VPC restrictions for the new version. This field is independent of vpcConfig. An empty request body {}, an empty runtimeConfig, or a runtimeConfig that contains only vpcBindings: null returns a 400 error and does not clear the bindings. The actual JSON received by the server prevails. An empty array omitted by the SDK cannot clear the list.
	//
	// example:
	//
	// ["vpc-examplea","vpc-exampleb"]
	VpcBindings []*string `json:"vpcBindings" xml:"vpcBindings" type:"Repeated"`
	// The VPC network configuration.
	VpcConfig *PublicUpdateTemplateVPCConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s PublicUpdateTemplateRuntimeConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateRuntimeConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateRuntimeConfig) GetCpu() *float32 {
	return s.Cpu
}

func (s *PublicUpdateTemplateRuntimeConfig) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *PublicUpdateTemplateRuntimeConfig) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *PublicUpdateTemplateRuntimeConfig) GetLogConfig() *PublicUpdateTemplateLogConfig {
	return s.LogConfig
}

func (s *PublicUpdateTemplateRuntimeConfig) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *PublicUpdateTemplateRuntimeConfig) GetSandboxConfig() *PublicUpdateTemplateSandboxConfig {
	return s.SandboxConfig
}

func (s *PublicUpdateTemplateRuntimeConfig) GetVpcBindings() []*string {
	return s.VpcBindings
}

func (s *PublicUpdateTemplateRuntimeConfig) GetVpcConfig() *PublicUpdateTemplateVPCConfig {
	return s.VpcConfig
}

func (s *PublicUpdateTemplateRuntimeConfig) SetCpu(v float32) *PublicUpdateTemplateRuntimeConfig {
	s.Cpu = &v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetDiskSize(v int32) *PublicUpdateTemplateRuntimeConfig {
	s.DiskSize = &v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetInternetAccess(v bool) *PublicUpdateTemplateRuntimeConfig {
	s.InternetAccess = &v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetLogConfig(v *PublicUpdateTemplateLogConfig) *PublicUpdateTemplateRuntimeConfig {
	s.LogConfig = v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetMemorySize(v int32) *PublicUpdateTemplateRuntimeConfig {
	s.MemorySize = &v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetSandboxConfig(v *PublicUpdateTemplateSandboxConfig) *PublicUpdateTemplateRuntimeConfig {
	s.SandboxConfig = v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetVpcBindings(v []*string) *PublicUpdateTemplateRuntimeConfig {
	s.VpcBindings = v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) SetVpcConfig(v *PublicUpdateTemplateVPCConfig) *PublicUpdateTemplateRuntimeConfig {
	s.VpcConfig = v
	return s
}

func (s *PublicUpdateTemplateRuntimeConfig) Validate() error {
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

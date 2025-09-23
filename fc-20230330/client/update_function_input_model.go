// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionInput interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v *InputCodeLocation) *UpdateFunctionInput
	GetCode() *InputCodeLocation
	SetCpu(v float32) *UpdateFunctionInput
	GetCpu() *float32
	SetCustomContainerConfig(v *CustomContainerConfig) *UpdateFunctionInput
	GetCustomContainerConfig() *CustomContainerConfig
	SetCustomDNS(v *CustomDNS) *UpdateFunctionInput
	GetCustomDNS() *CustomDNS
	SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateFunctionInput
	GetCustomRuntimeConfig() *CustomRuntimeConfig
	SetDescription(v string) *UpdateFunctionInput
	GetDescription() *string
	SetDisableOndemand(v bool) *UpdateFunctionInput
	GetDisableOndemand() *bool
	SetDiskSize(v int32) *UpdateFunctionInput
	GetDiskSize() *int32
	SetEnableLongLiving(v bool) *UpdateFunctionInput
	GetEnableLongLiving() *bool
	SetEnvironmentVariables(v map[string]*string) *UpdateFunctionInput
	GetEnvironmentVariables() map[string]*string
	SetGpuConfig(v *GPUConfig) *UpdateFunctionInput
	GetGpuConfig() *GPUConfig
	SetHandler(v string) *UpdateFunctionInput
	GetHandler() *string
	SetIdleTimeout(v int32) *UpdateFunctionInput
	GetIdleTimeout() *int32
	SetInstanceConcurrency(v int32) *UpdateFunctionInput
	GetInstanceConcurrency() *int32
	SetInstanceIsolationMode(v string) *UpdateFunctionInput
	GetInstanceIsolationMode() *string
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateFunctionInput
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInternetAccess(v bool) *UpdateFunctionInput
	GetInternetAccess() *bool
	SetLayers(v []*string) *UpdateFunctionInput
	GetLayers() []*string
	SetLogConfig(v *LogConfig) *UpdateFunctionInput
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *UpdateFunctionInput
	GetMemorySize() *int32
	SetNasConfig(v *NASConfig) *UpdateFunctionInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *UpdateFunctionInput
	GetOssMountConfig() *OSSMountConfig
	SetRole(v string) *UpdateFunctionInput
	GetRole() *string
	SetRuntime(v string) *UpdateFunctionInput
	GetRuntime() *string
	SetSessionAffinity(v string) *UpdateFunctionInput
	GetSessionAffinity() *string
	SetSessionAffinityConfig(v string) *UpdateFunctionInput
	GetSessionAffinityConfig() *string
	SetTimeout(v int32) *UpdateFunctionInput
	GetTimeout() *int32
	SetTracingConfig(v *TracingConfig) *UpdateFunctionInput
	GetTracingConfig() *TracingConfig
	SetVpcConfig(v *VPCConfig) *UpdateFunctionInput
	GetVpcConfig() *VPCConfig
}

type UpdateFunctionInput struct {
	Code *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Cpu                   *float32               `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS             *CustomDNS             `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig   *CustomRuntimeConfig   `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// example:
	//
	// my function
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	DisableOndemand *bool   `json:"disableOndemand,omitempty" xml:"disableOndemand,omitempty"`
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Deprecated
	EnableLongLiving     *bool              `json:"enableLongLiving,omitempty" xml:"enableLongLiving,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	GpuConfig            *GPUConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// example:
	//
	// index.handler
	Handler     *string `json:"handler,omitempty" xml:"handler,omitempty"`
	IdleTimeout *int32  `json:"idleTimeout,omitempty" xml:"idleTimeout,omitempty"`
	// example:
	//
	// 1
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceIsolationMode   *string                  `json:"instanceIsolationMode,omitempty" xml:"instanceIsolationMode,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// example:
	//
	// true
	InternetAccess *bool      `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	Layers         []*string  `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig      *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// example:
	//
	// 512
	MemorySize     *int32          `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig      *NASConfig      `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// MCP_SSE
	SessionAffinity       *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	// example:
	//
	// 60
	Timeout       *int32         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s UpdateFunctionInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInput) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInput) GetCode() *InputCodeLocation {
	return s.Code
}

func (s *UpdateFunctionInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *UpdateFunctionInput) GetCustomContainerConfig() *CustomContainerConfig {
	return s.CustomContainerConfig
}

func (s *UpdateFunctionInput) GetCustomDNS() *CustomDNS {
	return s.CustomDNS
}

func (s *UpdateFunctionInput) GetCustomRuntimeConfig() *CustomRuntimeConfig {
	return s.CustomRuntimeConfig
}

func (s *UpdateFunctionInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateFunctionInput) GetDisableOndemand() *bool {
	return s.DisableOndemand
}

func (s *UpdateFunctionInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpdateFunctionInput) GetEnableLongLiving() *bool {
	return s.EnableLongLiving
}

func (s *UpdateFunctionInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateFunctionInput) GetGpuConfig() *GPUConfig {
	return s.GpuConfig
}

func (s *UpdateFunctionInput) GetHandler() *string {
	return s.Handler
}

func (s *UpdateFunctionInput) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *UpdateFunctionInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *UpdateFunctionInput) GetInstanceIsolationMode() *string {
	return s.InstanceIsolationMode
}

func (s *UpdateFunctionInput) GetInstanceLifecycleConfig() *InstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *UpdateFunctionInput) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *UpdateFunctionInput) GetLayers() []*string {
	return s.Layers
}

func (s *UpdateFunctionInput) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *UpdateFunctionInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *UpdateFunctionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateFunctionInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *UpdateFunctionInput) GetRole() *string {
	return s.Role
}

func (s *UpdateFunctionInput) GetRuntime() *string {
	return s.Runtime
}

func (s *UpdateFunctionInput) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *UpdateFunctionInput) GetSessionAffinityConfig() *string {
	return s.SessionAffinityConfig
}

func (s *UpdateFunctionInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateFunctionInput) GetTracingConfig() *TracingConfig {
	return s.TracingConfig
}

func (s *UpdateFunctionInput) GetVpcConfig() *VPCConfig {
	return s.VpcConfig
}

func (s *UpdateFunctionInput) SetCode(v *InputCodeLocation) *UpdateFunctionInput {
	s.Code = v
	return s
}

func (s *UpdateFunctionInput) SetCpu(v float32) *UpdateFunctionInput {
	s.Cpu = &v
	return s
}

func (s *UpdateFunctionInput) SetCustomContainerConfig(v *CustomContainerConfig) *UpdateFunctionInput {
	s.CustomContainerConfig = v
	return s
}

func (s *UpdateFunctionInput) SetCustomDNS(v *CustomDNS) *UpdateFunctionInput {
	s.CustomDNS = v
	return s
}

func (s *UpdateFunctionInput) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateFunctionInput {
	s.CustomRuntimeConfig = v
	return s
}

func (s *UpdateFunctionInput) SetDescription(v string) *UpdateFunctionInput {
	s.Description = &v
	return s
}

func (s *UpdateFunctionInput) SetDisableOndemand(v bool) *UpdateFunctionInput {
	s.DisableOndemand = &v
	return s
}

func (s *UpdateFunctionInput) SetDiskSize(v int32) *UpdateFunctionInput {
	s.DiskSize = &v
	return s
}

func (s *UpdateFunctionInput) SetEnableLongLiving(v bool) *UpdateFunctionInput {
	s.EnableLongLiving = &v
	return s
}

func (s *UpdateFunctionInput) SetEnvironmentVariables(v map[string]*string) *UpdateFunctionInput {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateFunctionInput) SetGpuConfig(v *GPUConfig) *UpdateFunctionInput {
	s.GpuConfig = v
	return s
}

func (s *UpdateFunctionInput) SetHandler(v string) *UpdateFunctionInput {
	s.Handler = &v
	return s
}

func (s *UpdateFunctionInput) SetIdleTimeout(v int32) *UpdateFunctionInput {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateFunctionInput) SetInstanceConcurrency(v int32) *UpdateFunctionInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionInput) SetInstanceIsolationMode(v string) *UpdateFunctionInput {
	s.InstanceIsolationMode = &v
	return s
}

func (s *UpdateFunctionInput) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateFunctionInput {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *UpdateFunctionInput) SetInternetAccess(v bool) *UpdateFunctionInput {
	s.InternetAccess = &v
	return s
}

func (s *UpdateFunctionInput) SetLayers(v []*string) *UpdateFunctionInput {
	s.Layers = v
	return s
}

func (s *UpdateFunctionInput) SetLogConfig(v *LogConfig) *UpdateFunctionInput {
	s.LogConfig = v
	return s
}

func (s *UpdateFunctionInput) SetMemorySize(v int32) *UpdateFunctionInput {
	s.MemorySize = &v
	return s
}

func (s *UpdateFunctionInput) SetNasConfig(v *NASConfig) *UpdateFunctionInput {
	s.NasConfig = v
	return s
}

func (s *UpdateFunctionInput) SetOssMountConfig(v *OSSMountConfig) *UpdateFunctionInput {
	s.OssMountConfig = v
	return s
}

func (s *UpdateFunctionInput) SetRole(v string) *UpdateFunctionInput {
	s.Role = &v
	return s
}

func (s *UpdateFunctionInput) SetRuntime(v string) *UpdateFunctionInput {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionInput) SetSessionAffinity(v string) *UpdateFunctionInput {
	s.SessionAffinity = &v
	return s
}

func (s *UpdateFunctionInput) SetSessionAffinityConfig(v string) *UpdateFunctionInput {
	s.SessionAffinityConfig = &v
	return s
}

func (s *UpdateFunctionInput) SetTimeout(v int32) *UpdateFunctionInput {
	s.Timeout = &v
	return s
}

func (s *UpdateFunctionInput) SetTracingConfig(v *TracingConfig) *UpdateFunctionInput {
	s.TracingConfig = v
	return s
}

func (s *UpdateFunctionInput) SetVpcConfig(v *VPCConfig) *UpdateFunctionInput {
	s.VpcConfig = v
	return s
}

func (s *UpdateFunctionInput) Validate() error {
	return dara.Validate(s)
}

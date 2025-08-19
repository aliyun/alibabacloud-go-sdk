// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionInput interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v *InputCodeLocation) *CreateFunctionInput
	GetCode() *InputCodeLocation
	SetCpu(v float32) *CreateFunctionInput
	GetCpu() *float32
	SetCustomContainerConfig(v *CustomContainerConfig) *CreateFunctionInput
	GetCustomContainerConfig() *CustomContainerConfig
	SetCustomDNS(v *CustomDNS) *CreateFunctionInput
	GetCustomDNS() *CustomDNS
	SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateFunctionInput
	GetCustomRuntimeConfig() *CustomRuntimeConfig
	SetDescription(v string) *CreateFunctionInput
	GetDescription() *string
	SetDisableOndemand(v bool) *CreateFunctionInput
	GetDisableOndemand() *bool
	SetDiskSize(v int32) *CreateFunctionInput
	GetDiskSize() *int32
	SetEnableLongLiving(v bool) *CreateFunctionInput
	GetEnableLongLiving() *bool
	SetEnvironmentVariables(v map[string]*string) *CreateFunctionInput
	GetEnvironmentVariables() map[string]*string
	SetFunctionName(v string) *CreateFunctionInput
	GetFunctionName() *string
	SetGpuConfig(v *GPUConfig) *CreateFunctionInput
	GetGpuConfig() *GPUConfig
	SetHandler(v string) *CreateFunctionInput
	GetHandler() *string
	SetInstanceConcurrency(v int32) *CreateFunctionInput
	GetInstanceConcurrency() *int32
	SetInstanceIsolationMode(v string) *CreateFunctionInput
	GetInstanceIsolationMode() *string
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateFunctionInput
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInternetAccess(v bool) *CreateFunctionInput
	GetInternetAccess() *bool
	SetLayers(v []*string) *CreateFunctionInput
	GetLayers() []*string
	SetLogConfig(v *LogConfig) *CreateFunctionInput
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *CreateFunctionInput
	GetMemorySize() *int32
	SetNasConfig(v *NASConfig) *CreateFunctionInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *CreateFunctionInput
	GetOssMountConfig() *OSSMountConfig
	SetResourceGroupId(v string) *CreateFunctionInput
	GetResourceGroupId() *string
	SetRole(v string) *CreateFunctionInput
	GetRole() *string
	SetRuntime(v string) *CreateFunctionInput
	GetRuntime() *string
	SetSessionAffinity(v string) *CreateFunctionInput
	GetSessionAffinity() *string
	SetSessionAffinityConfig(v string) *CreateFunctionInput
	GetSessionAffinityConfig() *string
	SetTags(v []*Tag) *CreateFunctionInput
	GetTags() []*Tag
	SetTimeout(v int32) *CreateFunctionInput
	GetTimeout() *int32
	SetTracingConfig(v *TracingConfig) *CreateFunctionInput
	GetTracingConfig() *TracingConfig
	SetVpcConfig(v *VPCConfig) *CreateFunctionInput
	GetVpcConfig() *VPCConfig
}

type CreateFunctionInput struct {
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
	// This parameter is required.
	//
	// example:
	//
	// my-function-1
	FunctionName *string    `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuConfig    *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
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
	MemorySize      *int32          `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig       *NASConfig      `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig  *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	ResourceGroupId *string         `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// python3.10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// MCP_SSE
	SessionAffinity       *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	Tags                  []*Tag  `json:"tags" xml:"tags" type:"Repeated"`
	// example:
	//
	// 60
	Timeout       *int32         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s CreateFunctionInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionInput) GoString() string {
	return s.String()
}

func (s *CreateFunctionInput) GetCode() *InputCodeLocation {
	return s.Code
}

func (s *CreateFunctionInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateFunctionInput) GetCustomContainerConfig() *CustomContainerConfig {
	return s.CustomContainerConfig
}

func (s *CreateFunctionInput) GetCustomDNS() *CustomDNS {
	return s.CustomDNS
}

func (s *CreateFunctionInput) GetCustomRuntimeConfig() *CustomRuntimeConfig {
	return s.CustomRuntimeConfig
}

func (s *CreateFunctionInput) GetDescription() *string {
	return s.Description
}

func (s *CreateFunctionInput) GetDisableOndemand() *bool {
	return s.DisableOndemand
}

func (s *CreateFunctionInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateFunctionInput) GetEnableLongLiving() *bool {
	return s.EnableLongLiving
}

func (s *CreateFunctionInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateFunctionInput) GetFunctionName() *string {
	return s.FunctionName
}

func (s *CreateFunctionInput) GetGpuConfig() *GPUConfig {
	return s.GpuConfig
}

func (s *CreateFunctionInput) GetHandler() *string {
	return s.Handler
}

func (s *CreateFunctionInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *CreateFunctionInput) GetInstanceIsolationMode() *string {
	return s.InstanceIsolationMode
}

func (s *CreateFunctionInput) GetInstanceLifecycleConfig() *InstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *CreateFunctionInput) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *CreateFunctionInput) GetLayers() []*string {
	return s.Layers
}

func (s *CreateFunctionInput) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *CreateFunctionInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *CreateFunctionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateFunctionInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *CreateFunctionInput) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFunctionInput) GetRole() *string {
	return s.Role
}

func (s *CreateFunctionInput) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateFunctionInput) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *CreateFunctionInput) GetSessionAffinityConfig() *string {
	return s.SessionAffinityConfig
}

func (s *CreateFunctionInput) GetTags() []*Tag {
	return s.Tags
}

func (s *CreateFunctionInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateFunctionInput) GetTracingConfig() *TracingConfig {
	return s.TracingConfig
}

func (s *CreateFunctionInput) GetVpcConfig() *VPCConfig {
	return s.VpcConfig
}

func (s *CreateFunctionInput) SetCode(v *InputCodeLocation) *CreateFunctionInput {
	s.Code = v
	return s
}

func (s *CreateFunctionInput) SetCpu(v float32) *CreateFunctionInput {
	s.Cpu = &v
	return s
}

func (s *CreateFunctionInput) SetCustomContainerConfig(v *CustomContainerConfig) *CreateFunctionInput {
	s.CustomContainerConfig = v
	return s
}

func (s *CreateFunctionInput) SetCustomDNS(v *CustomDNS) *CreateFunctionInput {
	s.CustomDNS = v
	return s
}

func (s *CreateFunctionInput) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateFunctionInput {
	s.CustomRuntimeConfig = v
	return s
}

func (s *CreateFunctionInput) SetDescription(v string) *CreateFunctionInput {
	s.Description = &v
	return s
}

func (s *CreateFunctionInput) SetDisableOndemand(v bool) *CreateFunctionInput {
	s.DisableOndemand = &v
	return s
}

func (s *CreateFunctionInput) SetDiskSize(v int32) *CreateFunctionInput {
	s.DiskSize = &v
	return s
}

func (s *CreateFunctionInput) SetEnableLongLiving(v bool) *CreateFunctionInput {
	s.EnableLongLiving = &v
	return s
}

func (s *CreateFunctionInput) SetEnvironmentVariables(v map[string]*string) *CreateFunctionInput {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateFunctionInput) SetFunctionName(v string) *CreateFunctionInput {
	s.FunctionName = &v
	return s
}

func (s *CreateFunctionInput) SetGpuConfig(v *GPUConfig) *CreateFunctionInput {
	s.GpuConfig = v
	return s
}

func (s *CreateFunctionInput) SetHandler(v string) *CreateFunctionInput {
	s.Handler = &v
	return s
}

func (s *CreateFunctionInput) SetInstanceConcurrency(v int32) *CreateFunctionInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateFunctionInput) SetInstanceIsolationMode(v string) *CreateFunctionInput {
	s.InstanceIsolationMode = &v
	return s
}

func (s *CreateFunctionInput) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateFunctionInput {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *CreateFunctionInput) SetInternetAccess(v bool) *CreateFunctionInput {
	s.InternetAccess = &v
	return s
}

func (s *CreateFunctionInput) SetLayers(v []*string) *CreateFunctionInput {
	s.Layers = v
	return s
}

func (s *CreateFunctionInput) SetLogConfig(v *LogConfig) *CreateFunctionInput {
	s.LogConfig = v
	return s
}

func (s *CreateFunctionInput) SetMemorySize(v int32) *CreateFunctionInput {
	s.MemorySize = &v
	return s
}

func (s *CreateFunctionInput) SetNasConfig(v *NASConfig) *CreateFunctionInput {
	s.NasConfig = v
	return s
}

func (s *CreateFunctionInput) SetOssMountConfig(v *OSSMountConfig) *CreateFunctionInput {
	s.OssMountConfig = v
	return s
}

func (s *CreateFunctionInput) SetResourceGroupId(v string) *CreateFunctionInput {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFunctionInput) SetRole(v string) *CreateFunctionInput {
	s.Role = &v
	return s
}

func (s *CreateFunctionInput) SetRuntime(v string) *CreateFunctionInput {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionInput) SetSessionAffinity(v string) *CreateFunctionInput {
	s.SessionAffinity = &v
	return s
}

func (s *CreateFunctionInput) SetSessionAffinityConfig(v string) *CreateFunctionInput {
	s.SessionAffinityConfig = &v
	return s
}

func (s *CreateFunctionInput) SetTags(v []*Tag) *CreateFunctionInput {
	s.Tags = v
	return s
}

func (s *CreateFunctionInput) SetTimeout(v int32) *CreateFunctionInput {
	s.Timeout = &v
	return s
}

func (s *CreateFunctionInput) SetTracingConfig(v *TracingConfig) *CreateFunctionInput {
	s.TracingConfig = v
	return s
}

func (s *CreateFunctionInput) SetVpcConfig(v *VPCConfig) *CreateFunctionInput {
	s.VpcConfig = v
	return s
}

func (s *CreateFunctionInput) Validate() error {
	return dara.Validate(s)
}

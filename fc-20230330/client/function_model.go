// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunction interface {
	dara.Model
	String() string
	GoString() string
	SetCodeChecksum(v string) *Function
	GetCodeChecksum() *string
	SetCodeSize(v int64) *Function
	GetCodeSize() *int64
	SetCpu(v float32) *Function
	GetCpu() *float32
	SetCreatedTime(v string) *Function
	GetCreatedTime() *string
	SetCustomContainerConfig(v *CustomContainerConfig) *Function
	GetCustomContainerConfig() *CustomContainerConfig
	SetCustomDNS(v *CustomDNS) *Function
	GetCustomDNS() *CustomDNS
	SetCustomRuntimeConfig(v *CustomRuntimeConfig) *Function
	GetCustomRuntimeConfig() *CustomRuntimeConfig
	SetDescription(v string) *Function
	GetDescription() *string
	SetDisableOndemand(v bool) *Function
	GetDisableOndemand() *bool
	SetDiskSize(v int32) *Function
	GetDiskSize() *int32
	SetEnableLongLiving(v bool) *Function
	GetEnableLongLiving() *bool
	SetEnvironmentVariables(v map[string]*string) *Function
	GetEnvironmentVariables() map[string]*string
	SetFunctionArn(v string) *Function
	GetFunctionArn() *string
	SetFunctionId(v string) *Function
	GetFunctionId() *string
	SetFunctionName(v string) *Function
	GetFunctionName() *string
	SetGpuConfig(v *GPUConfig) *Function
	GetGpuConfig() *GPUConfig
	SetHandler(v string) *Function
	GetHandler() *string
	SetIdleTimeout(v int32) *Function
	GetIdleTimeout() *int32
	SetInstanceConcurrency(v int32) *Function
	GetInstanceConcurrency() *int32
	SetInstanceIsolationMode(v string) *Function
	GetInstanceIsolationMode() *string
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *Function
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInternetAccess(v bool) *Function
	GetInternetAccess() *bool
	SetInvocationRestriction(v *FunctionRestriction) *Function
	GetInvocationRestriction() *FunctionRestriction
	SetLastModifiedTime(v string) *Function
	GetLastModifiedTime() *string
	SetLastUpdateStatus(v string) *Function
	GetLastUpdateStatus() *string
	SetLastUpdateStatusReason(v string) *Function
	GetLastUpdateStatusReason() *string
	SetLastUpdateStatusReasonCode(v string) *Function
	GetLastUpdateStatusReasonCode() *string
	SetLayers(v []*FunctionLayer) *Function
	GetLayers() []*FunctionLayer
	SetLogConfig(v *LogConfig) *Function
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *Function
	GetMemorySize() *int32
	SetNasConfig(v *NASConfig) *Function
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *Function
	GetOssMountConfig() *OSSMountConfig
	SetResourceGroupId(v string) *Function
	GetResourceGroupId() *string
	SetRole(v string) *Function
	GetRole() *string
	SetRuntime(v string) *Function
	GetRuntime() *string
	SetSessionAffinity(v string) *Function
	GetSessionAffinity() *string
	SetSessionAffinityConfig(v string) *Function
	GetSessionAffinityConfig() *string
	SetState(v string) *Function
	GetState() *string
	SetStateReason(v string) *Function
	GetStateReason() *string
	SetStateReasonCode(v string) *Function
	GetStateReasonCode() *string
	SetTags(v []*Tag) *Function
	GetTags() []*Tag
	SetTimeout(v int32) *Function
	GetTimeout() *int32
	SetTracingConfig(v *TracingConfig) *Function
	GetTracingConfig() *TracingConfig
	SetVpcConfig(v *VPCConfig) *Function
	GetVpcConfig() *VPCConfig
}

type Function struct {
	// example:
	//
	// 2825179536350****
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// example:
	//
	// 412
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// example:
	//
	// 1
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 2023-04-01T08:15:27Z
	CreatedTime           *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
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
	// example:
	//
	// acs:fc:cn-shanghai:123:functions/functionName
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// aa715851-1c20-4b89-a8fb-***
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// example:
	//
	// my-function-1
	FunctionName *string    `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuConfig    *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
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
	InternetAccess        *bool                `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	InvocationRestriction *FunctionRestriction `json:"invocationRestriction,omitempty" xml:"invocationRestriction,omitempty"`
	// example:
	//
	// 2023-05-01T08:15:27Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// InProgress
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty" xml:"lastUpdateStatus,omitempty"`
	// example:
	//
	// The system is currently processing the acceleration optimization for the image.
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty" xml:"lastUpdateStatusReason,omitempty"`
	// example:
	//
	// ImageOptimizing
	LastUpdateStatusReasonCode *string          `json:"lastUpdateStatusReasonCode,omitempty" xml:"lastUpdateStatusReasonCode,omitempty"`
	Layers                     []*FunctionLayer `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig                  *LogConfig       `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
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
	// example:
	//
	// python3.10
	Runtime         *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	// example:
	//
	// {\"sseEndpointPath\":\"/sse\", \"sessionConcurrencyPerInstance\":20}
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	// example:
	//
	// Pending
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// Function creating
	StateReason *string `json:"stateReason,omitempty" xml:"stateReason,omitempty"`
	// example:
	//
	// Creating
	StateReasonCode *string `json:"stateReasonCode,omitempty" xml:"stateReasonCode,omitempty"`
	Tags            []*Tag  `json:"tags" xml:"tags" type:"Repeated"`
	// example:
	//
	// 60
	Timeout       *int32         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s Function) String() string {
	return dara.Prettify(s)
}

func (s Function) GoString() string {
	return s.String()
}

func (s *Function) GetCodeChecksum() *string {
	return s.CodeChecksum
}

func (s *Function) GetCodeSize() *int64 {
	return s.CodeSize
}

func (s *Function) GetCpu() *float32 {
	return s.Cpu
}

func (s *Function) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Function) GetCustomContainerConfig() *CustomContainerConfig {
	return s.CustomContainerConfig
}

func (s *Function) GetCustomDNS() *CustomDNS {
	return s.CustomDNS
}

func (s *Function) GetCustomRuntimeConfig() *CustomRuntimeConfig {
	return s.CustomRuntimeConfig
}

func (s *Function) GetDescription() *string {
	return s.Description
}

func (s *Function) GetDisableOndemand() *bool {
	return s.DisableOndemand
}

func (s *Function) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *Function) GetEnableLongLiving() *bool {
	return s.EnableLongLiving
}

func (s *Function) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *Function) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *Function) GetFunctionId() *string {
	return s.FunctionId
}

func (s *Function) GetFunctionName() *string {
	return s.FunctionName
}

func (s *Function) GetGpuConfig() *GPUConfig {
	return s.GpuConfig
}

func (s *Function) GetHandler() *string {
	return s.Handler
}

func (s *Function) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *Function) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *Function) GetInstanceIsolationMode() *string {
	return s.InstanceIsolationMode
}

func (s *Function) GetInstanceLifecycleConfig() *InstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *Function) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *Function) GetInvocationRestriction() *FunctionRestriction {
	return s.InvocationRestriction
}

func (s *Function) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Function) GetLastUpdateStatus() *string {
	return s.LastUpdateStatus
}

func (s *Function) GetLastUpdateStatusReason() *string {
	return s.LastUpdateStatusReason
}

func (s *Function) GetLastUpdateStatusReasonCode() *string {
	return s.LastUpdateStatusReasonCode
}

func (s *Function) GetLayers() []*FunctionLayer {
	return s.Layers
}

func (s *Function) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *Function) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *Function) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *Function) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *Function) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Function) GetRole() *string {
	return s.Role
}

func (s *Function) GetRuntime() *string {
	return s.Runtime
}

func (s *Function) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *Function) GetSessionAffinityConfig() *string {
	return s.SessionAffinityConfig
}

func (s *Function) GetState() *string {
	return s.State
}

func (s *Function) GetStateReason() *string {
	return s.StateReason
}

func (s *Function) GetStateReasonCode() *string {
	return s.StateReasonCode
}

func (s *Function) GetTags() []*Tag {
	return s.Tags
}

func (s *Function) GetTimeout() *int32 {
	return s.Timeout
}

func (s *Function) GetTracingConfig() *TracingConfig {
	return s.TracingConfig
}

func (s *Function) GetVpcConfig() *VPCConfig {
	return s.VpcConfig
}

func (s *Function) SetCodeChecksum(v string) *Function {
	s.CodeChecksum = &v
	return s
}

func (s *Function) SetCodeSize(v int64) *Function {
	s.CodeSize = &v
	return s
}

func (s *Function) SetCpu(v float32) *Function {
	s.Cpu = &v
	return s
}

func (s *Function) SetCreatedTime(v string) *Function {
	s.CreatedTime = &v
	return s
}

func (s *Function) SetCustomContainerConfig(v *CustomContainerConfig) *Function {
	s.CustomContainerConfig = v
	return s
}

func (s *Function) SetCustomDNS(v *CustomDNS) *Function {
	s.CustomDNS = v
	return s
}

func (s *Function) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *Function {
	s.CustomRuntimeConfig = v
	return s
}

func (s *Function) SetDescription(v string) *Function {
	s.Description = &v
	return s
}

func (s *Function) SetDisableOndemand(v bool) *Function {
	s.DisableOndemand = &v
	return s
}

func (s *Function) SetDiskSize(v int32) *Function {
	s.DiskSize = &v
	return s
}

func (s *Function) SetEnableLongLiving(v bool) *Function {
	s.EnableLongLiving = &v
	return s
}

func (s *Function) SetEnvironmentVariables(v map[string]*string) *Function {
	s.EnvironmentVariables = v
	return s
}

func (s *Function) SetFunctionArn(v string) *Function {
	s.FunctionArn = &v
	return s
}

func (s *Function) SetFunctionId(v string) *Function {
	s.FunctionId = &v
	return s
}

func (s *Function) SetFunctionName(v string) *Function {
	s.FunctionName = &v
	return s
}

func (s *Function) SetGpuConfig(v *GPUConfig) *Function {
	s.GpuConfig = v
	return s
}

func (s *Function) SetHandler(v string) *Function {
	s.Handler = &v
	return s
}

func (s *Function) SetIdleTimeout(v int32) *Function {
	s.IdleTimeout = &v
	return s
}

func (s *Function) SetInstanceConcurrency(v int32) *Function {
	s.InstanceConcurrency = &v
	return s
}

func (s *Function) SetInstanceIsolationMode(v string) *Function {
	s.InstanceIsolationMode = &v
	return s
}

func (s *Function) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *Function {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *Function) SetInternetAccess(v bool) *Function {
	s.InternetAccess = &v
	return s
}

func (s *Function) SetInvocationRestriction(v *FunctionRestriction) *Function {
	s.InvocationRestriction = v
	return s
}

func (s *Function) SetLastModifiedTime(v string) *Function {
	s.LastModifiedTime = &v
	return s
}

func (s *Function) SetLastUpdateStatus(v string) *Function {
	s.LastUpdateStatus = &v
	return s
}

func (s *Function) SetLastUpdateStatusReason(v string) *Function {
	s.LastUpdateStatusReason = &v
	return s
}

func (s *Function) SetLastUpdateStatusReasonCode(v string) *Function {
	s.LastUpdateStatusReasonCode = &v
	return s
}

func (s *Function) SetLayers(v []*FunctionLayer) *Function {
	s.Layers = v
	return s
}

func (s *Function) SetLogConfig(v *LogConfig) *Function {
	s.LogConfig = v
	return s
}

func (s *Function) SetMemorySize(v int32) *Function {
	s.MemorySize = &v
	return s
}

func (s *Function) SetNasConfig(v *NASConfig) *Function {
	s.NasConfig = v
	return s
}

func (s *Function) SetOssMountConfig(v *OSSMountConfig) *Function {
	s.OssMountConfig = v
	return s
}

func (s *Function) SetResourceGroupId(v string) *Function {
	s.ResourceGroupId = &v
	return s
}

func (s *Function) SetRole(v string) *Function {
	s.Role = &v
	return s
}

func (s *Function) SetRuntime(v string) *Function {
	s.Runtime = &v
	return s
}

func (s *Function) SetSessionAffinity(v string) *Function {
	s.SessionAffinity = &v
	return s
}

func (s *Function) SetSessionAffinityConfig(v string) *Function {
	s.SessionAffinityConfig = &v
	return s
}

func (s *Function) SetState(v string) *Function {
	s.State = &v
	return s
}

func (s *Function) SetStateReason(v string) *Function {
	s.StateReason = &v
	return s
}

func (s *Function) SetStateReasonCode(v string) *Function {
	s.StateReasonCode = &v
	return s
}

func (s *Function) SetTags(v []*Tag) *Function {
	s.Tags = v
	return s
}

func (s *Function) SetTimeout(v int32) *Function {
	s.Timeout = &v
	return s
}

func (s *Function) SetTracingConfig(v *TracingConfig) *Function {
	s.TracingConfig = v
	return s
}

func (s *Function) SetVpcConfig(v *VPCConfig) *Function {
	s.VpcConfig = v
	return s
}

func (s *Function) Validate() error {
	return dara.Validate(s)
}

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
	SetDisableInjectCredentials(v string) *Function
	GetDisableInjectCredentials() *string
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
	SetJuiceFsConfig(v *JuiceFsConfig) *Function
	GetJuiceFsConfig() *JuiceFsConfig
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
	SetLockInfo(v *FunctionLockInfo) *Function
	GetLockInfo() *FunctionLockInfo
	SetLogConfig(v *LogConfig) *Function
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *Function
	GetMemorySize() *int32
	SetMicroSandboxConfig(v *MicroSandboxConfig) *Function
	GetMicroSandboxConfig() *MicroSandboxConfig
	SetNasConfig(v *NASConfig) *Function
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *Function
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *Function
	GetPolarFsConfig() *PolarFsConfig
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
	// The CRC-64 value of the function code package.
	//
	// example:
	//
	// 2825179536350****
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// The size of the function code package returned by the system. Unit: bytes.
	//
	// example:
	//
	// 412
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// The CPU specification of the function. Unit: vCPU. The value must be a multiple of 0.05 vCPU. Minimum value: 0.05. Maximum value: 16. The ratio of cpu to memorySize (in GB) must be between 1:1 and 1:4.
	//
	// example:
	//
	// 1
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The time when the function was created.
	//
	// example:
	//
	// 2023-04-01T08:15:27Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The custom container runtime configuration. After this parameter is configured, the function can use a custom container image to execute the function. Specify either code or customContainerConfig.
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// The custom DNS configuration.
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// The custom runtime configuration.
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// The description of the function.
	//
	// example:
	//
	// my function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to disable STS token injection. Valid values:
	//
	// - None: injects STS tokens in all methods.
	//
	// - Env: does not inject STS tokens through environment variables.
	//
	// - Request: does not inject STS tokens through requests, including context and headers.
	//
	// - All: does not inject STS tokens in any method.
	//
	// example:
	//
	// Env
	DisableInjectCredentials *string `json:"disableInjectCredentials,omitempty" xml:"disableInjectCredentials,omitempty"`
	// Deprecated
	//
	// Specifies whether to disable the creation of on-demand instances. If this feature is enabled, on-demand instances are not created, and only provisioned instances can be used.
	DisableOndemand *bool `json:"disableOndemand,omitempty" xml:"disableOndemand,omitempty"`
	// The disk specification of the function. Unit: MB. Valid values: 512 and 10240.
	//
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Deprecated
	//
	// When a sessionAffinity type is set, configure the corresponding affinity settings. For MCP_SSE affinity, populate the MCPSSESessionAffinityConfig configuration. For cookie-based affinity, populate the CookieSessionAffinityConfig configuration. For header field affinity, populate the HeaderFieldSessionAffinityConfig configuration.
	EnableLongLiving *bool `json:"enableLongLiving,omitempty" xml:"enableLongLiving,omitempty"`
	// The environment variables of the function. You can access the configured environment variables in the runtime environment.
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// The Alibaba Cloud Resource Name (ARN) of the function.
	//
	// example:
	//
	// acs:fc:cn-shanghai:123:functions/functionName
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// The globally unique ID generated by the system for the function.
	//
	// example:
	//
	// aa715851-1c20-4b89-a8fb-***
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// my-function-1
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The GPU configuration of the function.
	GpuConfig *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// The function entry point. The specific format depends on the runtime.
	//
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The deferred instance release time.
	//
	// example:
	//
	// 100
	IdleTimeout *int32 `json:"idleTimeout,omitempty" xml:"idleTimeout,omitempty"`
	// The maximum concurrency per instance.
	//
	// example:
	//
	// 1
	InstanceConcurrency *int32 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	// The instance isolation mode.
	InstanceIsolationMode *string `json:"instanceIsolationMode,omitempty" xml:"instanceIsolationMode,omitempty"`
	// The instance lifecycle hook method configuration.
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// Specifies whether the function can access the Internet. Default value: true.
	//
	// example:
	//
	// true
	InternetAccess        *bool                `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	InvocationRestriction *FunctionRestriction `json:"invocationRestriction,omitempty" xml:"invocationRestriction,omitempty"`
	JuiceFsConfig         *JuiceFsConfig       `json:"juiceFsConfig,omitempty" xml:"juiceFsConfig,omitempty"`
	// The time when the function was last updated.
	//
	// example:
	//
	// 2023-05-01T08:15:27Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The status of the most recent function update operation. When a function is created, this value is Successful. Valid values:
	//
	// - Successful
	//
	// - Failed
	//
	// - InProgress.
	//
	// example:
	//
	// InProgress
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty" xml:"lastUpdateStatus,omitempty"`
	// The reason that caused the most recent function update operation to have the current status.
	//
	// example:
	//
	// The system is currently processing the acceleration optimization for the image.
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty" xml:"lastUpdateStatusReason,omitempty"`
	// The status code of the reason that caused the most recent function update operation to have the current status.
	//
	// example:
	//
	// ImageOptimizing
	LastUpdateStatusReasonCode *string `json:"lastUpdateStatusReasonCode,omitempty" xml:"lastUpdateStatusReasonCode,omitempty"`
	// The list of layers.
	Layers []*FunctionLayer `json:"layers" xml:"layers" type:"Repeated"`
	// example:
	//
	// {"lockedBy":"AgentRun","lockedAt":"2025-04-05T10:00:00Z","lockedResources":["function","trigger","version","alias"]}
	LockInfo *FunctionLockInfo `json:"lockInfo,omitempty" xml:"lockInfo,omitempty"`
	// The log configuration. Logs generated by the function are written to the configured Logstore.
	LogConfig *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory specification of the function. Unit: MB. The value must be a multiple of 64 MB. Minimum value: 128. Maximum value: 32768 (32 GB). The ratio of cpu to memorySize (in GB) must be between 1:1 and 1:4.
	//
	// example:
	//
	// 512
	MemorySize         *int32              `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	MicroSandboxConfig *MicroSandboxConfig `json:"microSandboxConfig,omitempty" xml:"microSandboxConfig,omitempty"`
	// The NAS configuration. After this parameter is configured, the function can access the specified NAS resources.
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The OSS mount configuration.
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// The PolarFs configuration. After this parameter is configured, the function can access the specified PolarFs resources.
	PolarFsConfig *PolarFsConfig `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The RAM role that the user grants to Function Compute. After this parameter is configured, Function Compute assumes this role to generate temporary access credentials. You can use the temporary access credentials of this role in the function to access specified Alibaba Cloud services such as OSS and OTS.
	//
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The runtime environment of the function. Currently supported runtime environments include: nodejs12, nodejs14, nodejs16, nodejs18, nodejs20, go1, python3, python3.9, python3.10, python3.12, java8, java11, php7.2, dotnetcore3.1, custom, custom.debian10, custom.debian11, custom.debian12, and custom-container.
	//
	// example:
	//
	// python3.10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The affinity policy for Function Compute invocation requests. To implement request affinity for the MCP SSE protocol, set this parameter to MCP_SSE. To use cookie-based affinity, set this parameter to GENERATED_COOKIE. To use header-based affinity, set this parameter to HEADER_FIELD. If this parameter is not set or is set to NONE, no affinity is applied, and requests are routed based on the default scheduling policy of Function Compute.
	//
	// example:
	//
	// MCP_SSE
	SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	// When a sessionAffinity type is set, configure the corresponding affinity settings. For MCP_SSE affinity, populate the MCPSSESessionAffinityConfig configuration. For cookie-based affinity, populate the CookieSessionAffinityConfig configuration. For header field affinity, populate the HeaderFieldSessionAffinityConfig configuration.
	//
	// example:
	//
	// {\\"sseEndpointPath\\":\\"/sse\\", \\"sessionConcurrencyPerInstance\\":20}
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	// The current state of the function.
	//
	// example:
	//
	// Pending
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason why the function is in the current state.
	//
	// example:
	//
	// Function creating
	StateReason *string `json:"stateReason,omitempty" xml:"stateReason,omitempty"`
	// The status code of the reason why the function is in the current state.
	//
	// example:
	//
	// Creating
	StateReasonCode *string `json:"stateReasonCode,omitempty" xml:"stateReasonCode,omitempty"`
	// The list of tags.
	Tags []*Tag `json:"tags" xml:"tags" type:"Repeated"`
	// The timeout period for the function execution. Unit: seconds. Minimum value: 1. Maximum value: 86400. Default value: 3. The function is terminated if it exceeds this time limit.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The Tracing Analysis configuration. After Function Compute is integrated with Tracing Analysis, you can record the time consumed by requests in Function Compute, view the cold start time of functions, and record the time consumed by internal operations of functions.
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	// The VPC configuration. After this parameter is configured, the function can access the specified VPC resources.
	VpcConfig *VPCConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
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

func (s *Function) GetDisableInjectCredentials() *string {
	return s.DisableInjectCredentials
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

func (s *Function) GetJuiceFsConfig() *JuiceFsConfig {
	return s.JuiceFsConfig
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

func (s *Function) GetLockInfo() *FunctionLockInfo {
	return s.LockInfo
}

func (s *Function) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *Function) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *Function) GetMicroSandboxConfig() *MicroSandboxConfig {
	return s.MicroSandboxConfig
}

func (s *Function) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *Function) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *Function) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
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

func (s *Function) SetDisableInjectCredentials(v string) *Function {
	s.DisableInjectCredentials = &v
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

func (s *Function) SetJuiceFsConfig(v *JuiceFsConfig) *Function {
	s.JuiceFsConfig = v
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

func (s *Function) SetLockInfo(v *FunctionLockInfo) *Function {
	s.LockInfo = v
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

func (s *Function) SetMicroSandboxConfig(v *MicroSandboxConfig) *Function {
	s.MicroSandboxConfig = v
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

func (s *Function) SetPolarFsConfig(v *PolarFsConfig) *Function {
	s.PolarFsConfig = v
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
	if s.CustomContainerConfig != nil {
		if err := s.CustomContainerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CustomDNS != nil {
		if err := s.CustomDNS.Validate(); err != nil {
			return err
		}
	}
	if s.CustomRuntimeConfig != nil {
		if err := s.CustomRuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.GpuConfig != nil {
		if err := s.GpuConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceLifecycleConfig != nil {
		if err := s.InstanceLifecycleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InvocationRestriction != nil {
		if err := s.InvocationRestriction.Validate(); err != nil {
			return err
		}
	}
	if s.JuiceFsConfig != nil {
		if err := s.JuiceFsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Layers != nil {
		for _, item := range s.Layers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LockInfo != nil {
		if err := s.LockInfo.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfig != nil {
		if err := s.LogConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MicroSandboxConfig != nil {
		if err := s.MicroSandboxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NasConfig != nil {
		if err := s.NasConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OssMountConfig != nil {
		if err := s.OssMountConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PolarFsConfig != nil {
		if err := s.PolarFsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TracingConfig != nil {
		if err := s.TracingConfig.Validate(); err != nil {
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

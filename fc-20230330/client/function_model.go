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
	// The CPU power allocated to the function. Unit: vCPUs. The value must be a multiple of 0.05. The minimum value is 0.05 and the maximum value is 16. The ratio of cpu to memorySize (in GB) must be from 1:1 to 1:4.
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
	// The configurations of the Custom Container runtime. After you configure a Custom Container runtime for your function, Function Compute can execute the function in a custom container image. Configure either code or customContainerConfig.
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// The custom DNS settings of the function.
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// The configurations of the custom runtime.
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// The description of the function.
	//
	// example:
	//
	// my function
	Description              *string `json:"description,omitempty" xml:"description,omitempty"`
	DisableInjectCredentials *string `json:"disableInjectCredentials,omitempty" xml:"disableInjectCredentials,omitempty"`
	// Deprecated
	DisableOndemand *bool `json:"disableOndemand,omitempty" xml:"disableOndemand,omitempty"`
	// The disk size of the function. Unit: MB. Valid values: 512 and 10240.
	//
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Deprecated
	EnableLongLiving *bool `json:"enableLongLiving,omitempty" xml:"enableLongLiving,omitempty"`
	// The environment variables of the function. You can access the specified environment variables in the runtime.
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// The identifier of the function resource.
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
	// The GPU configurations of the function.
	GpuConfig *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// The handler of the function. The format of the handler is related to the runtime you use.
	//
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// example:
	//
	// 100
	IdleTimeout *int32 `json:"idleTimeout,omitempty" xml:"idleTimeout,omitempty"`
	// The maximum number of requests that a function instance can process at a time.
	//
	// example:
	//
	// 1
	InstanceConcurrency   *int32  `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceIsolationMode *string `json:"instanceIsolationMode,omitempty" xml:"instanceIsolationMode,omitempty"`
	// The configurations of instance lifecycle hooks.
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// Specifies whether to allow the function to access the Internet. Default value: true.
	//
	// example:
	//
	// true
	InternetAccess        *bool                `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	InvocationRestriction *FunctionRestriction `json:"invocationRestriction,omitempty" xml:"invocationRestriction,omitempty"`
	// The last time the function was updated.
	//
	// example:
	//
	// 2023-05-01T08:15:27Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The status of the most recent update that was executed on the function. The initial value for this parameter is Successful once the function has been successfully created. Valid values: Successful, Failed, and InProgress.
	//
	// example:
	//
	// InProgress
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty" xml:"lastUpdateStatus,omitempty"`
	// The reason for the most recent update that was executed on the function.
	//
	// example:
	//
	// The system is currently processing the acceleration optimization for the image.
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty" xml:"lastUpdateStatusReason,omitempty"`
	// The reason code for the most recent update that was executed on the function.
	//
	// example:
	//
	// ImageOptimizing
	LastUpdateStatusReasonCode *string `json:"lastUpdateStatusReasonCode,omitempty" xml:"lastUpdateStatusReasonCode,omitempty"`
	// The layers.
	Layers []*FunctionLayer `json:"layers" xml:"layers" type:"Repeated"`
	// The logging configurations. Logs generated by the function are written to the specified Logstore.
	LogConfig *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory capacity for the function. Unit: MB. The value must be a multiple of 64. The minimum capacity is 128 MB and the maximum capacity is 32 GB. The ratio of cpu to memorySize (in GB) must be from 1:1 to 1:4.
	//
	// example:
	//
	// 512
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The File Storage NAS (NAS) configurations. The configurations allow the function to access the specified NAS file system.
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The OSS mounting configurations.
	OssMountConfig  *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	PolarFsConfig   *PolarFsConfig  `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	ResourceGroupId *string         `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Resource Access Management (RAM) role that is assigned to the function. Function Compute assumes the role to obtain a Security Token Service (STS) token, which serves as a temporary key for your function to access other Alibaba Cloud services, such as Object Storage Service (OSS) and Tablestore.
	//
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The runtime of the function. Valid values: nodejs8, nodejs10, nodejs12, nodejs14, nodejs16, nodejs18, nodejs20, go1, python3, python3.9, python3.10, java8, java11, php7.2, dotnetcore3.1, custom, custom.debian10, and custom-container.
	//
	// example:
	//
	// python3.10
	Runtime         *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	// example:
	//
	// {\"sseEndpointPath\":\"/sse\", \"sessionConcurrencyPerInstance\":20}
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	// The current state of the function.
	//
	// example:
	//
	// Pending
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason for the current state of the function.
	//
	// example:
	//
	// Function creating
	StateReason *string `json:"stateReason,omitempty" xml:"stateReason,omitempty"`
	// The reason code for the current state of the function.
	//
	// example:
	//
	// Creating
	StateReasonCode *string `json:"stateReasonCode,omitempty" xml:"stateReasonCode,omitempty"`
	// The tags.
	Tags []*Tag `json:"tags" xml:"tags" type:"Repeated"`
	// The timeout period for function execution. Unit: seconds. Default value: 3. Valid values: 1 to 86400. The execution of the function is terminated when the timeout period expires.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The configurations of Managed Service for OpenTelemetry. After Function Compute is integrated with Managed Service for OpenTelemetry, you can record the invocation duration of a request, view the cold start duration of a function, and track the execution duration of the function.
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	// The Virtual Private Cloud (VPC) configurations. The configurations allow the function to access the specified VPC resources.
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
	if s.Layers != nil {
		for _, item := range s.Layers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogConfig != nil {
		if err := s.LogConfig.Validate(); err != nil {
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

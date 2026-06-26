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
	SetDisableInjectCredentials(v string) *CreateFunctionInput
	GetDisableInjectCredentials() *string
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
	SetIdleTimeout(v int32) *CreateFunctionInput
	GetIdleTimeout() *int32
	SetInstanceConcurrency(v int32) *CreateFunctionInput
	GetInstanceConcurrency() *int32
	SetInstanceIsolationMode(v string) *CreateFunctionInput
	GetInstanceIsolationMode() *string
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateFunctionInput
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInternetAccess(v bool) *CreateFunctionInput
	GetInternetAccess() *bool
	SetJuiceFsConfig(v *JuiceFsConfig) *CreateFunctionInput
	GetJuiceFsConfig() *JuiceFsConfig
	SetLayers(v []*string) *CreateFunctionInput
	GetLayers() []*string
	SetLogConfig(v *LogConfig) *CreateFunctionInput
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *CreateFunctionInput
	GetMemorySize() *int32
	SetMicroSandboxConfig(v *MicroSandboxConfig) *CreateFunctionInput
	GetMicroSandboxConfig() *MicroSandboxConfig
	SetNasConfig(v *NASConfig) *CreateFunctionInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *CreateFunctionInput
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *CreateFunctionInput
	GetPolarFsConfig() *PolarFsConfig
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
	// The ZIP package of the function code. Specify either code or customContainerConfig.
	Code *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// The CPU specification of the function, in vCPUs. The value must be a multiple of 0.05 vCPU. Minimum value: 0.05. Maximum value: 16. The ratio of cpu to memorySize (in GB) must be between 1:1 and 1:4.
	//
	// example:
	//
	// 1
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The configuration for the custom container runtime. After this parameter is configured, the function can use a custom container image for execution. Specify either code or customContainerConfig.
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
	// - None: STS tokens are injected in all methods.
	//
	// - Env: STS tokens are not injected through environment variables.
	//
	// - Request: STS tokens are not injected in requests, including context and headers.
	//
	// - All: STS tokens are not injected in any method.
	//
	// example:
	//
	// Env
	DisableInjectCredentials *string `json:"disableInjectCredentials,omitempty" xml:"disableInjectCredentials,omitempty"`
	// Deprecated
	//
	// Specifies whether to disable the creation of on-demand instances. If this feature is enabled, on-demand instances are not created and only provisioned instances can be used.
	DisableOndemand *bool `json:"disableOndemand,omitempty" xml:"disableOndemand,omitempty"`
	// The disk specification of the function, in MB. Valid values: 512 and 10240.
	//
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// Deprecated
	//
	// Specifies whether to allow provisioned instances of GPU functions to be long-running. When this feature is enabled, function instances are not injected with STS tokens.
	EnableLongLiving *bool `json:"enableLongLiving,omitempty" xml:"enableLongLiving,omitempty"`
	// The environment variables of the function. You can access the configured environment variables in the runtime environment.
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// The name of the function. The name can contain only letters, digits, underscores (_), and hyphens (-). The name cannot start with a digit or hyphen (-). The name must be 1 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-function-1
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The GPU configuration of the function.
	GpuConfig *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// The function entry point. The specific format depends on the runtime.
	//
	// This parameter is required.
	//
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The deferred release time of the instance.
	//
	// example:
	//
	// 100
	IdleTimeout *int32 `json:"idleTimeout,omitempty" xml:"idleTimeout,omitempty"`
	// The maximum concurrency of an instance.
	//
	// example:
	//
	// 1
	InstanceConcurrency *int32 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	// The instance isolation mode.
	InstanceIsolationMode *string `json:"instanceIsolationMode,omitempty" xml:"instanceIsolationMode,omitempty"`
	// The instance lifecycle hook configuration.
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// Specifies whether the function can access the Internet. Default value: true.
	//
	// example:
	//
	// true
	InternetAccess *bool          `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	JuiceFsConfig  *JuiceFsConfig `json:"juiceFsConfig,omitempty" xml:"juiceFsConfig,omitempty"`
	// The list of layers. Multiple layers are merged in descending order of array index. Files in a layer with a smaller index overwrite files with the same name in a layer with a larger index.
	Layers []*string `json:"layers" xml:"layers" type:"Repeated"`
	// The log configuration. Logs generated by the function are written to the configured Logstore.
	LogConfig *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// The memory specification of the function, in MB. The value must be a multiple of 64 MB. Minimum value: 128. Maximum value: 32768 (32 GB). The ratio of cpu to memorySize (in GB) must be between 1:1 and 1:4.
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
	PolarFsConfig   *PolarFsConfig `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	ResourceGroupId *string        `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The RAM role that the user grants to Function Compute. After this parameter is set, Function Compute assumes this role to generate temporary access credentials. You can use the temporary access credentials of this role in the function to access specified Alibaba Cloud services, such as OSS and OTS.
	//
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The runtime environment of the function. Supported runtimes: nodejs12, nodejs14, nodejs16, nodejs18, nodejs20, go1, python3, python3.9, python3.10, python3.12, java8, java11, php7.2, dotnetcore3.1, custom, custom.debian10, custom.debian11, custom.debian12, and custom-container.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3.10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The affinity policy for Function Compute invocation requests. To implement request affinity for the MCP SSE protocol, set this parameter to MCP_SSE. To use cookie-based affinity, set this parameter to GENERATED_COOKIE. To use header-based affinity, set this parameter to HEADER_FIELD. If this parameter is not set or is set to NONE, no affinity is applied and requests are routed based on the default scheduling policy of Function Compute.
	//
	// example:
	//
	// MCP_SSE
	SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	// The affinity configuration that corresponds to the sessionAffinity type. For MCP_SSE affinity, specify MCPSSESessionAffinityConfig. For cookie-based affinity, specify CookieSessionAffinityConfig. For header field affinity, specify HeaderFieldSessionAffinityConfig.
	//
	// example:
	//
	// {\\"sseEndpointPath\\":\\"/sse\\", \\"sessionConcurrencyPerInstance\\":20}
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	// The list of tags.
	Tags []*Tag `json:"tags" xml:"tags" type:"Repeated"`
	// The timeout period for function execution, in seconds. Minimum value: 1. Maximum value: 86400. Default value: 3. The function is terminated if it exceeds this time limit.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The Tracing Analysis configuration. After Function Compute is integrated with Tracing Analysis, you can record the time consumed by requests in Function Compute, view the cold start time of functions, and record the time consumed within functions.
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	// The VPC configuration. After this parameter is configured, the function can access the specified VPC resources.
	VpcConfig *VPCConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
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

func (s *CreateFunctionInput) GetDisableInjectCredentials() *string {
	return s.DisableInjectCredentials
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

func (s *CreateFunctionInput) GetIdleTimeout() *int32 {
	return s.IdleTimeout
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

func (s *CreateFunctionInput) GetJuiceFsConfig() *JuiceFsConfig {
	return s.JuiceFsConfig
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

func (s *CreateFunctionInput) GetMicroSandboxConfig() *MicroSandboxConfig {
	return s.MicroSandboxConfig
}

func (s *CreateFunctionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateFunctionInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *CreateFunctionInput) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
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

func (s *CreateFunctionInput) SetDisableInjectCredentials(v string) *CreateFunctionInput {
	s.DisableInjectCredentials = &v
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

func (s *CreateFunctionInput) SetIdleTimeout(v int32) *CreateFunctionInput {
	s.IdleTimeout = &v
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

func (s *CreateFunctionInput) SetJuiceFsConfig(v *JuiceFsConfig) *CreateFunctionInput {
	s.JuiceFsConfig = v
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

func (s *CreateFunctionInput) SetMicroSandboxConfig(v *MicroSandboxConfig) *CreateFunctionInput {
	s.MicroSandboxConfig = v
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

func (s *CreateFunctionInput) SetPolarFsConfig(v *PolarFsConfig) *CreateFunctionInput {
	s.PolarFsConfig = v
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
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			return err
		}
	}
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
	if s.JuiceFsConfig != nil {
		if err := s.JuiceFsConfig.Validate(); err != nil {
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

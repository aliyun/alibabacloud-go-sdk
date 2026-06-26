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
	SetDisableInjectCredentials(v string) *UpdateFunctionInput
	GetDisableInjectCredentials() *string
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
	SetJuiceFsConfig(v *JuiceFsConfig) *UpdateFunctionInput
	GetJuiceFsConfig() *JuiceFsConfig
	SetLayers(v []*string) *UpdateFunctionInput
	GetLayers() []*string
	SetLogConfig(v *LogConfig) *UpdateFunctionInput
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *UpdateFunctionInput
	GetMemorySize() *int32
	SetMicroSandboxConfig(v *MicroSandboxConfig) *UpdateFunctionInput
	GetMicroSandboxConfig() *MicroSandboxConfig
	SetNasConfig(v *NASConfig) *UpdateFunctionInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *UpdateFunctionInput
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *UpdateFunctionInput
	GetPolarFsConfig() *PolarFsConfig
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
	// The ZIP package of the function code. Specify either code or customContainerConfig.
	Code *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// The CPU specification of the function. Unit: vCPU. The value must be a multiple of 0.05 vCPU.
	//
	// example:
	//
	// 1
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The configuration of the custom container runtime. After this parameter is configured, the function can use a custom container image for execution. Specify either code or customContainerConfig.
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
	// - Request: STS tokens are not injected through requests, including context and headers.
	//
	// - All: STS tokens are not injected in any method.
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
	// Specifies whether to allow provisioned instances of GPU functions to be long-running. When this feature is enabled, function instances that are created are not injected with STS tokens.
	EnableLongLiving *bool `json:"enableLongLiving,omitempty" xml:"enableLongLiving,omitempty"`
	// The environment variables of the function. You can access the configured environment variables in the runtime environment.
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// The GPU configuration of the function.
	GpuConfig *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// The function entry point. The specific format depends on the runtime.
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
	// Specifies whether to allow access to the Internet.
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
	// The memory specification of the function. Unit: MB. The value must be a multiple of 64 MB. The memory specification varies based on the function instance type.
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
	// The Alibaba Cloud Resource Access Management (RAM) role that grants Function Compute the required permissions. Scenarios include: 1. Sending logs generated by the function to your Logstore. 2. Generating temporary access tokens for the function to access other cloud resources during the execute procedure.
	//
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The runtime environment of the function.
	//
	// example:
	//
	// nodejs14
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The affinity policy for Function Compute invocation requests. To implement request affinity for the MCP SSE protocol, set this parameter to MCP_SSE. To use cookie-based affinity, set this parameter to GENERATED_COOKIE. To use header-based affinity, set this parameter to HEADER_FIELD. If this parameter is not set or is set to NONE, no affinity is applied, and requests are routed based on the default scheduling policy of Function Compute.
	//
	// example:
	//
	// MCP_SSE
	SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	// The affinity configuration that corresponds to the sessionAffinity type. For MCP_SSE affinity, configure MCPSSESessionAffinityConfig. For cookie-based affinity, configure CookieSessionAffinityConfig. For header field affinity, configure HeaderFieldSessionAffinityConfig.
	//
	// example:
	//
	// {\\"sseEndpointPath\\":\\"/sse\\", \\"sessionConcurrencyPerInstance\\":20}
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
	// The timeout period for function execution. Unit: seconds. Minimum value: 1. Default value: 3. The function is terminated if it exceeds this time limit.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The Tracing Analysis configuration. After Function Compute is integrated with Tracing Analysis, you can record the time consumed by requests in Function Compute, view the cold start time of functions, and record the time consumed by internal function operations.
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	// The VPC configuration. After this parameter is configured, the function can access the specified VPC resources.
	VpcConfig *VPCConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
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

func (s *UpdateFunctionInput) GetDisableInjectCredentials() *string {
	return s.DisableInjectCredentials
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

func (s *UpdateFunctionInput) GetJuiceFsConfig() *JuiceFsConfig {
	return s.JuiceFsConfig
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

func (s *UpdateFunctionInput) GetMicroSandboxConfig() *MicroSandboxConfig {
	return s.MicroSandboxConfig
}

func (s *UpdateFunctionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateFunctionInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *UpdateFunctionInput) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
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

func (s *UpdateFunctionInput) SetDisableInjectCredentials(v string) *UpdateFunctionInput {
	s.DisableInjectCredentials = &v
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

func (s *UpdateFunctionInput) SetJuiceFsConfig(v *JuiceFsConfig) *UpdateFunctionInput {
	s.JuiceFsConfig = v
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

func (s *UpdateFunctionInput) SetMicroSandboxConfig(v *MicroSandboxConfig) *UpdateFunctionInput {
	s.MicroSandboxConfig = v
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

func (s *UpdateFunctionInput) SetPolarFsConfig(v *PolarFsConfig) *UpdateFunctionInput {
	s.PolarFsConfig = v
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

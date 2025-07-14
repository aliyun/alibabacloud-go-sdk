// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationInput interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v string) *UpdateApplicationInput
	GetArgs() *string
	SetCaPort(v int32) *UpdateApplicationInput
	GetCaPort() *int32
	SetCode(v *InputCodeLocation) *UpdateApplicationInput
	GetCode() *InputCodeLocation
	SetCommand(v string) *UpdateApplicationInput
	GetCommand() *string
	SetCpu(v float32) *UpdateApplicationInput
	GetCpu() *float32
	SetCustomDNS(v *CustomDNS) *UpdateApplicationInput
	GetCustomDNS() *CustomDNS
	SetCustomHealthCheckConfig(v *CustomHealthCheckConfig) *UpdateApplicationInput
	GetCustomHealthCheckConfig() *CustomHealthCheckConfig
	SetCustomHostAlias(v *CustomHostAlias) *UpdateApplicationInput
	GetCustomHostAlias() *CustomHostAlias
	SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateApplicationInput
	GetCustomRuntimeConfig() *CustomRuntimeConfig
	SetDescription(v string) *UpdateApplicationInput
	GetDescription() *string
	SetDiskSize(v int32) *UpdateApplicationInput
	GetDiskSize() *int32
	SetEffectiveImmediately(v bool) *UpdateApplicationInput
	GetEffectiveImmediately() *bool
	SetEnableAppMetric(v bool) *UpdateApplicationInput
	GetEnableAppMetric() *bool
	SetEnvironmentVariables(v map[string]*string) *UpdateApplicationInput
	GetEnvironmentVariables() map[string]*string
	SetGpuMemorySize(v int32) *UpdateApplicationInput
	GetGpuMemorySize() *int32
	SetHandler(v string) *UpdateApplicationInput
	GetHandler() *string
	SetHttpTriggerConfig(v *HTTPTriggerConfig) *UpdateApplicationInput
	GetHttpTriggerConfig() *HTTPTriggerConfig
	SetImageConfig(v *ImageConfig) *UpdateApplicationInput
	GetImageConfig() *ImageConfig
	SetInitializationTimeout(v int32) *UpdateApplicationInput
	GetInitializationTimeout() *int32
	SetInitializer(v string) *UpdateApplicationInput
	GetInitializer() *string
	SetInstanceConcurrency(v int32) *UpdateApplicationInput
	GetInstanceConcurrency() *int32
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateApplicationInput
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInstanceSoftConcurrency(v int32) *UpdateApplicationInput
	GetInstanceSoftConcurrency() *int32
	SetInstanceType(v string) *UpdateApplicationInput
	GetInstanceType() *string
	SetInternetAccess(v bool) *UpdateApplicationInput
	GetInternetAccess() *bool
	SetLayers(v []*string) *UpdateApplicationInput
	GetLayers() []*string
	SetLivenessProbe(v *Probe) *UpdateApplicationInput
	GetLivenessProbe() *Probe
	SetLogConfig(v *LogConfig) *UpdateApplicationInput
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *UpdateApplicationInput
	GetMemorySize() *int32
	SetNamespaceID(v string) *UpdateApplicationInput
	GetNamespaceID() *string
	SetNasConfig(v *NASConfig) *UpdateApplicationInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *UpdateApplicationInput
	GetOssMountConfig() *OSSMountConfig
	SetProgrammingLanguage(v string) *UpdateApplicationInput
	GetProgrammingLanguage() *string
	SetRuntime(v string) *UpdateApplicationInput
	GetRuntime() *string
	SetScaleConfig(v *ScaleConfig) *UpdateApplicationInput
	GetScaleConfig() *ScaleConfig
	SetSlsConfig(v *SLSConfig) *UpdateApplicationInput
	GetSlsConfig() *SLSConfig
	SetStartupProbe(v *Probe) *UpdateApplicationInput
	GetStartupProbe() *Probe
	SetTimeout(v int32) *UpdateApplicationInput
	GetTimeout() *int32
	SetTracingConfig(v *TracingConfig) *UpdateApplicationInput
	GetTracingConfig() *TracingConfig
	SetVpcConfig(v *VPCConfig) *UpdateApplicationInput
	GetVpcConfig() *VPCConfig
}

type UpdateApplicationInput struct {
	Args                    *string                  `json:"args,omitempty" xml:"args,omitempty"`
	CaPort                  *int32                   `json:"caPort,omitempty" xml:"caPort,omitempty"`
	Code                    *InputCodeLocation       `json:"code,omitempty" xml:"code,omitempty"`
	Command                 *string                  `json:"command,omitempty" xml:"command,omitempty"`
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomDNS               *CustomDNS               `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomHealthCheckConfig *CustomHealthCheckConfig `json:"customHealthCheckConfig,omitempty" xml:"customHealthCheckConfig,omitempty"`
	CustomHostAlias         *CustomHostAlias         `json:"customHostAlias,omitempty" xml:"customHostAlias,omitempty"`
	CustomRuntimeConfig     *CustomRuntimeConfig     `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EffectiveImmediately    *bool                    `json:"effectiveImmediately,omitempty" xml:"effectiveImmediately,omitempty"`
	EnableAppMetric         *bool                    `json:"enableAppMetric,omitempty" xml:"enableAppMetric,omitempty"`
	EnvironmentVariables    map[string]*string       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuMemorySize           *int32                   `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	Handler                 *string                  `json:"handler,omitempty" xml:"handler,omitempty"`
	HttpTriggerConfig       *HTTPTriggerConfig       `json:"httpTriggerConfig,omitempty" xml:"httpTriggerConfig,omitempty"`
	ImageConfig             *ImageConfig             `json:"imageConfig,omitempty" xml:"imageConfig,omitempty"`
	InitializationTimeout   *int32                   `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceSoftConcurrency *int32                   `json:"instanceSoftConcurrency,omitempty" xml:"instanceSoftConcurrency,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	InternetAccess          *bool                    `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	Layers                  []*string                `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	LivenessProbe           *Probe                   `json:"livenessProbe,omitempty" xml:"livenessProbe,omitempty"`
	LogConfig               *LogConfig               `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	MemorySize              *int32                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NamespaceID             *string                  `json:"namespaceID,omitempty" xml:"namespaceID,omitempty"`
	NasConfig               *NASConfig               `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig          *OSSMountConfig          `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	ProgrammingLanguage     *string                  `json:"programmingLanguage,omitempty" xml:"programmingLanguage,omitempty"`
	Runtime                 *string                  `json:"runtime,omitempty" xml:"runtime,omitempty"`
	ScaleConfig             *ScaleConfig             `json:"scaleConfig,omitempty" xml:"scaleConfig,omitempty"`
	SlsConfig               *SLSConfig               `json:"slsConfig,omitempty" xml:"slsConfig,omitempty"`
	StartupProbe            *Probe                   `json:"startupProbe,omitempty" xml:"startupProbe,omitempty"`
	Timeout                 *int32                   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig           *TracingConfig           `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig               *VPCConfig               `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s UpdateApplicationInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationInput) GoString() string {
	return s.String()
}

func (s *UpdateApplicationInput) GetArgs() *string {
	return s.Args
}

func (s *UpdateApplicationInput) GetCaPort() *int32 {
	return s.CaPort
}

func (s *UpdateApplicationInput) GetCode() *InputCodeLocation {
	return s.Code
}

func (s *UpdateApplicationInput) GetCommand() *string {
	return s.Command
}

func (s *UpdateApplicationInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *UpdateApplicationInput) GetCustomDNS() *CustomDNS {
	return s.CustomDNS
}

func (s *UpdateApplicationInput) GetCustomHealthCheckConfig() *CustomHealthCheckConfig {
	return s.CustomHealthCheckConfig
}

func (s *UpdateApplicationInput) GetCustomHostAlias() *CustomHostAlias {
	return s.CustomHostAlias
}

func (s *UpdateApplicationInput) GetCustomRuntimeConfig() *CustomRuntimeConfig {
	return s.CustomRuntimeConfig
}

func (s *UpdateApplicationInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpdateApplicationInput) GetEffectiveImmediately() *bool {
	return s.EffectiveImmediately
}

func (s *UpdateApplicationInput) GetEnableAppMetric() *bool {
	return s.EnableAppMetric
}

func (s *UpdateApplicationInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateApplicationInput) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *UpdateApplicationInput) GetHandler() *string {
	return s.Handler
}

func (s *UpdateApplicationInput) GetHttpTriggerConfig() *HTTPTriggerConfig {
	return s.HttpTriggerConfig
}

func (s *UpdateApplicationInput) GetImageConfig() *ImageConfig {
	return s.ImageConfig
}

func (s *UpdateApplicationInput) GetInitializationTimeout() *int32 {
	return s.InitializationTimeout
}

func (s *UpdateApplicationInput) GetInitializer() *string {
	return s.Initializer
}

func (s *UpdateApplicationInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *UpdateApplicationInput) GetInstanceLifecycleConfig() *InstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *UpdateApplicationInput) GetInstanceSoftConcurrency() *int32 {
	return s.InstanceSoftConcurrency
}

func (s *UpdateApplicationInput) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateApplicationInput) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *UpdateApplicationInput) GetLayers() []*string {
	return s.Layers
}

func (s *UpdateApplicationInput) GetLivenessProbe() *Probe {
	return s.LivenessProbe
}

func (s *UpdateApplicationInput) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *UpdateApplicationInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *UpdateApplicationInput) GetNamespaceID() *string {
	return s.NamespaceID
}

func (s *UpdateApplicationInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateApplicationInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *UpdateApplicationInput) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *UpdateApplicationInput) GetRuntime() *string {
	return s.Runtime
}

func (s *UpdateApplicationInput) GetScaleConfig() *ScaleConfig {
	return s.ScaleConfig
}

func (s *UpdateApplicationInput) GetSlsConfig() *SLSConfig {
	return s.SlsConfig
}

func (s *UpdateApplicationInput) GetStartupProbe() *Probe {
	return s.StartupProbe
}

func (s *UpdateApplicationInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateApplicationInput) GetTracingConfig() *TracingConfig {
	return s.TracingConfig
}

func (s *UpdateApplicationInput) GetVpcConfig() *VPCConfig {
	return s.VpcConfig
}

func (s *UpdateApplicationInput) SetArgs(v string) *UpdateApplicationInput {
	s.Args = &v
	return s
}

func (s *UpdateApplicationInput) SetCaPort(v int32) *UpdateApplicationInput {
	s.CaPort = &v
	return s
}

func (s *UpdateApplicationInput) SetCode(v *InputCodeLocation) *UpdateApplicationInput {
	s.Code = v
	return s
}

func (s *UpdateApplicationInput) SetCommand(v string) *UpdateApplicationInput {
	s.Command = &v
	return s
}

func (s *UpdateApplicationInput) SetCpu(v float32) *UpdateApplicationInput {
	s.Cpu = &v
	return s
}

func (s *UpdateApplicationInput) SetCustomDNS(v *CustomDNS) *UpdateApplicationInput {
	s.CustomDNS = v
	return s
}

func (s *UpdateApplicationInput) SetCustomHealthCheckConfig(v *CustomHealthCheckConfig) *UpdateApplicationInput {
	s.CustomHealthCheckConfig = v
	return s
}

func (s *UpdateApplicationInput) SetCustomHostAlias(v *CustomHostAlias) *UpdateApplicationInput {
	s.CustomHostAlias = v
	return s
}

func (s *UpdateApplicationInput) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateApplicationInput {
	s.CustomRuntimeConfig = v
	return s
}

func (s *UpdateApplicationInput) SetDescription(v string) *UpdateApplicationInput {
	s.Description = &v
	return s
}

func (s *UpdateApplicationInput) SetDiskSize(v int32) *UpdateApplicationInput {
	s.DiskSize = &v
	return s
}

func (s *UpdateApplicationInput) SetEffectiveImmediately(v bool) *UpdateApplicationInput {
	s.EffectiveImmediately = &v
	return s
}

func (s *UpdateApplicationInput) SetEnableAppMetric(v bool) *UpdateApplicationInput {
	s.EnableAppMetric = &v
	return s
}

func (s *UpdateApplicationInput) SetEnvironmentVariables(v map[string]*string) *UpdateApplicationInput {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateApplicationInput) SetGpuMemorySize(v int32) *UpdateApplicationInput {
	s.GpuMemorySize = &v
	return s
}

func (s *UpdateApplicationInput) SetHandler(v string) *UpdateApplicationInput {
	s.Handler = &v
	return s
}

func (s *UpdateApplicationInput) SetHttpTriggerConfig(v *HTTPTriggerConfig) *UpdateApplicationInput {
	s.HttpTriggerConfig = v
	return s
}

func (s *UpdateApplicationInput) SetImageConfig(v *ImageConfig) *UpdateApplicationInput {
	s.ImageConfig = v
	return s
}

func (s *UpdateApplicationInput) SetInitializationTimeout(v int32) *UpdateApplicationInput {
	s.InitializationTimeout = &v
	return s
}

func (s *UpdateApplicationInput) SetInitializer(v string) *UpdateApplicationInput {
	s.Initializer = &v
	return s
}

func (s *UpdateApplicationInput) SetInstanceConcurrency(v int32) *UpdateApplicationInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateApplicationInput) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateApplicationInput {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *UpdateApplicationInput) SetInstanceSoftConcurrency(v int32) *UpdateApplicationInput {
	s.InstanceSoftConcurrency = &v
	return s
}

func (s *UpdateApplicationInput) SetInstanceType(v string) *UpdateApplicationInput {
	s.InstanceType = &v
	return s
}

func (s *UpdateApplicationInput) SetInternetAccess(v bool) *UpdateApplicationInput {
	s.InternetAccess = &v
	return s
}

func (s *UpdateApplicationInput) SetLayers(v []*string) *UpdateApplicationInput {
	s.Layers = v
	return s
}

func (s *UpdateApplicationInput) SetLivenessProbe(v *Probe) *UpdateApplicationInput {
	s.LivenessProbe = v
	return s
}

func (s *UpdateApplicationInput) SetLogConfig(v *LogConfig) *UpdateApplicationInput {
	s.LogConfig = v
	return s
}

func (s *UpdateApplicationInput) SetMemorySize(v int32) *UpdateApplicationInput {
	s.MemorySize = &v
	return s
}

func (s *UpdateApplicationInput) SetNamespaceID(v string) *UpdateApplicationInput {
	s.NamespaceID = &v
	return s
}

func (s *UpdateApplicationInput) SetNasConfig(v *NASConfig) *UpdateApplicationInput {
	s.NasConfig = v
	return s
}

func (s *UpdateApplicationInput) SetOssMountConfig(v *OSSMountConfig) *UpdateApplicationInput {
	s.OssMountConfig = v
	return s
}

func (s *UpdateApplicationInput) SetProgrammingLanguage(v string) *UpdateApplicationInput {
	s.ProgrammingLanguage = &v
	return s
}

func (s *UpdateApplicationInput) SetRuntime(v string) *UpdateApplicationInput {
	s.Runtime = &v
	return s
}

func (s *UpdateApplicationInput) SetScaleConfig(v *ScaleConfig) *UpdateApplicationInput {
	s.ScaleConfig = v
	return s
}

func (s *UpdateApplicationInput) SetSlsConfig(v *SLSConfig) *UpdateApplicationInput {
	s.SlsConfig = v
	return s
}

func (s *UpdateApplicationInput) SetStartupProbe(v *Probe) *UpdateApplicationInput {
	s.StartupProbe = v
	return s
}

func (s *UpdateApplicationInput) SetTimeout(v int32) *UpdateApplicationInput {
	s.Timeout = &v
	return s
}

func (s *UpdateApplicationInput) SetTracingConfig(v *TracingConfig) *UpdateApplicationInput {
	s.TracingConfig = v
	return s
}

func (s *UpdateApplicationInput) SetVpcConfig(v *VPCConfig) *UpdateApplicationInput {
	s.VpcConfig = v
	return s
}

func (s *UpdateApplicationInput) Validate() error {
	return dara.Validate(s)
}

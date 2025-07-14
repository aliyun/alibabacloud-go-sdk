// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationInput interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *CreateApplicationInput
	GetApplicationName() *string
	SetArgs(v string) *CreateApplicationInput
	GetArgs() *string
	SetCaPort(v int32) *CreateApplicationInput
	GetCaPort() *int32
	SetCode(v *InputCodeLocation) *CreateApplicationInput
	GetCode() *InputCodeLocation
	SetCommand(v string) *CreateApplicationInput
	GetCommand() *string
	SetCpu(v float32) *CreateApplicationInput
	GetCpu() *float32
	SetCustomDNS(v *CustomDNS) *CreateApplicationInput
	GetCustomDNS() *CustomDNS
	SetCustomHealthCheckConfig(v *CustomHealthCheckConfig) *CreateApplicationInput
	GetCustomHealthCheckConfig() *CustomHealthCheckConfig
	SetCustomHostAlias(v *CustomHostAlias) *CreateApplicationInput
	GetCustomHostAlias() *CustomHostAlias
	SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateApplicationInput
	GetCustomRuntimeConfig() *CustomRuntimeConfig
	SetDescription(v string) *CreateApplicationInput
	GetDescription() *string
	SetDiskSize(v int32) *CreateApplicationInput
	GetDiskSize() *int32
	SetEnableAppMetric(v bool) *CreateApplicationInput
	GetEnableAppMetric() *bool
	SetEnvironmentVariables(v map[string]*string) *CreateApplicationInput
	GetEnvironmentVariables() map[string]*string
	SetGpuMemorySize(v int32) *CreateApplicationInput
	GetGpuMemorySize() *int32
	SetHandler(v string) *CreateApplicationInput
	GetHandler() *string
	SetHttpTriggerConfig(v *HTTPTriggerConfig) *CreateApplicationInput
	GetHttpTriggerConfig() *HTTPTriggerConfig
	SetImageConfig(v *ImageConfig) *CreateApplicationInput
	GetImageConfig() *ImageConfig
	SetInitializationTimeout(v int32) *CreateApplicationInput
	GetInitializationTimeout() *int32
	SetInitializer(v string) *CreateApplicationInput
	GetInitializer() *string
	SetInstanceConcurrency(v int32) *CreateApplicationInput
	GetInstanceConcurrency() *int32
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateApplicationInput
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInstanceSoftConcurrency(v int32) *CreateApplicationInput
	GetInstanceSoftConcurrency() *int32
	SetInstanceType(v string) *CreateApplicationInput
	GetInstanceType() *string
	SetInternetAccess(v bool) *CreateApplicationInput
	GetInternetAccess() *bool
	SetLayers(v []*string) *CreateApplicationInput
	GetLayers() []*string
	SetLivenessProbe(v *Probe) *CreateApplicationInput
	GetLivenessProbe() *Probe
	SetLogConfig(v *LogConfig) *CreateApplicationInput
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *CreateApplicationInput
	GetMemorySize() *int32
	SetNamespaceID(v string) *CreateApplicationInput
	GetNamespaceID() *string
	SetNasConfig(v *NASConfig) *CreateApplicationInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *CreateApplicationInput
	GetOssMountConfig() *OSSMountConfig
	SetProgrammingLanguage(v string) *CreateApplicationInput
	GetProgrammingLanguage() *string
	SetRuntime(v string) *CreateApplicationInput
	GetRuntime() *string
	SetScaleConfig(v *ScaleConfig) *CreateApplicationInput
	GetScaleConfig() *ScaleConfig
	SetSlsConfig(v *SLSConfig) *CreateApplicationInput
	GetSlsConfig() *SLSConfig
	SetStartupProbe(v *Probe) *CreateApplicationInput
	GetStartupProbe() *Probe
	SetTimeout(v int32) *CreateApplicationInput
	GetTimeout() *int32
	SetTracingConfig(v *TracingConfig) *CreateApplicationInput
	GetTracingConfig() *TracingConfig
	SetVpcConfig(v *VPCConfig) *CreateApplicationInput
	GetVpcConfig() *VPCConfig
}

type CreateApplicationInput struct {
	ApplicationName         *string                  `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
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

func (s CreateApplicationInput) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationInput) GoString() string {
	return s.String()
}

func (s *CreateApplicationInput) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateApplicationInput) GetArgs() *string {
	return s.Args
}

func (s *CreateApplicationInput) GetCaPort() *int32 {
	return s.CaPort
}

func (s *CreateApplicationInput) GetCode() *InputCodeLocation {
	return s.Code
}

func (s *CreateApplicationInput) GetCommand() *string {
	return s.Command
}

func (s *CreateApplicationInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateApplicationInput) GetCustomDNS() *CustomDNS {
	return s.CustomDNS
}

func (s *CreateApplicationInput) GetCustomHealthCheckConfig() *CustomHealthCheckConfig {
	return s.CustomHealthCheckConfig
}

func (s *CreateApplicationInput) GetCustomHostAlias() *CustomHostAlias {
	return s.CustomHostAlias
}

func (s *CreateApplicationInput) GetCustomRuntimeConfig() *CustomRuntimeConfig {
	return s.CustomRuntimeConfig
}

func (s *CreateApplicationInput) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateApplicationInput) GetEnableAppMetric() *bool {
	return s.EnableAppMetric
}

func (s *CreateApplicationInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateApplicationInput) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *CreateApplicationInput) GetHandler() *string {
	return s.Handler
}

func (s *CreateApplicationInput) GetHttpTriggerConfig() *HTTPTriggerConfig {
	return s.HttpTriggerConfig
}

func (s *CreateApplicationInput) GetImageConfig() *ImageConfig {
	return s.ImageConfig
}

func (s *CreateApplicationInput) GetInitializationTimeout() *int32 {
	return s.InitializationTimeout
}

func (s *CreateApplicationInput) GetInitializer() *string {
	return s.Initializer
}

func (s *CreateApplicationInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *CreateApplicationInput) GetInstanceLifecycleConfig() *InstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *CreateApplicationInput) GetInstanceSoftConcurrency() *int32 {
	return s.InstanceSoftConcurrency
}

func (s *CreateApplicationInput) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateApplicationInput) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *CreateApplicationInput) GetLayers() []*string {
	return s.Layers
}

func (s *CreateApplicationInput) GetLivenessProbe() *Probe {
	return s.LivenessProbe
}

func (s *CreateApplicationInput) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *CreateApplicationInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *CreateApplicationInput) GetNamespaceID() *string {
	return s.NamespaceID
}

func (s *CreateApplicationInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateApplicationInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *CreateApplicationInput) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *CreateApplicationInput) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateApplicationInput) GetScaleConfig() *ScaleConfig {
	return s.ScaleConfig
}

func (s *CreateApplicationInput) GetSlsConfig() *SLSConfig {
	return s.SlsConfig
}

func (s *CreateApplicationInput) GetStartupProbe() *Probe {
	return s.StartupProbe
}

func (s *CreateApplicationInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateApplicationInput) GetTracingConfig() *TracingConfig {
	return s.TracingConfig
}

func (s *CreateApplicationInput) GetVpcConfig() *VPCConfig {
	return s.VpcConfig
}

func (s *CreateApplicationInput) SetApplicationName(v string) *CreateApplicationInput {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationInput) SetArgs(v string) *CreateApplicationInput {
	s.Args = &v
	return s
}

func (s *CreateApplicationInput) SetCaPort(v int32) *CreateApplicationInput {
	s.CaPort = &v
	return s
}

func (s *CreateApplicationInput) SetCode(v *InputCodeLocation) *CreateApplicationInput {
	s.Code = v
	return s
}

func (s *CreateApplicationInput) SetCommand(v string) *CreateApplicationInput {
	s.Command = &v
	return s
}

func (s *CreateApplicationInput) SetCpu(v float32) *CreateApplicationInput {
	s.Cpu = &v
	return s
}

func (s *CreateApplicationInput) SetCustomDNS(v *CustomDNS) *CreateApplicationInput {
	s.CustomDNS = v
	return s
}

func (s *CreateApplicationInput) SetCustomHealthCheckConfig(v *CustomHealthCheckConfig) *CreateApplicationInput {
	s.CustomHealthCheckConfig = v
	return s
}

func (s *CreateApplicationInput) SetCustomHostAlias(v *CustomHostAlias) *CreateApplicationInput {
	s.CustomHostAlias = v
	return s
}

func (s *CreateApplicationInput) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateApplicationInput {
	s.CustomRuntimeConfig = v
	return s
}

func (s *CreateApplicationInput) SetDescription(v string) *CreateApplicationInput {
	s.Description = &v
	return s
}

func (s *CreateApplicationInput) SetDiskSize(v int32) *CreateApplicationInput {
	s.DiskSize = &v
	return s
}

func (s *CreateApplicationInput) SetEnableAppMetric(v bool) *CreateApplicationInput {
	s.EnableAppMetric = &v
	return s
}

func (s *CreateApplicationInput) SetEnvironmentVariables(v map[string]*string) *CreateApplicationInput {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateApplicationInput) SetGpuMemorySize(v int32) *CreateApplicationInput {
	s.GpuMemorySize = &v
	return s
}

func (s *CreateApplicationInput) SetHandler(v string) *CreateApplicationInput {
	s.Handler = &v
	return s
}

func (s *CreateApplicationInput) SetHttpTriggerConfig(v *HTTPTriggerConfig) *CreateApplicationInput {
	s.HttpTriggerConfig = v
	return s
}

func (s *CreateApplicationInput) SetImageConfig(v *ImageConfig) *CreateApplicationInput {
	s.ImageConfig = v
	return s
}

func (s *CreateApplicationInput) SetInitializationTimeout(v int32) *CreateApplicationInput {
	s.InitializationTimeout = &v
	return s
}

func (s *CreateApplicationInput) SetInitializer(v string) *CreateApplicationInput {
	s.Initializer = &v
	return s
}

func (s *CreateApplicationInput) SetInstanceConcurrency(v int32) *CreateApplicationInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateApplicationInput) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateApplicationInput {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *CreateApplicationInput) SetInstanceSoftConcurrency(v int32) *CreateApplicationInput {
	s.InstanceSoftConcurrency = &v
	return s
}

func (s *CreateApplicationInput) SetInstanceType(v string) *CreateApplicationInput {
	s.InstanceType = &v
	return s
}

func (s *CreateApplicationInput) SetInternetAccess(v bool) *CreateApplicationInput {
	s.InternetAccess = &v
	return s
}

func (s *CreateApplicationInput) SetLayers(v []*string) *CreateApplicationInput {
	s.Layers = v
	return s
}

func (s *CreateApplicationInput) SetLivenessProbe(v *Probe) *CreateApplicationInput {
	s.LivenessProbe = v
	return s
}

func (s *CreateApplicationInput) SetLogConfig(v *LogConfig) *CreateApplicationInput {
	s.LogConfig = v
	return s
}

func (s *CreateApplicationInput) SetMemorySize(v int32) *CreateApplicationInput {
	s.MemorySize = &v
	return s
}

func (s *CreateApplicationInput) SetNamespaceID(v string) *CreateApplicationInput {
	s.NamespaceID = &v
	return s
}

func (s *CreateApplicationInput) SetNasConfig(v *NASConfig) *CreateApplicationInput {
	s.NasConfig = v
	return s
}

func (s *CreateApplicationInput) SetOssMountConfig(v *OSSMountConfig) *CreateApplicationInput {
	s.OssMountConfig = v
	return s
}

func (s *CreateApplicationInput) SetProgrammingLanguage(v string) *CreateApplicationInput {
	s.ProgrammingLanguage = &v
	return s
}

func (s *CreateApplicationInput) SetRuntime(v string) *CreateApplicationInput {
	s.Runtime = &v
	return s
}

func (s *CreateApplicationInput) SetScaleConfig(v *ScaleConfig) *CreateApplicationInput {
	s.ScaleConfig = v
	return s
}

func (s *CreateApplicationInput) SetSlsConfig(v *SLSConfig) *CreateApplicationInput {
	s.SlsConfig = v
	return s
}

func (s *CreateApplicationInput) SetStartupProbe(v *Probe) *CreateApplicationInput {
	s.StartupProbe = v
	return s
}

func (s *CreateApplicationInput) SetTimeout(v int32) *CreateApplicationInput {
	s.Timeout = &v
	return s
}

func (s *CreateApplicationInput) SetTracingConfig(v *TracingConfig) *CreateApplicationInput {
	s.TracingConfig = v
	return s
}

func (s *CreateApplicationInput) SetVpcConfig(v *VPCConfig) *CreateApplicationInput {
	s.VpcConfig = v
	return s
}

func (s *CreateApplicationInput) Validate() error {
	return dara.Validate(s)
}

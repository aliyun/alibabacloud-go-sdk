// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplication interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *Application
	GetRequestId() *string
	SetApplicationId(v string) *Application
	GetApplicationId() *string
	SetApplicationName(v string) *Application
	GetApplicationName() *string
	SetArgs(v string) *Application
	GetArgs() *string
	SetCaPort(v int32) *Application
	GetCaPort() *int32
	SetCodeChecksum(v string) *Application
	GetCodeChecksum() *string
	SetCodeSize(v int64) *Application
	GetCodeSize() *int64
	SetCommand(v string) *Application
	GetCommand() *string
	SetCpu(v float32) *Application
	GetCpu() *float32
	SetCreatedTime(v string) *Application
	GetCreatedTime() *string
	SetCustomDNS(v *CustomDNS) *Application
	GetCustomDNS() *CustomDNS
	SetCustomDomainName(v string) *Application
	GetCustomDomainName() *string
	SetCustomHealthCheckConfig(v *CustomHealthCheckConfig) *Application
	GetCustomHealthCheckConfig() *CustomHealthCheckConfig
	SetCustomHostAlias(v *CustomHostAlias) *Application
	GetCustomHostAlias() *CustomHostAlias
	SetCustomRuntimeConfig(v *CustomRuntimeConfig) *Application
	GetCustomRuntimeConfig() *CustomRuntimeConfig
	SetDescription(v string) *Application
	GetDescription() *string
	SetDiskSize(v int32) *Application
	GetDiskSize() *int32
	SetEnableAppMetric(v bool) *Application
	GetEnableAppMetric() *bool
	SetEnableArmsAdvanced(v bool) *Application
	GetEnableArmsAdvanced() *bool
	SetEnvironmentVariables(v map[string]*string) *Application
	GetEnvironmentVariables() map[string]*string
	SetGpuMemorySize(v int32) *Application
	GetGpuMemorySize() *int32
	SetHandler(v string) *Application
	GetHandler() *string
	SetHttpTriggerConfig(v *HTTPTriggerConfig) *Application
	GetHttpTriggerConfig() *HTTPTriggerConfig
	SetImageConfig(v *ImageConfig) *Application
	GetImageConfig() *ImageConfig
	SetInitializationTimeout(v int32) *Application
	GetInitializationTimeout() *int32
	SetInitializer(v string) *Application
	GetInitializer() *string
	SetInstanceConcurrency(v int32) *Application
	GetInstanceConcurrency() *int32
	SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *Application
	GetInstanceLifecycleConfig() *InstanceLifecycleConfig
	SetInstanceSoftConcurrency(v int32) *Application
	GetInstanceSoftConcurrency() *int32
	SetInstanceType(v string) *Application
	GetInstanceType() *string
	SetInternetAccess(v bool) *Application
	GetInternetAccess() *bool
	SetLastModifiedTime(v string) *Application
	GetLastModifiedTime() *string
	SetLayers(v []*string) *Application
	GetLayers() []*string
	SetLayersArnV2(v []*string) *Application
	GetLayersArnV2() []*string
	SetLivenessProbe(v *Probe) *Application
	GetLivenessProbe() *Probe
	SetLogConfig(v *LogConfig) *Application
	GetLogConfig() *LogConfig
	SetMemorySize(v int32) *Application
	GetMemorySize() *int32
	SetNamespace(v string) *Application
	GetNamespace() *string
	SetNamespaceID(v string) *Application
	GetNamespaceID() *string
	SetNamespaceName(v string) *Application
	GetNamespaceName() *string
	SetNasConfig(v *NASConfig) *Application
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *Application
	GetOssMountConfig() *OSSMountConfig
	SetProgrammingLanguage(v string) *Application
	GetProgrammingLanguage() *string
	SetRuntime(v string) *Application
	GetRuntime() *string
	SetScaleConfig(v *ScaleConfig) *Application
	GetScaleConfig() *ScaleConfig
	SetSlsConfig(v *SLSConfig) *Application
	GetSlsConfig() *SLSConfig
	SetStartupProbe(v *Probe) *Application
	GetStartupProbe() *Probe
	SetTimeout(v int32) *Application
	GetTimeout() *int32
	SetTracingConfig(v *TracingConfig) *Application
	GetTracingConfig() *TracingConfig
	SetUrlInternet(v string) *Application
	GetUrlInternet() *string
	SetUrlIntranet(v string) *Application
	GetUrlIntranet() *string
	SetVersion(v *Version) *Application
	GetVersion() *Version
	SetVpcConfig(v *VPCConfig) *Application
	GetVpcConfig() *VPCConfig
}

type Application struct {
	RequestId               *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApplicationId           *string                  `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	ApplicationName         *string                  `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	Args                    *string                  `json:"args,omitempty" xml:"args,omitempty"`
	CaPort                  *int32                   `json:"caPort,omitempty" xml:"caPort,omitempty"`
	CodeChecksum            *string                  `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	CodeSize                *int64                   `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	Command                 *string                  `json:"command,omitempty" xml:"command,omitempty"`
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CreatedTime             *string                  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomDNS               *CustomDNS               `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomDomainName        *string                  `json:"customDomainName,omitempty" xml:"customDomainName,omitempty"`
	CustomHealthCheckConfig *CustomHealthCheckConfig `json:"customHealthCheckConfig,omitempty" xml:"customHealthCheckConfig,omitempty"`
	CustomHostAlias         *CustomHostAlias         `json:"customHostAlias,omitempty" xml:"customHostAlias,omitempty"`
	CustomRuntimeConfig     *CustomRuntimeConfig     `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnableAppMetric         *bool                    `json:"enableAppMetric,omitempty" xml:"enableAppMetric,omitempty"`
	EnableArmsAdvanced      *bool                    `json:"enableArmsAdvanced,omitempty" xml:"enableArmsAdvanced,omitempty"`
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
	LastModifiedTime        *string                  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Layers                  []*string                `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	LayersArnV2             []*string                `json:"layersArnV2,omitempty" xml:"layersArnV2,omitempty" type:"Repeated"`
	LivenessProbe           *Probe                   `json:"livenessProbe,omitempty" xml:"livenessProbe,omitempty"`
	LogConfig               *LogConfig               `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	MemorySize              *int32                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	Namespace               *string                  `json:"namespace,omitempty" xml:"namespace,omitempty"`
	NamespaceID             *string                  `json:"namespaceID,omitempty" xml:"namespaceID,omitempty"`
	NamespaceName           *string                  `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	NasConfig               *NASConfig               `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig          *OSSMountConfig          `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	ProgrammingLanguage     *string                  `json:"programmingLanguage,omitempty" xml:"programmingLanguage,omitempty"`
	Runtime                 *string                  `json:"runtime,omitempty" xml:"runtime,omitempty"`
	ScaleConfig             *ScaleConfig             `json:"scaleConfig,omitempty" xml:"scaleConfig,omitempty"`
	SlsConfig               *SLSConfig               `json:"slsConfig,omitempty" xml:"slsConfig,omitempty"`
	StartupProbe            *Probe                   `json:"startupProbe,omitempty" xml:"startupProbe,omitempty"`
	Timeout                 *int32                   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig           *TracingConfig           `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	UrlInternet             *string                  `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet             *string                  `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
	Version                 *Version                 `json:"version,omitempty" xml:"version,omitempty"`
	VpcConfig               *VPCConfig               `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s Application) String() string {
	return dara.Prettify(s)
}

func (s Application) GoString() string {
	return s.String()
}

func (s *Application) GetRequestId() *string {
	return s.RequestId
}

func (s *Application) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *Application) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *Application) GetArgs() *string {
	return s.Args
}

func (s *Application) GetCaPort() *int32 {
	return s.CaPort
}

func (s *Application) GetCodeChecksum() *string {
	return s.CodeChecksum
}

func (s *Application) GetCodeSize() *int64 {
	return s.CodeSize
}

func (s *Application) GetCommand() *string {
	return s.Command
}

func (s *Application) GetCpu() *float32 {
	return s.Cpu
}

func (s *Application) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Application) GetCustomDNS() *CustomDNS {
	return s.CustomDNS
}

func (s *Application) GetCustomDomainName() *string {
	return s.CustomDomainName
}

func (s *Application) GetCustomHealthCheckConfig() *CustomHealthCheckConfig {
	return s.CustomHealthCheckConfig
}

func (s *Application) GetCustomHostAlias() *CustomHostAlias {
	return s.CustomHostAlias
}

func (s *Application) GetCustomRuntimeConfig() *CustomRuntimeConfig {
	return s.CustomRuntimeConfig
}

func (s *Application) GetDescription() *string {
	return s.Description
}

func (s *Application) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *Application) GetEnableAppMetric() *bool {
	return s.EnableAppMetric
}

func (s *Application) GetEnableArmsAdvanced() *bool {
	return s.EnableArmsAdvanced
}

func (s *Application) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *Application) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *Application) GetHandler() *string {
	return s.Handler
}

func (s *Application) GetHttpTriggerConfig() *HTTPTriggerConfig {
	return s.HttpTriggerConfig
}

func (s *Application) GetImageConfig() *ImageConfig {
	return s.ImageConfig
}

func (s *Application) GetInitializationTimeout() *int32 {
	return s.InitializationTimeout
}

func (s *Application) GetInitializer() *string {
	return s.Initializer
}

func (s *Application) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *Application) GetInstanceLifecycleConfig() *InstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *Application) GetInstanceSoftConcurrency() *int32 {
	return s.InstanceSoftConcurrency
}

func (s *Application) GetInstanceType() *string {
	return s.InstanceType
}

func (s *Application) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *Application) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Application) GetLayers() []*string {
	return s.Layers
}

func (s *Application) GetLayersArnV2() []*string {
	return s.LayersArnV2
}

func (s *Application) GetLivenessProbe() *Probe {
	return s.LivenessProbe
}

func (s *Application) GetLogConfig() *LogConfig {
	return s.LogConfig
}

func (s *Application) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *Application) GetNamespace() *string {
	return s.Namespace
}

func (s *Application) GetNamespaceID() *string {
	return s.NamespaceID
}

func (s *Application) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *Application) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *Application) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *Application) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *Application) GetRuntime() *string {
	return s.Runtime
}

func (s *Application) GetScaleConfig() *ScaleConfig {
	return s.ScaleConfig
}

func (s *Application) GetSlsConfig() *SLSConfig {
	return s.SlsConfig
}

func (s *Application) GetStartupProbe() *Probe {
	return s.StartupProbe
}

func (s *Application) GetTimeout() *int32 {
	return s.Timeout
}

func (s *Application) GetTracingConfig() *TracingConfig {
	return s.TracingConfig
}

func (s *Application) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *Application) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *Application) GetVersion() *Version {
	return s.Version
}

func (s *Application) GetVpcConfig() *VPCConfig {
	return s.VpcConfig
}

func (s *Application) SetRequestId(v string) *Application {
	s.RequestId = &v
	return s
}

func (s *Application) SetApplicationId(v string) *Application {
	s.ApplicationId = &v
	return s
}

func (s *Application) SetApplicationName(v string) *Application {
	s.ApplicationName = &v
	return s
}

func (s *Application) SetArgs(v string) *Application {
	s.Args = &v
	return s
}

func (s *Application) SetCaPort(v int32) *Application {
	s.CaPort = &v
	return s
}

func (s *Application) SetCodeChecksum(v string) *Application {
	s.CodeChecksum = &v
	return s
}

func (s *Application) SetCodeSize(v int64) *Application {
	s.CodeSize = &v
	return s
}

func (s *Application) SetCommand(v string) *Application {
	s.Command = &v
	return s
}

func (s *Application) SetCpu(v float32) *Application {
	s.Cpu = &v
	return s
}

func (s *Application) SetCreatedTime(v string) *Application {
	s.CreatedTime = &v
	return s
}

func (s *Application) SetCustomDNS(v *CustomDNS) *Application {
	s.CustomDNS = v
	return s
}

func (s *Application) SetCustomDomainName(v string) *Application {
	s.CustomDomainName = &v
	return s
}

func (s *Application) SetCustomHealthCheckConfig(v *CustomHealthCheckConfig) *Application {
	s.CustomHealthCheckConfig = v
	return s
}

func (s *Application) SetCustomHostAlias(v *CustomHostAlias) *Application {
	s.CustomHostAlias = v
	return s
}

func (s *Application) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *Application {
	s.CustomRuntimeConfig = v
	return s
}

func (s *Application) SetDescription(v string) *Application {
	s.Description = &v
	return s
}

func (s *Application) SetDiskSize(v int32) *Application {
	s.DiskSize = &v
	return s
}

func (s *Application) SetEnableAppMetric(v bool) *Application {
	s.EnableAppMetric = &v
	return s
}

func (s *Application) SetEnableArmsAdvanced(v bool) *Application {
	s.EnableArmsAdvanced = &v
	return s
}

func (s *Application) SetEnvironmentVariables(v map[string]*string) *Application {
	s.EnvironmentVariables = v
	return s
}

func (s *Application) SetGpuMemorySize(v int32) *Application {
	s.GpuMemorySize = &v
	return s
}

func (s *Application) SetHandler(v string) *Application {
	s.Handler = &v
	return s
}

func (s *Application) SetHttpTriggerConfig(v *HTTPTriggerConfig) *Application {
	s.HttpTriggerConfig = v
	return s
}

func (s *Application) SetImageConfig(v *ImageConfig) *Application {
	s.ImageConfig = v
	return s
}

func (s *Application) SetInitializationTimeout(v int32) *Application {
	s.InitializationTimeout = &v
	return s
}

func (s *Application) SetInitializer(v string) *Application {
	s.Initializer = &v
	return s
}

func (s *Application) SetInstanceConcurrency(v int32) *Application {
	s.InstanceConcurrency = &v
	return s
}

func (s *Application) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *Application {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *Application) SetInstanceSoftConcurrency(v int32) *Application {
	s.InstanceSoftConcurrency = &v
	return s
}

func (s *Application) SetInstanceType(v string) *Application {
	s.InstanceType = &v
	return s
}

func (s *Application) SetInternetAccess(v bool) *Application {
	s.InternetAccess = &v
	return s
}

func (s *Application) SetLastModifiedTime(v string) *Application {
	s.LastModifiedTime = &v
	return s
}

func (s *Application) SetLayers(v []*string) *Application {
	s.Layers = v
	return s
}

func (s *Application) SetLayersArnV2(v []*string) *Application {
	s.LayersArnV2 = v
	return s
}

func (s *Application) SetLivenessProbe(v *Probe) *Application {
	s.LivenessProbe = v
	return s
}

func (s *Application) SetLogConfig(v *LogConfig) *Application {
	s.LogConfig = v
	return s
}

func (s *Application) SetMemorySize(v int32) *Application {
	s.MemorySize = &v
	return s
}

func (s *Application) SetNamespace(v string) *Application {
	s.Namespace = &v
	return s
}

func (s *Application) SetNamespaceID(v string) *Application {
	s.NamespaceID = &v
	return s
}

func (s *Application) SetNamespaceName(v string) *Application {
	s.NamespaceName = &v
	return s
}

func (s *Application) SetNasConfig(v *NASConfig) *Application {
	s.NasConfig = v
	return s
}

func (s *Application) SetOssMountConfig(v *OSSMountConfig) *Application {
	s.OssMountConfig = v
	return s
}

func (s *Application) SetProgrammingLanguage(v string) *Application {
	s.ProgrammingLanguage = &v
	return s
}

func (s *Application) SetRuntime(v string) *Application {
	s.Runtime = &v
	return s
}

func (s *Application) SetScaleConfig(v *ScaleConfig) *Application {
	s.ScaleConfig = v
	return s
}

func (s *Application) SetSlsConfig(v *SLSConfig) *Application {
	s.SlsConfig = v
	return s
}

func (s *Application) SetStartupProbe(v *Probe) *Application {
	s.StartupProbe = v
	return s
}

func (s *Application) SetTimeout(v int32) *Application {
	s.Timeout = &v
	return s
}

func (s *Application) SetTracingConfig(v *TracingConfig) *Application {
	s.TracingConfig = v
	return s
}

func (s *Application) SetUrlInternet(v string) *Application {
	s.UrlInternet = &v
	return s
}

func (s *Application) SetUrlIntranet(v string) *Application {
	s.UrlIntranet = &v
	return s
}

func (s *Application) SetVersion(v *Version) *Application {
	s.Version = v
	return s
}

func (s *Application) SetVpcConfig(v *VPCConfig) *Application {
	s.VpcConfig = v
	return s
}

func (s *Application) Validate() error {
	if s.CustomDNS != nil {
		if err := s.CustomDNS.Validate(); err != nil {
			return err
		}
	}
	if s.CustomHealthCheckConfig != nil {
		if err := s.CustomHealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CustomHostAlias != nil {
		if err := s.CustomHostAlias.Validate(); err != nil {
			return err
		}
	}
	if s.CustomRuntimeConfig != nil {
		if err := s.CustomRuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HttpTriggerConfig != nil {
		if err := s.HttpTriggerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ImageConfig != nil {
		if err := s.ImageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceLifecycleConfig != nil {
		if err := s.InstanceLifecycleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LivenessProbe != nil {
		if err := s.LivenessProbe.Validate(); err != nil {
			return err
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
	if s.ScaleConfig != nil {
		if err := s.ScaleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SlsConfig != nil {
		if err := s.SlsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StartupProbe != nil {
		if err := s.StartupProbe.Validate(); err != nil {
			return err
		}
	}
	if s.TracingConfig != nil {
		if err := s.TracingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Version != nil {
		if err := s.Version.Validate(); err != nil {
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainer interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v string) *Container
	GetArgs() *string
	SetCommand(v string) *Container
	GetCommand() *string
	SetEnvironmentVariables(v map[string]*string) *Container
	GetEnvironmentVariables() map[string]*string
	SetImage(v string) *Container
	GetImage() *string
	SetImageRegistryConfig(v *ImageRegistryConfig) *Container
	GetImageRegistryConfig() *ImageRegistryConfig
	SetMetricsCollectConfig(v *MetricsCollectConfig) *Container
	GetMetricsCollectConfig() *MetricsCollectConfig
	SetPort(v int32) *Container
	GetPort() *int32
	SetRequestConcurrency(v int32) *Container
	GetRequestConcurrency() *int32
	SetRequestTimeout(v int32) *Container
	GetRequestTimeout() *int32
	SetResources(v *ContainerResources) *Container
	GetResources() *ContainerResources
	SetSLSCollectConfigs(v *SLSCollectConfigs) *Container
	GetSLSCollectConfigs() *SLSCollectConfigs
	SetStartupProbe(v *StartupProbe) *Container
	GetStartupProbe() *StartupProbe
	SetWebNASConfig(v *WebNASConfig) *Container
	GetWebNASConfig() *WebNASConfig
	SetWebOSSConfig(v *WebOSSConfig) *Container
	GetWebOSSConfig() *WebOSSConfig
}

type Container struct {
	// example:
	//
	// ["abc", ">", "file0"]
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// ["/bin/sh"]
	Command              *string            `json:"Command,omitempty" xml:"Command,omitempty"`
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/serverless_devsxxxxx
	Image                *string               `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageRegistryConfig  *ImageRegistryConfig  `json:"ImageRegistryConfig,omitempty" xml:"ImageRegistryConfig,omitempty"`
	MetricsCollectConfig *MetricsCollectConfig `json:"MetricsCollectConfig,omitempty" xml:"MetricsCollectConfig,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 100
	RequestConcurrency *int32 `json:"RequestConcurrency,omitempty" xml:"RequestConcurrency,omitempty"`
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// This parameter is required.
	Resources         *ContainerResources `json:"Resources,omitempty" xml:"Resources,omitempty"`
	SLSCollectConfigs *SLSCollectConfigs  `json:"SLSCollectConfigs,omitempty" xml:"SLSCollectConfigs,omitempty"`
	StartupProbe      *StartupProbe       `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	WebNASConfig      *WebNASConfig       `json:"WebNASConfig,omitempty" xml:"WebNASConfig,omitempty"`
	WebOSSConfig      *WebOSSConfig       `json:"WebOSSConfig,omitempty" xml:"WebOSSConfig,omitempty"`
}

func (s Container) String() string {
	return dara.Prettify(s)
}

func (s Container) GoString() string {
	return s.String()
}

func (s *Container) GetArgs() *string {
	return s.Args
}

func (s *Container) GetCommand() *string {
	return s.Command
}

func (s *Container) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *Container) GetImage() *string {
	return s.Image
}

func (s *Container) GetImageRegistryConfig() *ImageRegistryConfig {
	return s.ImageRegistryConfig
}

func (s *Container) GetMetricsCollectConfig() *MetricsCollectConfig {
	return s.MetricsCollectConfig
}

func (s *Container) GetPort() *int32 {
	return s.Port
}

func (s *Container) GetRequestConcurrency() *int32 {
	return s.RequestConcurrency
}

func (s *Container) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *Container) GetResources() *ContainerResources {
	return s.Resources
}

func (s *Container) GetSLSCollectConfigs() *SLSCollectConfigs {
	return s.SLSCollectConfigs
}

func (s *Container) GetStartupProbe() *StartupProbe {
	return s.StartupProbe
}

func (s *Container) GetWebNASConfig() *WebNASConfig {
	return s.WebNASConfig
}

func (s *Container) GetWebOSSConfig() *WebOSSConfig {
	return s.WebOSSConfig
}

func (s *Container) SetArgs(v string) *Container {
	s.Args = &v
	return s
}

func (s *Container) SetCommand(v string) *Container {
	s.Command = &v
	return s
}

func (s *Container) SetEnvironmentVariables(v map[string]*string) *Container {
	s.EnvironmentVariables = v
	return s
}

func (s *Container) SetImage(v string) *Container {
	s.Image = &v
	return s
}

func (s *Container) SetImageRegistryConfig(v *ImageRegistryConfig) *Container {
	s.ImageRegistryConfig = v
	return s
}

func (s *Container) SetMetricsCollectConfig(v *MetricsCollectConfig) *Container {
	s.MetricsCollectConfig = v
	return s
}

func (s *Container) SetPort(v int32) *Container {
	s.Port = &v
	return s
}

func (s *Container) SetRequestConcurrency(v int32) *Container {
	s.RequestConcurrency = &v
	return s
}

func (s *Container) SetRequestTimeout(v int32) *Container {
	s.RequestTimeout = &v
	return s
}

func (s *Container) SetResources(v *ContainerResources) *Container {
	s.Resources = v
	return s
}

func (s *Container) SetSLSCollectConfigs(v *SLSCollectConfigs) *Container {
	s.SLSCollectConfigs = v
	return s
}

func (s *Container) SetStartupProbe(v *StartupProbe) *Container {
	s.StartupProbe = v
	return s
}

func (s *Container) SetWebNASConfig(v *WebNASConfig) *Container {
	s.WebNASConfig = v
	return s
}

func (s *Container) SetWebOSSConfig(v *WebOSSConfig) *Container {
	s.WebOSSConfig = v
	return s
}

func (s *Container) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomContainerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationInfo(v *AccelerationInfo) *CustomContainerConfig
	GetAccelerationInfo() *AccelerationInfo
	SetAccelerationType(v string) *CustomContainerConfig
	GetAccelerationType() *string
	SetAcrInstanceId(v string) *CustomContainerConfig
	GetAcrInstanceId() *string
	SetCommand(v []*string) *CustomContainerConfig
	GetCommand() []*string
	SetEntrypoint(v []*string) *CustomContainerConfig
	GetEntrypoint() []*string
	SetHealthCheckConfig(v *CustomHealthCheckConfig) *CustomContainerConfig
	GetHealthCheckConfig() *CustomHealthCheckConfig
	SetImage(v string) *CustomContainerConfig
	GetImage() *string
	SetPort(v int32) *CustomContainerConfig
	GetPort() *int32
	SetRegistryConfig(v *RegistryConfig) *CustomContainerConfig
	GetRegistryConfig() *RegistryConfig
	SetResolvedImageUri(v string) *CustomContainerConfig
	GetResolvedImageUri() *string
}

type CustomContainerConfig struct {
	AccelerationInfo *AccelerationInfo `json:"accelerationInfo,omitempty" xml:"accelerationInfo,omitempty"`
	// example:
	//
	// deprecated
	AccelerationType *string `json:"accelerationType,omitempty" xml:"accelerationType,omitempty"`
	// example:
	//
	// deprecated
	AcrInstanceId     *string                  `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	Command           []*string                `json:"command" xml:"command" type:"Repeated"`
	Entrypoint        []*string                `json:"entrypoint" xml:"entrypoint" type:"Repeated"`
	HealthCheckConfig *CustomHealthCheckConfig `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/fc-demo/helloworld:v1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// example:
	//
	// 9000
	Port           *int32          `json:"port,omitempty" xml:"port,omitempty"`
	RegistryConfig *RegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// example:
	//
	// stand-sh-registry-vpc.cn-shanghai.cr.aliyuncs.com/fc-demo2/springboot-helloworld@sha256:68d1****0d64d6
	ResolvedImageUri *string `json:"resolvedImageUri,omitempty" xml:"resolvedImageUri,omitempty"`
}

func (s CustomContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s CustomContainerConfig) GoString() string {
	return s.String()
}

func (s *CustomContainerConfig) GetAccelerationInfo() *AccelerationInfo {
	return s.AccelerationInfo
}

func (s *CustomContainerConfig) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *CustomContainerConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CustomContainerConfig) GetCommand() []*string {
	return s.Command
}

func (s *CustomContainerConfig) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *CustomContainerConfig) GetHealthCheckConfig() *CustomHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *CustomContainerConfig) GetImage() *string {
	return s.Image
}

func (s *CustomContainerConfig) GetPort() *int32 {
	return s.Port
}

func (s *CustomContainerConfig) GetRegistryConfig() *RegistryConfig {
	return s.RegistryConfig
}

func (s *CustomContainerConfig) GetResolvedImageUri() *string {
	return s.ResolvedImageUri
}

func (s *CustomContainerConfig) SetAccelerationInfo(v *AccelerationInfo) *CustomContainerConfig {
	s.AccelerationInfo = v
	return s
}

func (s *CustomContainerConfig) SetAccelerationType(v string) *CustomContainerConfig {
	s.AccelerationType = &v
	return s
}

func (s *CustomContainerConfig) SetAcrInstanceId(v string) *CustomContainerConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *CustomContainerConfig) SetCommand(v []*string) *CustomContainerConfig {
	s.Command = v
	return s
}

func (s *CustomContainerConfig) SetEntrypoint(v []*string) *CustomContainerConfig {
	s.Entrypoint = v
	return s
}

func (s *CustomContainerConfig) SetHealthCheckConfig(v *CustomHealthCheckConfig) *CustomContainerConfig {
	s.HealthCheckConfig = v
	return s
}

func (s *CustomContainerConfig) SetImage(v string) *CustomContainerConfig {
	s.Image = &v
	return s
}

func (s *CustomContainerConfig) SetPort(v int32) *CustomContainerConfig {
	s.Port = &v
	return s
}

func (s *CustomContainerConfig) SetRegistryConfig(v *RegistryConfig) *CustomContainerConfig {
	s.RegistryConfig = v
	return s
}

func (s *CustomContainerConfig) SetResolvedImageUri(v string) *CustomContainerConfig {
	s.ResolvedImageUri = &v
	return s
}

func (s *CustomContainerConfig) Validate() error {
	if s.AccelerationInfo != nil {
		if err := s.AccelerationInfo.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

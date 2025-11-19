// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProxy interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v float32) *ModelProxy
	GetCpu() *float32
	SetCreatedAt(v string) *ModelProxy
	GetCreatedAt() *string
	SetCredentialName(v string) *ModelProxy
	GetCredentialName() *string
	SetDescription(v string) *ModelProxy
	GetDescription() *string
	SetEndpoint(v string) *ModelProxy
	GetEndpoint() *string
	SetFunctionName(v string) *ModelProxy
	GetFunctionName() *string
	SetLastUpdatedAt(v string) *ModelProxy
	GetLastUpdatedAt() *string
	SetLitellmVersion(v string) *ModelProxy
	GetLitellmVersion() *string
	SetLogConfiguration(v *LogConfiguration) *ModelProxy
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *ModelProxy
	GetMemory() *int32
	SetModelProxyId(v string) *ModelProxy
	GetModelProxyId() *string
	SetModelProxyName(v string) *ModelProxy
	GetModelProxyName() *string
	SetModelType(v string) *ModelProxy
	GetModelType() *string
	SetNetworkConfiguration(v *NetworkConfiguration) *ModelProxy
	GetNetworkConfiguration() *NetworkConfiguration
	SetProxyConfig(v *ProxyConfig) *ModelProxy
	GetProxyConfig() *ProxyConfig
	SetProxyMode(v string) *ModelProxy
	GetProxyMode() *string
	SetServiceRegionId(v string) *ModelProxy
	GetServiceRegionId() *string
	SetStatus(v string) *ModelProxy
	GetStatus() *string
	SetStatusReason(v string) *ModelProxy
	GetStatusReason() *string
}

type ModelProxy struct {
	Cpu                  *float32              `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CreatedAt            *string               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialName       *string               `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	Endpoint             *string               `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FunctionName         *string               `json:"functionName,omitempty" xml:"functionName,omitempty"`
	LastUpdatedAt        *string               `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	LitellmVersion       *string               `json:"litellmVersion,omitempty" xml:"litellmVersion,omitempty"`
	LogConfiguration     *LogConfiguration     `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	Memory               *int32                `json:"memory,omitempty" xml:"memory,omitempty"`
	ModelProxyId         *string               `json:"modelProxyId,omitempty" xml:"modelProxyId,omitempty"`
	ModelProxyName       *string               `json:"modelProxyName,omitempty" xml:"modelProxyName,omitempty"`
	ModelType            *string               `json:"modelType,omitempty" xml:"modelType,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	ProxyConfig          *ProxyConfig          `json:"proxyConfig,omitempty" xml:"proxyConfig,omitempty"`
	ProxyMode            *string               `json:"proxyMode,omitempty" xml:"proxyMode,omitempty"`
	ServiceRegionId      *string               `json:"serviceRegionId,omitempty" xml:"serviceRegionId,omitempty"`
	Status               *string               `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason         *string               `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
}

func (s ModelProxy) String() string {
	return dara.Prettify(s)
}

func (s ModelProxy) GoString() string {
	return s.String()
}

func (s *ModelProxy) GetCpu() *float32 {
	return s.Cpu
}

func (s *ModelProxy) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ModelProxy) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ModelProxy) GetDescription() *string {
	return s.Description
}

func (s *ModelProxy) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModelProxy) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ModelProxy) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *ModelProxy) GetLitellmVersion() *string {
	return s.LitellmVersion
}

func (s *ModelProxy) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *ModelProxy) GetMemory() *int32 {
	return s.Memory
}

func (s *ModelProxy) GetModelProxyId() *string {
	return s.ModelProxyId
}

func (s *ModelProxy) GetModelProxyName() *string {
	return s.ModelProxyName
}

func (s *ModelProxy) GetModelType() *string {
	return s.ModelType
}

func (s *ModelProxy) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *ModelProxy) GetProxyConfig() *ProxyConfig {
	return s.ProxyConfig
}

func (s *ModelProxy) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *ModelProxy) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ModelProxy) GetStatus() *string {
	return s.Status
}

func (s *ModelProxy) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ModelProxy) SetCpu(v float32) *ModelProxy {
	s.Cpu = &v
	return s
}

func (s *ModelProxy) SetCreatedAt(v string) *ModelProxy {
	s.CreatedAt = &v
	return s
}

func (s *ModelProxy) SetCredentialName(v string) *ModelProxy {
	s.CredentialName = &v
	return s
}

func (s *ModelProxy) SetDescription(v string) *ModelProxy {
	s.Description = &v
	return s
}

func (s *ModelProxy) SetEndpoint(v string) *ModelProxy {
	s.Endpoint = &v
	return s
}

func (s *ModelProxy) SetFunctionName(v string) *ModelProxy {
	s.FunctionName = &v
	return s
}

func (s *ModelProxy) SetLastUpdatedAt(v string) *ModelProxy {
	s.LastUpdatedAt = &v
	return s
}

func (s *ModelProxy) SetLitellmVersion(v string) *ModelProxy {
	s.LitellmVersion = &v
	return s
}

func (s *ModelProxy) SetLogConfiguration(v *LogConfiguration) *ModelProxy {
	s.LogConfiguration = v
	return s
}

func (s *ModelProxy) SetMemory(v int32) *ModelProxy {
	s.Memory = &v
	return s
}

func (s *ModelProxy) SetModelProxyId(v string) *ModelProxy {
	s.ModelProxyId = &v
	return s
}

func (s *ModelProxy) SetModelProxyName(v string) *ModelProxy {
	s.ModelProxyName = &v
	return s
}

func (s *ModelProxy) SetModelType(v string) *ModelProxy {
	s.ModelType = &v
	return s
}

func (s *ModelProxy) SetNetworkConfiguration(v *NetworkConfiguration) *ModelProxy {
	s.NetworkConfiguration = v
	return s
}

func (s *ModelProxy) SetProxyConfig(v *ProxyConfig) *ModelProxy {
	s.ProxyConfig = v
	return s
}

func (s *ModelProxy) SetProxyMode(v string) *ModelProxy {
	s.ProxyMode = &v
	return s
}

func (s *ModelProxy) SetServiceRegionId(v string) *ModelProxy {
	s.ServiceRegionId = &v
	return s
}

func (s *ModelProxy) SetStatus(v string) *ModelProxy {
	s.Status = &v
	return s
}

func (s *ModelProxy) SetStatusReason(v string) *ModelProxy {
	s.StatusReason = &v
	return s
}

func (s *ModelProxy) Validate() error {
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ProxyConfig != nil {
		if err := s.ProxyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

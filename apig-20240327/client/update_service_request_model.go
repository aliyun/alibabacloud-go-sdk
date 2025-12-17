// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *UpdateServiceRequest
	GetAddresses() []*string
	SetAgentServiceConfig(v *AgentServiceConfig) *UpdateServiceRequest
	GetAgentServiceConfig() *AgentServiceConfig
	SetAiServiceConfig(v *AiServiceConfig) *UpdateServiceRequest
	GetAiServiceConfig() *AiServiceConfig
	SetDnsServers(v []*string) *UpdateServiceRequest
	GetDnsServers() []*string
	SetHealthCheckConfig(v *UpdateServiceRequestHealthCheckConfig) *UpdateServiceRequest
	GetHealthCheckConfig() *UpdateServiceRequestHealthCheckConfig
	SetHealthyPanicThreshold(v float32) *UpdateServiceRequest
	GetHealthyPanicThreshold() *float32
	SetOutlierDetectionConfig(v *UpdateServiceRequestOutlierDetectionConfig) *UpdateServiceRequest
	GetOutlierDetectionConfig() *UpdateServiceRequestOutlierDetectionConfig
	SetPorts(v []*UpdateServiceRequestPorts) *UpdateServiceRequest
	GetPorts() []*UpdateServiceRequestPorts
	SetProtocol(v string) *UpdateServiceRequest
	GetProtocol() *string
}

type UpdateServiceRequest struct {
	Addresses              []*string                                   `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	AgentServiceConfig     *AgentServiceConfig                         `json:"agentServiceConfig,omitempty" xml:"agentServiceConfig,omitempty"`
	AiServiceConfig        *AiServiceConfig                            `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	DnsServers             []*string                                   `json:"dnsServers,omitempty" xml:"dnsServers,omitempty" type:"Repeated"`
	HealthCheckConfig      *UpdateServiceRequestHealthCheckConfig      `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty" type:"Struct"`
	HealthyPanicThreshold  *float32                                    `json:"healthyPanicThreshold,omitempty" xml:"healthyPanicThreshold,omitempty"`
	OutlierDetectionConfig *UpdateServiceRequestOutlierDetectionConfig `json:"outlierDetectionConfig,omitempty" xml:"outlierDetectionConfig,omitempty" type:"Struct"`
	Ports                  []*UpdateServiceRequestPorts                `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdateServiceRequest) GetAgentServiceConfig() *AgentServiceConfig {
	return s.AgentServiceConfig
}

func (s *UpdateServiceRequest) GetAiServiceConfig() *AiServiceConfig {
	return s.AiServiceConfig
}

func (s *UpdateServiceRequest) GetDnsServers() []*string {
	return s.DnsServers
}

func (s *UpdateServiceRequest) GetHealthCheckConfig() *UpdateServiceRequestHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *UpdateServiceRequest) GetHealthyPanicThreshold() *float32 {
	return s.HealthyPanicThreshold
}

func (s *UpdateServiceRequest) GetOutlierDetectionConfig() *UpdateServiceRequestOutlierDetectionConfig {
	return s.OutlierDetectionConfig
}

func (s *UpdateServiceRequest) GetPorts() []*UpdateServiceRequestPorts {
	return s.Ports
}

func (s *UpdateServiceRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateServiceRequest) SetAddresses(v []*string) *UpdateServiceRequest {
	s.Addresses = v
	return s
}

func (s *UpdateServiceRequest) SetAgentServiceConfig(v *AgentServiceConfig) *UpdateServiceRequest {
	s.AgentServiceConfig = v
	return s
}

func (s *UpdateServiceRequest) SetAiServiceConfig(v *AiServiceConfig) *UpdateServiceRequest {
	s.AiServiceConfig = v
	return s
}

func (s *UpdateServiceRequest) SetDnsServers(v []*string) *UpdateServiceRequest {
	s.DnsServers = v
	return s
}

func (s *UpdateServiceRequest) SetHealthCheckConfig(v *UpdateServiceRequestHealthCheckConfig) *UpdateServiceRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *UpdateServiceRequest) SetHealthyPanicThreshold(v float32) *UpdateServiceRequest {
	s.HealthyPanicThreshold = &v
	return s
}

func (s *UpdateServiceRequest) SetOutlierDetectionConfig(v *UpdateServiceRequestOutlierDetectionConfig) *UpdateServiceRequest {
	s.OutlierDetectionConfig = v
	return s
}

func (s *UpdateServiceRequest) SetPorts(v []*UpdateServiceRequestPorts) *UpdateServiceRequest {
	s.Ports = v
	return s
}

func (s *UpdateServiceRequest) SetProtocol(v string) *UpdateServiceRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateServiceRequest) Validate() error {
	if s.AgentServiceConfig != nil {
		if err := s.AgentServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiServiceConfig != nil {
		if err := s.AiServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OutlierDetectionConfig != nil {
		if err := s.OutlierDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateServiceRequestHealthCheckConfig struct {
	// example:
	//
	// true
	Enable           *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	ExpectedStatuses []*string `json:"expectedStatuses,omitempty" xml:"expectedStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	// example:
	//
	// dev.itemcener.com
	HttpHost *string `json:"httpHost,omitempty" xml:"httpHost,omitempty"`
	// example:
	//
	// /healthz
	HttpPath *string `json:"httpPath,omitempty" xml:"httpPath,omitempty"`
	// example:
	//
	// 2
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// 2
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// 22
	UnhealthyThreshold *int32 `json:"unhealthyThreshold,omitempty" xml:"unhealthyThreshold,omitempty"`
}

func (s UpdateServiceRequestHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestHealthCheckConfig) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateServiceRequestHealthCheckConfig) GetExpectedStatuses() []*string {
	return s.ExpectedStatuses
}

func (s *UpdateServiceRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateServiceRequestHealthCheckConfig) GetHttpHost() *string {
	return s.HttpHost
}

func (s *UpdateServiceRequestHealthCheckConfig) GetHttpPath() *string {
	return s.HttpPath
}

func (s *UpdateServiceRequestHealthCheckConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateServiceRequestHealthCheckConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateServiceRequestHealthCheckConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateServiceRequestHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateServiceRequestHealthCheckConfig) SetEnable(v bool) *UpdateServiceRequestHealthCheckConfig {
	s.Enable = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetExpectedStatuses(v []*string) *UpdateServiceRequestHealthCheckConfig {
	s.ExpectedStatuses = v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetHealthyThreshold(v int32) *UpdateServiceRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetHttpHost(v string) *UpdateServiceRequestHealthCheckConfig {
	s.HttpHost = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetHttpPath(v string) *UpdateServiceRequestHealthCheckConfig {
	s.HttpPath = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetInterval(v int32) *UpdateServiceRequestHealthCheckConfig {
	s.Interval = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetProtocol(v string) *UpdateServiceRequestHealthCheckConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetTimeout(v int32) *UpdateServiceRequestHealthCheckConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *UpdateServiceRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateServiceRequestHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestOutlierDetectionConfig struct {
	BaseEjectionTime              *int32 `json:"baseEjectionTime,omitempty" xml:"baseEjectionTime,omitempty"`
	Enable                        *bool  `json:"enable,omitempty" xml:"enable,omitempty"`
	FailurePercentageMinimumHosts *int32 `json:"failurePercentageMinimumHosts,omitempty" xml:"failurePercentageMinimumHosts,omitempty"`
	FailurePercentageThreshold    *int32 `json:"failurePercentageThreshold,omitempty" xml:"failurePercentageThreshold,omitempty"`
	Interval                      *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
}

func (s UpdateServiceRequestOutlierDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestOutlierDetectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestOutlierDetectionConfig) GetBaseEjectionTime() *int32 {
	return s.BaseEjectionTime
}

func (s *UpdateServiceRequestOutlierDetectionConfig) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateServiceRequestOutlierDetectionConfig) GetFailurePercentageMinimumHosts() *int32 {
	return s.FailurePercentageMinimumHosts
}

func (s *UpdateServiceRequestOutlierDetectionConfig) GetFailurePercentageThreshold() *int32 {
	return s.FailurePercentageThreshold
}

func (s *UpdateServiceRequestOutlierDetectionConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateServiceRequestOutlierDetectionConfig) SetBaseEjectionTime(v int32) *UpdateServiceRequestOutlierDetectionConfig {
	s.BaseEjectionTime = &v
	return s
}

func (s *UpdateServiceRequestOutlierDetectionConfig) SetEnable(v bool) *UpdateServiceRequestOutlierDetectionConfig {
	s.Enable = &v
	return s
}

func (s *UpdateServiceRequestOutlierDetectionConfig) SetFailurePercentageMinimumHosts(v int32) *UpdateServiceRequestOutlierDetectionConfig {
	s.FailurePercentageMinimumHosts = &v
	return s
}

func (s *UpdateServiceRequestOutlierDetectionConfig) SetFailurePercentageThreshold(v int32) *UpdateServiceRequestOutlierDetectionConfig {
	s.FailurePercentageThreshold = &v
	return s
}

func (s *UpdateServiceRequestOutlierDetectionConfig) SetInterval(v int32) *UpdateServiceRequestOutlierDetectionConfig {
	s.Interval = &v
	return s
}

func (s *UpdateServiceRequestOutlierDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestPorts struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Port     *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateServiceRequestPorts) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestPorts) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestPorts) GetName() *string {
	return s.Name
}

func (s *UpdateServiceRequestPorts) GetPort() *int32 {
	return s.Port
}

func (s *UpdateServiceRequestPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateServiceRequestPorts) SetName(v string) *UpdateServiceRequestPorts {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestPorts) SetPort(v int32) *UpdateServiceRequestPorts {
	s.Port = &v
	return s
}

func (s *UpdateServiceRequestPorts) SetProtocol(v string) *UpdateServiceRequestPorts {
	s.Protocol = &v
	return s
}

func (s *UpdateServiceRequestPorts) Validate() error {
	return dara.Validate(s)
}

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
	SetModelProviderId(v string) *UpdateServiceRequest
	GetModelProviderId() *string
	SetOutlierDetectionConfig(v *UpdateServiceRequestOutlierDetectionConfig) *UpdateServiceRequest
	GetOutlierDetectionConfig() *UpdateServiceRequestOutlierDetectionConfig
	SetPorts(v []*UpdateServiceRequestPorts) *UpdateServiceRequest
	GetPorts() []*UpdateServiceRequestPorts
	SetProtocol(v string) *UpdateServiceRequest
	GetProtocol() *string
}

type UpdateServiceRequest struct {
	// The list of domain names or fixed addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The agent service configuration.
	AgentServiceConfig *AgentServiceConfig `json:"agentServiceConfig,omitempty" xml:"agentServiceConfig,omitempty"`
	// The AI service configuration.
	AiServiceConfig *AiServiceConfig `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	// The DNS server addresses.
	DnsServers []*string `json:"dnsServers,omitempty" xml:"dnsServers,omitempty" type:"Repeated"`
	// The health check configuration of the service.
	HealthCheckConfig *UpdateServiceRequestHealthCheckConfig `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty" type:"Struct"`
	// The health check threshold.
	//
	// example:
	//
	// 80
	HealthyPanicThreshold *float32 `json:"healthyPanicThreshold,omitempty" xml:"healthyPanicThreshold,omitempty"`
	ModelProviderId       *string  `json:"modelProviderId,omitempty" xml:"modelProviderId,omitempty"`
	// The passive health check parameter settings.
	OutlierDetectionConfig *UpdateServiceRequestOutlierDetectionConfig `json:"outlierDetectionConfig,omitempty" xml:"outlierDetectionConfig,omitempty" type:"Struct"`
	// The port information.
	Ports []*UpdateServiceRequestPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// The protocol of the service.
	//
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

func (s *UpdateServiceRequest) GetModelProviderId() *string {
	return s.ModelProviderId
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

func (s *UpdateServiceRequest) SetModelProviderId(v string) *UpdateServiceRequest {
	s.ModelProviderId = &v
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
	// Specifies whether to enable health checks for the service.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of expected HTTP status codes that indicate a healthy response. This parameter is required when the protocol is HTTP.
	ExpectedStatuses []*string `json:"expectedStatuses,omitempty" xml:"expectedStatuses,omitempty" type:"Repeated"`
	// The healthy threshold for health checks.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	// The domain name for health checks. This parameter is optional and can be configured when the protocol is HTTP.
	//
	// example:
	//
	// dev.itemcener.com
	HttpHost *string `json:"httpHost,omitempty" xml:"httpHost,omitempty"`
	// The request path for health checks. This parameter is required when the protocol is HTTP.
	//
	// example:
	//
	// /healthz
	HttpPath *string `json:"httpPath,omitempty" xml:"httpPath,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 2
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// The protocol used for health checks.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The response timeout period for health checks. Unit: seconds.
	//
	// example:
	//
	// 2
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The unhealthy threshold for health checks.
	//
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
	// The base ejection time, which is the initial isolation duration after a node is ejected (for example, 30 seconds). The isolation time is calculated by using the following formula: k × base_ejection_time (the initial value of k is 1). Each ejection increases the isolation time (k is incremented by 1). If consecutive checks are healthy, the isolation time is gradually reduced (k is decremented by 1).
	//
	// example:
	//
	// 30
	BaseEjectionTime *int32 `json:"baseEjectionTime,omitempty" xml:"baseEjectionTime,omitempty"`
	// enable
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The panic threshold.
	//
	// When the proportion of healthy nodes in the service is greater than the panic threshold, health checks function normally. Requests are sent only to healthy nodes and not to ejected nodes. When the proportion of healthy nodes in the service is less than or equal to the panic threshold, health checks are effectively disabled. Requests are sent to all nodes, including ejected nodes.
	//
	// example:
	//
	// 1
	FailurePercentageMinimumHosts *int32 `json:"failurePercentageMinimumHosts,omitempty" xml:"failurePercentageMinimumHosts,omitempty"`
	// The failure percentage threshold. When the percentage of failed requests on a node reaches this threshold, the system triggers the ejection mechanism for the node.
	//
	// example:
	//
	// 80
	FailurePercentageThreshold *int32 `json:"failurePercentageThreshold,omitempty" xml:"failurePercentageThreshold,omitempty"`
	// The detection interval.
	//
	// example:
	//
	// 30
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
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
	// The port name.
	//
	// example:
	//
	// catalog
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol.
	//
	// example:
	//
	// TCP|UDP
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

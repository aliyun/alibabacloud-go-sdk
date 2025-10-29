// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iService interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *Service
	GetAddresses() []*string
	SetAgentServiceConfig(v *AgentServiceConfig) *Service
	GetAgentServiceConfig() *AgentServiceConfig
	SetAiServiceConfig(v *AiServiceConfig) *Service
	GetAiServiceConfig() *AiServiceConfig
	SetCreateTimestamp(v int64) *Service
	GetCreateTimestamp() *int64
	SetExpressType(v string) *Service
	GetExpressType() *string
	SetGatewayId(v string) *Service
	GetGatewayId() *string
	SetGroupName(v string) *Service
	GetGroupName() *string
	SetHealthCheck(v *ServiceHealthCheck) *Service
	GetHealthCheck() *ServiceHealthCheck
	SetHealthStatus(v string) *Service
	GetHealthStatus() *string
	SetLabelDetails(v []*LabelDetail) *Service
	GetLabelDetails() []*LabelDetail
	SetName(v string) *Service
	GetName() *string
	SetNamespace(v string) *Service
	GetNamespace() *string
	SetOutlierEndpoints(v []*string) *Service
	GetOutlierEndpoints() []*string
	SetPorts(v []*ServicePorts) *Service
	GetPorts() []*ServicePorts
	SetProtocol(v string) *Service
	GetProtocol() *string
	SetQualifier(v string) *Service
	GetQualifier() *string
	SetResourceGroupId(v string) *Service
	GetResourceGroupId() *string
	SetServiceId(v string) *Service
	GetServiceId() *string
	SetSourceType(v string) *Service
	GetSourceType() *string
	SetUnhealthyEndpoints(v []*string) *Service
	GetUnhealthyEndpoints() []*string
	SetUpdateTimestamp(v int64) *Service
	GetUpdateTimestamp() *int64
}

type Service struct {
	Addresses          []*string           `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	AgentServiceConfig *AgentServiceConfig `json:"agentServiceConfig,omitempty" xml:"agentServiceConfig,omitempty"`
	AiServiceConfig    *AiServiceConfig    `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	CreateTimestamp    *int64              `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// StartExecution
	ExpressType *string `json:"expressType,omitempty" xml:"expressType,omitempty"`
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// publich
	GroupName        *string             `json:"groupName,omitempty" xml:"groupName,omitempty"`
	HealthCheck      *ServiceHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty"`
	HealthStatus     *string             `json:"healthStatus,omitempty" xml:"healthStatus,omitempty"`
	LabelDetails     []*LabelDetail      `json:"labelDetails,omitempty" xml:"labelDetails,omitempty" type:"Repeated"`
	Name             *string             `json:"name,omitempty" xml:"name,omitempty"`
	Namespace        *string             `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OutlierEndpoints []*string           `json:"outlierEndpoints,omitempty" xml:"outlierEndpoints,omitempty" type:"Repeated"`
	Ports            []*ServicePorts     `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// rg-xxx
	ResourceGroupId    *string   `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ServiceId          *string   `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	SourceType         *string   `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UnhealthyEndpoints []*string `json:"unhealthyEndpoints,omitempty" xml:"unhealthyEndpoints,omitempty" type:"Repeated"`
	UpdateTimestamp    *int64    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s Service) String() string {
	return dara.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) GetAddresses() []*string {
	return s.Addresses
}

func (s *Service) GetAgentServiceConfig() *AgentServiceConfig {
	return s.AgentServiceConfig
}

func (s *Service) GetAiServiceConfig() *AiServiceConfig {
	return s.AiServiceConfig
}

func (s *Service) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *Service) GetExpressType() *string {
	return s.ExpressType
}

func (s *Service) GetGatewayId() *string {
	return s.GatewayId
}

func (s *Service) GetGroupName() *string {
	return s.GroupName
}

func (s *Service) GetHealthCheck() *ServiceHealthCheck {
	return s.HealthCheck
}

func (s *Service) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *Service) GetLabelDetails() []*LabelDetail {
	return s.LabelDetails
}

func (s *Service) GetName() *string {
	return s.Name
}

func (s *Service) GetNamespace() *string {
	return s.Namespace
}

func (s *Service) GetOutlierEndpoints() []*string {
	return s.OutlierEndpoints
}

func (s *Service) GetPorts() []*ServicePorts {
	return s.Ports
}

func (s *Service) GetProtocol() *string {
	return s.Protocol
}

func (s *Service) GetQualifier() *string {
	return s.Qualifier
}

func (s *Service) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Service) GetServiceId() *string {
	return s.ServiceId
}

func (s *Service) GetSourceType() *string {
	return s.SourceType
}

func (s *Service) GetUnhealthyEndpoints() []*string {
	return s.UnhealthyEndpoints
}

func (s *Service) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *Service) SetAddresses(v []*string) *Service {
	s.Addresses = v
	return s
}

func (s *Service) SetAgentServiceConfig(v *AgentServiceConfig) *Service {
	s.AgentServiceConfig = v
	return s
}

func (s *Service) SetAiServiceConfig(v *AiServiceConfig) *Service {
	s.AiServiceConfig = v
	return s
}

func (s *Service) SetCreateTimestamp(v int64) *Service {
	s.CreateTimestamp = &v
	return s
}

func (s *Service) SetExpressType(v string) *Service {
	s.ExpressType = &v
	return s
}

func (s *Service) SetGatewayId(v string) *Service {
	s.GatewayId = &v
	return s
}

func (s *Service) SetGroupName(v string) *Service {
	s.GroupName = &v
	return s
}

func (s *Service) SetHealthCheck(v *ServiceHealthCheck) *Service {
	s.HealthCheck = v
	return s
}

func (s *Service) SetHealthStatus(v string) *Service {
	s.HealthStatus = &v
	return s
}

func (s *Service) SetLabelDetails(v []*LabelDetail) *Service {
	s.LabelDetails = v
	return s
}

func (s *Service) SetName(v string) *Service {
	s.Name = &v
	return s
}

func (s *Service) SetNamespace(v string) *Service {
	s.Namespace = &v
	return s
}

func (s *Service) SetOutlierEndpoints(v []*string) *Service {
	s.OutlierEndpoints = v
	return s
}

func (s *Service) SetPorts(v []*ServicePorts) *Service {
	s.Ports = v
	return s
}

func (s *Service) SetProtocol(v string) *Service {
	s.Protocol = &v
	return s
}

func (s *Service) SetQualifier(v string) *Service {
	s.Qualifier = &v
	return s
}

func (s *Service) SetResourceGroupId(v string) *Service {
	s.ResourceGroupId = &v
	return s
}

func (s *Service) SetServiceId(v string) *Service {
	s.ServiceId = &v
	return s
}

func (s *Service) SetSourceType(v string) *Service {
	s.SourceType = &v
	return s
}

func (s *Service) SetUnhealthyEndpoints(v []*string) *Service {
	s.UnhealthyEndpoints = v
	return s
}

func (s *Service) SetUpdateTimestamp(v int64) *Service {
	s.UpdateTimestamp = &v
	return s
}

func (s *Service) Validate() error {
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
	if s.HealthCheck != nil {
		if err := s.HealthCheck.Validate(); err != nil {
			return err
		}
	}
	if s.LabelDetails != nil {
		for _, item := range s.LabelDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type ServicePorts struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Port     *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ServicePorts) String() string {
	return dara.Prettify(s)
}

func (s ServicePorts) GoString() string {
	return s.String()
}

func (s *ServicePorts) GetName() *string {
	return s.Name
}

func (s *ServicePorts) GetPort() *int32 {
	return s.Port
}

func (s *ServicePorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ServicePorts) SetName(v string) *ServicePorts {
	s.Name = &v
	return s
}

func (s *ServicePorts) SetPort(v int32) *ServicePorts {
	s.Port = &v
	return s
}

func (s *ServicePorts) SetProtocol(v string) *ServicePorts {
	s.Protocol = &v
	return s
}

func (s *ServicePorts) Validate() error {
	return dara.Validate(s)
}

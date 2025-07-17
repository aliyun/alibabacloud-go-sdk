// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *CreateServiceRequest
	GetGatewayId() *string
	SetResourceGroupId(v string) *CreateServiceRequest
	GetResourceGroupId() *string
	SetServiceConfigs(v []*CreateServiceRequestServiceConfigs) *CreateServiceRequest
	GetServiceConfigs() []*CreateServiceRequestServiceConfigs
	SetSourceType(v string) *CreateServiceRequest
	GetSourceType() *string
}

type CreateServiceRequest struct {
	// The gateway instance ID.
	//
	// example:
	//
	// gw-cq7l5s5lhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The list of service configurations.
	ServiceConfigs []*CreateServiceRequestServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// The service source. Valid values:
	//
	// 	- MSE_NACOS: a service in an MSE Nacos instance
	//
	// 	- K8S: a service in a Kubernetes (K8s) cluster in Container Service for Kubernetes (ACK)
	//
	// 	- VIP: a fixed IP address
	//
	// 	- DNS: a Domain Name System (DNS) domain name
	//
	// 	- FC3: a service in Function Compute
	//
	// 	- SAE_K8S_SERVICE: a service in a K8s cluster in Serverless App Engine (SAE)
	//
	// Enumerated values:
	//
	// 	- SAE_K8S_SERVICE
	//
	// 	- K8S
	//
	// 	- FC3
	//
	// 	- DNS
	//
	// 	- VIP
	//
	// 	- MSE_NACOS
	//
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateServiceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceRequest) GetServiceConfigs() []*CreateServiceRequestServiceConfigs {
	return s.ServiceConfigs
}

func (s *CreateServiceRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateServiceRequest) SetGatewayId(v string) *CreateServiceRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateServiceRequest) SetResourceGroupId(v string) *CreateServiceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceConfigs(v []*CreateServiceRequestServiceConfigs) *CreateServiceRequest {
	s.ServiceConfigs = v
	return s
}

func (s *CreateServiceRequest) SetSourceType(v string) *CreateServiceRequest {
	s.SourceType = &v
	return s
}

func (s *CreateServiceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateServiceRequestServiceConfigs struct {
	// The list of domain names or fixed addresses.
	Addresses          []*string           `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	AgentServiceConfig *AgentServiceConfig `json:"agentServiceConfig,omitempty" xml:"agentServiceConfig,omitempty"`
	// The AI service configurations.
	AiServiceConfig *AiServiceConfig `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	// The list of DNS service addresses.
	DnsServers []*string `json:"dnsServers,omitempty" xml:"dnsServers,omitempty" type:"Repeated"`
	// The service group name. This parameter is required if sourceType is set to MSE_NACOS.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The service name.
	//
	// example:
	//
	// user-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service namespace. This parameter is required when sourceType is set to K8S or MSE_NACOS.
	//
	// 	- If sourceType is set to K8S, this parameter specifies the namespace where the K8s service resides.
	//
	// 	- If sourceType is set to MSE_NACOS, this parameter specifies a namespace in Nacos.
	//
	// example:
	//
	// PUBLIC
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The function version or alias.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s CreateServiceRequestServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequestServiceConfigs) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceConfigs) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateServiceRequestServiceConfigs) GetAgentServiceConfig() *AgentServiceConfig {
	return s.AgentServiceConfig
}

func (s *CreateServiceRequestServiceConfigs) GetAiServiceConfig() *AiServiceConfig {
	return s.AiServiceConfig
}

func (s *CreateServiceRequestServiceConfigs) GetDnsServers() []*string {
	return s.DnsServers
}

func (s *CreateServiceRequestServiceConfigs) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateServiceRequestServiceConfigs) GetName() *string {
	return s.Name
}

func (s *CreateServiceRequestServiceConfigs) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateServiceRequestServiceConfigs) GetQualifier() *string {
	return s.Qualifier
}

func (s *CreateServiceRequestServiceConfigs) SetAddresses(v []*string) *CreateServiceRequestServiceConfigs {
	s.Addresses = v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetAgentServiceConfig(v *AgentServiceConfig) *CreateServiceRequestServiceConfigs {
	s.AgentServiceConfig = v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetAiServiceConfig(v *AiServiceConfig) *CreateServiceRequestServiceConfigs {
	s.AiServiceConfig = v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetDnsServers(v []*string) *CreateServiceRequestServiceConfigs {
	s.DnsServers = v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetGroupName(v string) *CreateServiceRequestServiceConfigs {
	s.GroupName = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetName(v string) *CreateServiceRequestServiceConfigs {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetNamespace(v string) *CreateServiceRequestServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetQualifier(v string) *CreateServiceRequestServiceConfigs {
	s.Qualifier = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) Validate() error {
	return dara.Validate(s)
}

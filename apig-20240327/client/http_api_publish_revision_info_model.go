// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiPublishRevisionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBackendScene(v string) *HttpApiPublishRevisionInfo
	GetBackendScene() *string
	SetBackendType(v string) *HttpApiPublishRevisionInfo
	GetBackendType() *string
	SetCloudProductConfig(v *HttpApiPublishRevisionInfoCloudProductConfig) *HttpApiPublishRevisionInfo
	GetCloudProductConfig() *HttpApiPublishRevisionInfoCloudProductConfig
	SetCreateTimestamp(v int64) *HttpApiPublishRevisionInfo
	GetCreateTimestamp() *int64
	SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo
	GetCustomDomains() []*HttpApiDomainInfo
	SetDnsConfigs(v []*HttpApiPublishRevisionInfoDnsConfigs) *HttpApiPublishRevisionInfo
	GetDnsConfigs() []*HttpApiPublishRevisionInfoDnsConfigs
	SetEnvironmentInfo(v *HttpApiPublishRevisionInfoEnvironmentInfo) *HttpApiPublishRevisionInfo
	GetEnvironmentInfo() *HttpApiPublishRevisionInfoEnvironmentInfo
	SetIsCurrentVersion(v bool) *HttpApiPublishRevisionInfo
	GetIsCurrentVersion() *bool
	SetOperations(v []*HttpApiOperationInfo) *HttpApiPublishRevisionInfo
	GetOperations() []*HttpApiOperationInfo
	SetRevisionId(v string) *HttpApiPublishRevisionInfo
	GetRevisionId() *string
	SetServiceConfigs(v []*HttpApiPublishRevisionInfoServiceConfigs) *HttpApiPublishRevisionInfo
	GetServiceConfigs() []*HttpApiPublishRevisionInfoServiceConfigs
	SetSubDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo
	GetSubDomains() []*HttpApiDomainInfo
	SetVipConfigs(v []*HttpApiPublishRevisionInfoVipConfigs) *HttpApiPublishRevisionInfo
	GetVipConfigs() []*HttpApiPublishRevisionInfoVipConfigs
}

type HttpApiPublishRevisionInfo struct {
	// The publishing scenario.
	//
	// Valid values:
	//
	// 	- SingleService
	//
	// 	- MultiServiceByRatio
	//
	// 	- MultiServiceByContent
	//
	// 	- MultiServiceByTag
	//
	// 	- Mock
	//
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// The type of the backend service.
	//
	// Valid values:
	//
	// 	- DNS: a DNS domain name
	//
	// 	- Service: an existing service
	//
	// 	- VIP: a fixed IP address
	//
	// 	- CloudProduct: a cloud service
	//
	// example:
	//
	// Service
	BackendType *string `json:"backendType,omitempty" xml:"backendType,omitempty"`
	// The cloud service configurations.
	CloudProductConfig *HttpApiPublishRevisionInfoCloudProductConfig `json:"cloudProductConfig,omitempty" xml:"cloudProductConfig,omitempty" type:"Struct"`
	// The publishing timestamp.
	//
	// example:
	//
	// 1718807057927
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The custom domain names.
	CustomDomains []*HttpApiDomainInfo `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	// The configurations of DNS domain names. For single-service publishing, only one entry is allowed. For other scenarios, multiple entries are allowed.
	DnsConfigs []*HttpApiPublishRevisionInfoDnsConfigs `json:"dnsConfigs,omitempty" xml:"dnsConfigs,omitempty" type:"Repeated"`
	// The environment information.
	//
	// example:
	//
	// env-xxx
	EnvironmentInfo *HttpApiPublishRevisionInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// Specifies whether the current version is used.
	//
	// example:
	//
	// true
	IsCurrentVersion *bool `json:"isCurrentVersion,omitempty" xml:"isCurrentVersion,omitempty"`
	// The operations.
	Operations []*HttpApiOperationInfo `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// The published version.
	//
	// example:
	//
	// apr-xxx
	RevisionId *string `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
	// The configurations of existing services. For single-service publishing, only one entry is allowed. For other scenarios, multiple entries are allowed.
	ServiceConfigs []*HttpApiPublishRevisionInfoServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// The default domain names of the environment.
	//
	// example:
	//
	// env-xxx.com
	SubDomains []*HttpApiDomainInfo `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
	// The configurations of fixed IP addresses. For single-service publishing, only one entry is allowed. For other scenarios, multiple entries are allowed.
	VipConfigs []*HttpApiPublishRevisionInfoVipConfigs `json:"vipConfigs,omitempty" xml:"vipConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPublishRevisionInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfo) GetBackendScene() *string {
	return s.BackendScene
}

func (s *HttpApiPublishRevisionInfo) GetBackendType() *string {
	return s.BackendType
}

func (s *HttpApiPublishRevisionInfo) GetCloudProductConfig() *HttpApiPublishRevisionInfoCloudProductConfig {
	return s.CloudProductConfig
}

func (s *HttpApiPublishRevisionInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *HttpApiPublishRevisionInfo) GetCustomDomains() []*HttpApiDomainInfo {
	return s.CustomDomains
}

func (s *HttpApiPublishRevisionInfo) GetDnsConfigs() []*HttpApiPublishRevisionInfoDnsConfigs {
	return s.DnsConfigs
}

func (s *HttpApiPublishRevisionInfo) GetEnvironmentInfo() *HttpApiPublishRevisionInfoEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *HttpApiPublishRevisionInfo) GetIsCurrentVersion() *bool {
	return s.IsCurrentVersion
}

func (s *HttpApiPublishRevisionInfo) GetOperations() []*HttpApiOperationInfo {
	return s.Operations
}

func (s *HttpApiPublishRevisionInfo) GetRevisionId() *string {
	return s.RevisionId
}

func (s *HttpApiPublishRevisionInfo) GetServiceConfigs() []*HttpApiPublishRevisionInfoServiceConfigs {
	return s.ServiceConfigs
}

func (s *HttpApiPublishRevisionInfo) GetSubDomains() []*HttpApiDomainInfo {
	return s.SubDomains
}

func (s *HttpApiPublishRevisionInfo) GetVipConfigs() []*HttpApiPublishRevisionInfoVipConfigs {
	return s.VipConfigs
}

func (s *HttpApiPublishRevisionInfo) SetBackendScene(v string) *HttpApiPublishRevisionInfo {
	s.BackendScene = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetBackendType(v string) *HttpApiPublishRevisionInfo {
	s.BackendType = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCloudProductConfig(v *HttpApiPublishRevisionInfoCloudProductConfig) *HttpApiPublishRevisionInfo {
	s.CloudProductConfig = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCreateTimestamp(v int64) *HttpApiPublishRevisionInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo {
	s.CustomDomains = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetDnsConfigs(v []*HttpApiPublishRevisionInfoDnsConfigs) *HttpApiPublishRevisionInfo {
	s.DnsConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetEnvironmentInfo(v *HttpApiPublishRevisionInfoEnvironmentInfo) *HttpApiPublishRevisionInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetIsCurrentVersion(v bool) *HttpApiPublishRevisionInfo {
	s.IsCurrentVersion = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetOperations(v []*HttpApiOperationInfo) *HttpApiPublishRevisionInfo {
	s.Operations = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetRevisionId(v string) *HttpApiPublishRevisionInfo {
	s.RevisionId = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetServiceConfigs(v []*HttpApiPublishRevisionInfoServiceConfigs) *HttpApiPublishRevisionInfo {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetSubDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo {
	s.SubDomains = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetVipConfigs(v []*HttpApiPublishRevisionInfoVipConfigs) *HttpApiPublishRevisionInfo {
	s.VipConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) Validate() error {
	if s.CloudProductConfig != nil {
		if err := s.CloudProductConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CustomDomains != nil {
		for _, item := range s.CustomDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DnsConfigs != nil {
		for _, item := range s.DnsConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Operations != nil {
		for _, item := range s.Operations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubDomains != nil {
		for _, item := range s.SubDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VipConfigs != nil {
		for _, item := range s.VipConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoCloudProductConfig struct {
	// The type of the cloud service.
	//
	// example:
	//
	// FC
	CloudProductType *string `json:"cloudProductType,omitempty" xml:"cloudProductType,omitempty"`
	// The ACK configurations.
	ContainerServiceConfigs []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs `json:"containerServiceConfigs,omitempty" xml:"containerServiceConfigs,omitempty" type:"Repeated"`
	// The Function Compute configurations.
	FunctionConfigs []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs `json:"functionConfigs,omitempty" xml:"functionConfigs,omitempty" type:"Repeated"`
	// The MSE Nacos configurations.
	MseNacosConfigs []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs `json:"mseNacosConfigs,omitempty" xml:"mseNacosConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetCloudProductType() *string {
	return s.CloudProductType
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetContainerServiceConfigs() []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	return s.ContainerServiceConfigs
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetFunctionConfigs() []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	return s.FunctionConfigs
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetMseNacosConfigs() []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	return s.MseNacosConfigs
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetCloudProductType(v string) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.CloudProductType = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetContainerServiceConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.ContainerServiceConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetFunctionConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.FunctionConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetMseNacosConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.MseNacosConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) Validate() error {
	if s.ContainerServiceConfigs != nil {
		for _, item := range s.ContainerServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FunctionConfigs != nil {
		for _, item := range s.FunctionConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MseNacosConfigs != nil {
		for _, item := range s.MseNacosConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs struct {
	// The associated service ID.
	//
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// The matching conditions.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The K8s service name.
	//
	// example:
	//
	// demo-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The K8s namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The service port.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetNamespace() *string {
	return s.Namespace
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetWeight() *string {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetNamespace(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetPort(v int32) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetProtocol(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetWeight(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs struct {
	// The associated service ID.
	//
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// The matching conditions.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The function name.
	//
	// example:
	//
	// demo-function
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The function version or alias.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetQualifier() *string {
	return s.Qualifier
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetQualifier(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Qualifier = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs struct {
	// The associated service ID.
	//
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// The service group.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The matching conditions.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The Nacos service name.
	//
	// example:
	//
	// spring-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Nacos namespace.
	//
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetGroupName() *string {
	return s.GroupName
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetNamespace() *string {
	return s.Namespace
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetGroupName(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.GroupName = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetNamespace(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoDnsConfigs struct {
	// The DNS domain names.
	DnsList []*string `json:"dnsList,omitempty" xml:"dnsList,omitempty" type:"Repeated"`
	// The matching condition. This condition is valid only in content-based routing.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The weight. Valid values: [1,100]. This parameter is valid only in proportional routing.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoDnsConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoDnsConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) GetDnsList() []*string {
	return s.DnsList
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetDnsList(v []*string) *HttpApiPublishRevisionInfoDnsConfigs {
	s.DnsList = v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoDnsConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoDnsConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoEnvironmentInfo struct {
	// The environment alias.
	//
	// example:
	//
	// Test
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The instance information.
	GatewayInfo *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// The environment name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiPublishRevisionInfoEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetAlias() *string {
	return s.Alias
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetGatewayInfo() *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	return s.GatewayInfo
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetAlias(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.Alias = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetEnvironmentId(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetGatewayInfo(v *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.GatewayInfo = v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetName(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) Validate() error {
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo struct {
	// The instance ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Instance 1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) SetGatewayId(v string) *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) SetName(v string) *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoServiceConfigs struct {
	// The service ID.
	//
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// The matching conditions.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The service port.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetPort(v int32) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetProtocol(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetVersion(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Version = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPublishRevisionInfoVipConfigs struct {
	// The IP addresses.
	Endpoints []*string `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// The matching condition. This condition is valid only in content-based routing.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The weight. Valid values: [1,100]. This parameter is valid only in proportional routing.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoVipConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoVipConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoVipConfigs) GetEndpoints() []*string {
	return s.Endpoints
}

func (s *HttpApiPublishRevisionInfoVipConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoVipConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetEndpoints(v []*string) *HttpApiPublishRevisionInfoVipConfigs {
	s.Endpoints = v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoVipConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoVipConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

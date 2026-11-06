// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayResponseBody
	GetCode() *string
	SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody
	GetData() *GetGatewayResponseBodyData
	SetMessage(v string) *GetGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayResponseBody
	GetRequestId() *string
}

type GetGatewayResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetGatewayResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F138FFC-6E2B-56C1-9BAB-A67462E339D1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayResponseBody) GetData() *GetGatewayResponseBodyData {
	return s.Data
}

func (s *GetGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayResponseBody) SetCode(v string) *GetGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResponseBody) SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResponseBody) SetMessage(v string) *GetGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayResponseBody) SetRequestId(v string) *GetGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayResponseBodyData struct {
	// The billing method. Valid values:
	//
	// - POSTPAY: Pay-as-you-go.
	//
	// - PREPAY: Subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The source from which the gateway was created. Valid values:
	//
	// - Console: The console.
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// The creation timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// Indicates whether elastic scaling is enabled. A value of true indicates that elastic scaling is enabled. A value of false indicates that elastic scaling is disabled. A value of null or a missing value indicates that the status is unknown. This parameter is for query and display purposes only.
	//
	// example:
	//
	// false
	ElasticEnabled *bool `json:"elasticEnabled,omitempty" xml:"elasticEnabled,omitempty"`
	// The list of environments associated with the gateway.
	Environments []*GetGatewayResponseBodyDataEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The expiration timestamp for subscription instances. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// The gateway edition. Valid values:
	//
	// - Professional: Professional instance.
	//
	// - Serverless: Serverless instance.
	//
	// example:
	//
	// Serverless
	GatewayEdition *string `json:"gatewayEdition,omitempty" xml:"gatewayEdition,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The AI multi-tenant V2 operation mode. The default value is ENTERPRISE. This parameter is allowed only for AI and MultiTenantServerless.
	//
	// example:
	//
	// STANDARD
	GatewayMode *string `json:"gatewayMode,omitempty" xml:"gatewayMode,omitempty"`
	// The gateway type. Valid values:
	//
	// - API: API gateway.
	//
	// - AI: AI gateway.
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The network ISP type.
	//
	// example:
	//
	// BGP
	Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
	// The list of ingress addresses for the gateway.
	LoadBalancers []*GetGatewayResponseBodyDataLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	// The O&M window.
	MaintenancePeriod *GetGatewayResponseBodyDataMaintenancePeriod `json:"maintenancePeriod,omitempty" xml:"maintenancePeriod,omitempty" type:"Struct"`
	// The gateway name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network access configuration of the gateway, including the access endpoint type and the security group used by the internal endpoint. This parameter is returned only when the gateway is configured with an internal endpoint.
	NetworkAccessConfig *GetGatewayResponseBodyDataNetworkAccessConfig `json:"networkAccessConfig,omitempty" xml:"networkAccessConfig,omitempty" type:"Struct"`
	// The PrivateLink information of the gateway, including the endpoint ID, access domain name, connection status, and zone information.
	PrivateLink *GetGatewayResponseBodyDataPrivateLink `json:"privateLink,omitempty" xml:"privateLink,omitempty" type:"Struct"`
	// The number of nodes in the gateway instance.
	//
	// example:
	//
	// 2
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm2ij6pwxsvua
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The security group of the gateway.
	SecurityGroup *GetGatewayResponseBodyDataSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// The gateway specifications. Valid values:
	//
	// - apigw.small.x1: Small.
	//
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The gateway status. Valid values:
	//
	// - Running: The gateway is running.
	//
	// - Creating: The gateway is being created.
	//
	// - CreateFailed: The gateway failed to be created.
	//
	// - Upgrading: The gateway is being upgraded.
	//
	// - UpgradeFailed: The gateway failed to be upgraded.
	//
	// - Restarting: The gateway is being restarted.
	//
	// - RestartFailed: The gateway failed to be restarted.
	//
	// - Deleting: The gateway is being released.
	//
	// - DeleteFailed: The gateway failed to be released.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of default access domain name information for the gateway. This parameter is not returned if no default domain name information is available.
	SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	// The resource tags.
	Tags []*GetGatewayResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The target version of the gateway. You can upgrade the gateway if this value is different from the current version.
	//
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// The tenant ID returned in STANDARD mode.
	//
	// example:
	//
	// pt-xxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The update timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	// The vSwitch associated with the gateway.
	VSwitch *GetGatewayResponseBodyDataVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The gateway version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The VPC associated with the gateway.
	Vpc *GetGatewayResponseBodyDataVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	// The list of zones associated with the gateway.
	Zones []*GetGatewayResponseBodyDataZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s GetGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetGatewayResponseBodyData) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *GetGatewayResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetGatewayResponseBodyData) GetElasticEnabled() *bool {
	return s.ElasticEnabled
}

func (s *GetGatewayResponseBodyData) GetEnvironments() []*GetGatewayResponseBodyDataEnvironments {
	return s.Environments
}

func (s *GetGatewayResponseBodyData) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *GetGatewayResponseBodyData) GetGatewayEdition() *string {
	return s.GatewayEdition
}

func (s *GetGatewayResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetGatewayResponseBodyData) GetGatewayMode() *string {
	return s.GatewayMode
}

func (s *GetGatewayResponseBodyData) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetGatewayResponseBodyData) GetIsp() *string {
	return s.Isp
}

func (s *GetGatewayResponseBodyData) GetLoadBalancers() []*GetGatewayResponseBodyDataLoadBalancers {
	return s.LoadBalancers
}

func (s *GetGatewayResponseBodyData) GetMaintenancePeriod() *GetGatewayResponseBodyDataMaintenancePeriod {
	return s.MaintenancePeriod
}

func (s *GetGatewayResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyData) GetNetworkAccessConfig() *GetGatewayResponseBodyDataNetworkAccessConfig {
	return s.NetworkAccessConfig
}

func (s *GetGatewayResponseBodyData) GetPrivateLink() *GetGatewayResponseBodyDataPrivateLink {
	return s.PrivateLink
}

func (s *GetGatewayResponseBodyData) GetReplicas() *string {
	return s.Replicas
}

func (s *GetGatewayResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetGatewayResponseBodyData) GetSecurityGroup() *GetGatewayResponseBodyDataSecurityGroup {
	return s.SecurityGroup
}

func (s *GetGatewayResponseBodyData) GetSpec() *string {
	return s.Spec
}

func (s *GetGatewayResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayResponseBodyData) GetSubDomainInfos() []*SubDomainInfo {
	return s.SubDomainInfos
}

func (s *GetGatewayResponseBodyData) GetTags() []*GetGatewayResponseBodyDataTags {
	return s.Tags
}

func (s *GetGatewayResponseBodyData) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *GetGatewayResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGatewayResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetGatewayResponseBodyData) GetVSwitch() *GetGatewayResponseBodyDataVSwitch {
	return s.VSwitch
}

func (s *GetGatewayResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetGatewayResponseBodyData) GetVpc() *GetGatewayResponseBodyDataVpc {
	return s.Vpc
}

func (s *GetGatewayResponseBodyData) GetZones() []*GetGatewayResponseBodyDataZones {
	return s.Zones
}

func (s *GetGatewayResponseBodyData) SetChargeType(v string) *GetGatewayResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCreateFrom(v string) *GetGatewayResponseBodyData {
	s.CreateFrom = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCreateTimestamp(v int64) *GetGatewayResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetElasticEnabled(v bool) *GetGatewayResponseBodyData {
	s.ElasticEnabled = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEnvironments(v []*GetGatewayResponseBodyDataEnvironments) *GetGatewayResponseBodyData {
	s.Environments = v
	return s
}

func (s *GetGatewayResponseBodyData) SetExpireTimestamp(v int64) *GetGatewayResponseBodyData {
	s.ExpireTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayEdition(v string) *GetGatewayResponseBodyData {
	s.GatewayEdition = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayId(v string) *GetGatewayResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayMode(v string) *GetGatewayResponseBodyData {
	s.GatewayMode = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayType(v string) *GetGatewayResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetIsp(v string) *GetGatewayResponseBodyData {
	s.Isp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLoadBalancers(v []*GetGatewayResponseBodyDataLoadBalancers) *GetGatewayResponseBodyData {
	s.LoadBalancers = v
	return s
}

func (s *GetGatewayResponseBodyData) SetMaintenancePeriod(v *GetGatewayResponseBodyDataMaintenancePeriod) *GetGatewayResponseBodyData {
	s.MaintenancePeriod = v
	return s
}

func (s *GetGatewayResponseBodyData) SetName(v string) *GetGatewayResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetNetworkAccessConfig(v *GetGatewayResponseBodyDataNetworkAccessConfig) *GetGatewayResponseBodyData {
	s.NetworkAccessConfig = v
	return s
}

func (s *GetGatewayResponseBodyData) SetPrivateLink(v *GetGatewayResponseBodyDataPrivateLink) *GetGatewayResponseBodyData {
	s.PrivateLink = v
	return s
}

func (s *GetGatewayResponseBodyData) SetReplicas(v string) *GetGatewayResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetResourceGroupId(v string) *GetGatewayResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSecurityGroup(v *GetGatewayResponseBodyDataSecurityGroup) *GetGatewayResponseBodyData {
	s.SecurityGroup = v
	return s
}

func (s *GetGatewayResponseBodyData) SetSpec(v string) *GetGatewayResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetStatus(v string) *GetGatewayResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSubDomainInfos(v []*SubDomainInfo) *GetGatewayResponseBodyData {
	s.SubDomainInfos = v
	return s
}

func (s *GetGatewayResponseBodyData) SetTags(v []*GetGatewayResponseBodyDataTags) *GetGatewayResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetGatewayResponseBodyData) SetTargetVersion(v string) *GetGatewayResponseBodyData {
	s.TargetVersion = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetTenantId(v string) *GetGatewayResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetUpdateTimestamp(v int64) *GetGatewayResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVSwitch(v *GetGatewayResponseBodyDataVSwitch) *GetGatewayResponseBodyData {
	s.VSwitch = v
	return s
}

func (s *GetGatewayResponseBodyData) SetVersion(v string) *GetGatewayResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVpc(v *GetGatewayResponseBodyDataVpc) *GetGatewayResponseBodyData {
	s.Vpc = v
	return s
}

func (s *GetGatewayResponseBodyData) SetZones(v []*GetGatewayResponseBodyDataZones) *GetGatewayResponseBodyData {
	s.Zones = v
	return s
}

func (s *GetGatewayResponseBodyData) Validate() error {
	if s.Environments != nil {
		for _, item := range s.Environments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LoadBalancers != nil {
		for _, item := range s.LoadBalancers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MaintenancePeriod != nil {
		if err := s.MaintenancePeriod.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkAccessConfig != nil {
		if err := s.NetworkAccessConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateLink != nil {
		if err := s.PrivateLink.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroup != nil {
		if err := s.SecurityGroup.Validate(); err != nil {
			return err
		}
	}
	if s.SubDomainInfos != nil {
		for _, item := range s.SubDomainInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitch != nil {
		if err := s.VSwitch.Validate(); err != nil {
			return err
		}
	}
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayResponseBodyDataEnvironments struct {
	// The environment alias.
	//
	// example:
	//
	// Default environment
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cp9uhudlht***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// default-gw-cp9ugg5***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetGatewayResponseBodyDataEnvironments) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataEnvironments) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataEnvironments) GetAlias() *string {
	return s.Alias
}

func (s *GetGatewayResponseBodyDataEnvironments) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetGatewayResponseBodyDataEnvironments) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataEnvironments) SetAlias(v string) *GetGatewayResponseBodyDataEnvironments {
	s.Alias = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) SetEnvironmentId(v string) *GetGatewayResponseBodyDataEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) SetName(v string) *GetGatewayResponseBodyDataEnvironments {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataLoadBalancers struct {
	// The load balancing address.
	//
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The IP version. Valid values:
	//
	// - ipv4: IPv4.
	//
	// - ipv6: IPv6.
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// The load balancing address type. Valid values:
	//
	// - Internet: Internet.
	//
	// - Intranet: Private network.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// Indicates whether the address is the default ingress address for the gateway.
	//
	// example:
	//
	// true
	GatewayDefault *bool `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	// The list of IPv4 addresses.
	Ipv4Addresses []*string `json:"ipv4Addresses,omitempty" xml:"ipv4Addresses,omitempty" type:"Repeated"`
	// The list of IPv6 addresses.
	Ipv6Addresses []*string `json:"ipv6Addresses,omitempty" xml:"ipv6Addresses,omitempty" type:"Repeated"`
	// The load balancing ID.
	//
	// example:
	//
	// nlb-xoh3pghru7c***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The load balancing provisioning pattern for the gateway. Valid values:
	//
	// - Managed: Managed by Cloud-native API Gateway.
	//
	// example:
	//
	// Managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The list of listened ports.
	Ports []*GetGatewayResponseBodyDataLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// The status of the load balancing. Valid values:
	//
	// - Ready: Active.
	//
	// - NotCreate: Not associated with an instance.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The load balancing type. Valid values:
	//
	// - NLB: Network Load Balancer (NLB).
	//
	// - CLB: Classic Load Balancer.
	//
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetGatewayResponseBodyDataLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataLoadBalancers) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetAddress() *string {
	return s.Address
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetAddressType() *string {
	return s.AddressType
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetGatewayDefault() *bool {
	return s.GatewayDefault
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetIpv4Addresses() []*string {
	return s.Ipv4Addresses
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetIpv6Addresses() []*string {
	return s.Ipv6Addresses
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetMode() *string {
	return s.Mode
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetPorts() []*GetGatewayResponseBodyDataLoadBalancersPorts {
	return s.Ports
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetType() *string {
	return s.Type
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddress(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Address = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddressIpVersion(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddressType(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetGatewayDefault(v bool) *GetGatewayResponseBodyDataLoadBalancers {
	s.GatewayDefault = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetIpv4Addresses(v []*string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Ipv4Addresses = v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetIpv6Addresses(v []*string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Ipv6Addresses = v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetLoadBalancerId(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetMode(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Mode = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetPorts(v []*GetGatewayResponseBodyDataLoadBalancersPorts) *GetGatewayResponseBodyDataLoadBalancers {
	s.Ports = v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetStatus(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetType(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Type = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) Validate() error {
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

type GetGatewayResponseBodyDataLoadBalancersPorts struct {
	// The port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. Valid values:
	//
	// - TCP: TCP.
	//
	// - UDP: UDP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetGatewayResponseBodyDataLoadBalancersPorts) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataLoadBalancersPorts) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) GetPort() *int32 {
	return s.Port
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) SetPort(v int32) *GetGatewayResponseBodyDataLoadBalancersPorts {
	s.Port = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) SetProtocol(v string) *GetGatewayResponseBodyDataLoadBalancersPorts {
	s.Protocol = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataMaintenancePeriod struct {
	// The O&M end time in HH:mm format.
	//
	// example:
	//
	// 06:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The O&M start time in HH:mm format.
	//
	// example:
	//
	// 02:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetGatewayResponseBodyDataMaintenancePeriod) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataMaintenancePeriod) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) SetEndTime(v string) *GetGatewayResponseBodyDataMaintenancePeriod {
	s.EndTime = &v
	return s
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) SetStartTime(v string) *GetGatewayResponseBodyDataMaintenancePeriod {
	s.StartTime = &v
	return s
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataNetworkAccessConfig struct {
	// The ID of the security group configured for the internal endpoint. The security group belongs to the VPC selected by the user and is an unmanaged security group.
	//
	// example:
	//
	// sg-bp1example123456789
	EndpointSecurityGroupId *string `json:"endpointSecurityGroupId,omitempty" xml:"endpointSecurityGroupId,omitempty"`
	// The access endpoint type of the gateway. Valid values:
	//
	// - Intranet: Internal endpoint only.
	//
	// - InternetAndIntranet: Internet and internal endpoints.
	//
	// example:
	//
	// InternetAndIntranet
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetGatewayResponseBodyDataNetworkAccessConfig) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataNetworkAccessConfig) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataNetworkAccessConfig) GetEndpointSecurityGroupId() *string {
	return s.EndpointSecurityGroupId
}

func (s *GetGatewayResponseBodyDataNetworkAccessConfig) GetType() *string {
	return s.Type
}

func (s *GetGatewayResponseBodyDataNetworkAccessConfig) SetEndpointSecurityGroupId(v string) *GetGatewayResponseBodyDataNetworkAccessConfig {
	s.EndpointSecurityGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyDataNetworkAccessConfig) SetType(v string) *GetGatewayResponseBodyDataNetworkAccessConfig {
	s.Type = &v
	return s
}

func (s *GetGatewayResponseBodyDataNetworkAccessConfig) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataPrivateLink struct {
	// The connection status between the internal endpoint and the endpoint service. A value of Connected indicates that the endpoint is connected.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"connectionStatus,omitempty" xml:"connectionStatus,omitempty"`
	// The access domain name of the internal endpoint, which can be used to access the gateway from the user VPC.
	//
	// example:
	//
	// ep-bp1example123456789.epsrv-bp1example123456789.cn-hangzhou.privatelink.aliyuncs.com
	EndpointDomain *string `json:"endpointDomain,omitempty" xml:"endpointDomain,omitempty"`
	// The internal endpoint ID.
	//
	// example:
	//
	// ep-bp1example123456789
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	// The ID of the security group configured for the internal endpoint.
	//
	// example:
	//
	// sg-bp1example123456789
	EndpointSecurityGroupId *string `json:"endpointSecurityGroupId,omitempty" xml:"endpointSecurityGroupId,omitempty"`
	// The status of the internal endpoint. A value of Active indicates that the endpoint is available.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"endpointStatus,omitempty" xml:"endpointStatus,omitempty"`
	// The error code of the most recent fault that occurred during the creation or delete procedure of the private network ingress. This parameter is not returned if no error message is available.
	//
	// example:
	//
	// Forbidden.RAM
	LastErrorCode *string `json:"lastErrorCode,omitempty" xml:"lastErrorCode,omitempty"`
	// The description of the most recent fault that occurred during the creation or delete procedure of the private network ingress. You can use this description along with the error code to troubleshoot the issue. This parameter is not returned if no error message is available.
	//
	// example:
	//
	// private ingress CheckProductOpen failed (Forbidden.RAM)
	LastErrorMessage *string `json:"lastErrorMessage,omitempty" xml:"lastErrorMessage,omitempty"`
	// The time when the system last retrieved the status of the private endpoint. The time is in RFC 3339 format. This parameter is not returned if the status has not been retrieved.
	//
	// example:
	//
	// 2026-11-02T10:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
	// Indicates whether the private endpoint may continue to incur charges. A value of true indicates that the endpoint is created and its release is not confirmed. A value of false indicates that no charge prompt is available. Actual charges are subject to the bill.
	//
	// example:
	//
	// true
	ResourceMayBeCharged *bool `json:"resourceMayBeCharged,omitempty" xml:"resourceMayBeCharged,omitempty"`
	// The current processing stage of the private endpoint. A value of Ready indicates that the private endpoint is ready.
	//
	// example:
	//
	// Ready
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
	// The list of zone information for the private endpoint.
	Zones []*GetGatewayResponseBodyDataPrivateLinkZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s GetGatewayResponseBodyDataPrivateLink) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataPrivateLink) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetEndpointDomain() *string {
	return s.EndpointDomain
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetEndpointSecurityGroupId() *string {
	return s.EndpointSecurityGroupId
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetLastErrorCode() *string {
	return s.LastErrorCode
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetLastErrorMessage() *string {
	return s.LastErrorMessage
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetResourceMayBeCharged() *bool {
	return s.ResourceMayBeCharged
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetStage() *string {
	return s.Stage
}

func (s *GetGatewayResponseBodyDataPrivateLink) GetZones() []*GetGatewayResponseBodyDataPrivateLinkZones {
	return s.Zones
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetConnectionStatus(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.ConnectionStatus = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetEndpointDomain(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.EndpointDomain = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetEndpointId(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.EndpointId = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetEndpointSecurityGroupId(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.EndpointSecurityGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetEndpointStatus(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.EndpointStatus = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetLastErrorCode(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.LastErrorCode = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetLastErrorMessage(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.LastErrorMessage = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetObservedAt(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.ObservedAt = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetResourceMayBeCharged(v bool) *GetGatewayResponseBodyDataPrivateLink {
	s.ResourceMayBeCharged = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetStage(v string) *GetGatewayResponseBodyDataPrivateLink {
	s.Stage = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) SetZones(v []*GetGatewayResponseBodyDataPrivateLinkZones) *GetGatewayResponseBodyDataPrivateLink {
	s.Zones = v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLink) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayResponseBodyDataPrivateLinkZones struct {
	// The connection status of the private endpoint in the zone. A value of Connected indicates that the endpoint is connected.
	//
	// example:
	//
	// Connected
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The zone ID of the private endpoint.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetGatewayResponseBodyDataPrivateLinkZones) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataPrivateLinkZones) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataPrivateLinkZones) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayResponseBodyDataPrivateLinkZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetGatewayResponseBodyDataPrivateLinkZones) SetStatus(v string) *GetGatewayResponseBodyDataPrivateLinkZones {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLinkZones) SetZoneId(v string) *GetGatewayResponseBodyDataPrivateLinkZones {
	s.ZoneId = &v
	return s
}

func (s *GetGatewayResponseBodyDataPrivateLinkZones) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataSecurityGroup struct {
	// The name of the security group.
	//
	// example:
	//
	// APIG-sg-gw-cq7ke5ll***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp16tafq9***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s GetGatewayResponseBodyDataSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataSecurityGroup) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataSecurityGroup) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetGatewayResponseBodyDataSecurityGroup) SetName(v string) *GetGatewayResponseBodyDataSecurityGroup {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataSecurityGroup) SetSecurityGroupId(v string) *GetGatewayResponseBodyDataSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyDataSecurityGroup) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataTags struct {
	// The key of the resource tag.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the resource tag.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetGatewayResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetGatewayResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetGatewayResponseBodyDataTags) SetKey(v string) *GetGatewayResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetGatewayResponseBodyDataTags) SetValue(v string) *GetGatewayResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetGatewayResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataVSwitch struct {
	// The name of the vSwitch.
	//
	// example:
	//
	// Hangzhou VPC vSwitch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1c7ggkj***
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s GetGatewayResponseBodyDataVSwitch) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataVSwitch) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataVSwitch) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetGatewayResponseBodyDataVSwitch) SetName(v string) *GetGatewayResponseBodyDataVSwitch {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataVSwitch) SetVSwitchId(v string) *GetGatewayResponseBodyDataVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *GetGatewayResponseBodyDataVSwitch) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataVpc struct {
	// The name of the VPC.
	//
	// example:
	//
	// Hangzhou VPC
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1llj52lvj6xc***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetGatewayResponseBodyDataVpc) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataVpc) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataVpc) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetGatewayResponseBodyDataVpc) SetName(v string) *GetGatewayResponseBodyDataVpc {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataVpc) SetVpcId(v string) *GetGatewayResponseBodyDataVpc {
	s.VpcId = &v
	return s
}

func (s *GetGatewayResponseBodyDataVpc) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataZones struct {
	// The name of the zone.
	//
	// example:
	//
	// Hangzhou Zone E
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The vSwitch.
	VSwitch *GetGatewayResponseBodyDataZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetGatewayResponseBodyDataZones) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataZones) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataZones) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataZones) GetVSwitch() *GetGatewayResponseBodyDataZonesVSwitch {
	return s.VSwitch
}

func (s *GetGatewayResponseBodyDataZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetGatewayResponseBodyDataZones) SetName(v string) *GetGatewayResponseBodyDataZones {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataZones) SetVSwitch(v *GetGatewayResponseBodyDataZonesVSwitch) *GetGatewayResponseBodyDataZones {
	s.VSwitch = v
	return s
}

func (s *GetGatewayResponseBodyDataZones) SetZoneId(v string) *GetGatewayResponseBodyDataZones {
	s.ZoneId = &v
	return s
}

func (s *GetGatewayResponseBodyDataZones) Validate() error {
	if s.VSwitch != nil {
		if err := s.VSwitch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayResponseBodyDataZonesVSwitch struct {
	// The name of the vSwitch.
	//
	// example:
	//
	// Hangzhou VPC vSwitch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1c7ggkj***
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s GetGatewayResponseBodyDataZonesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataZonesVSwitch) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) SetName(v string) *GetGatewayResponseBodyDataZonesVSwitch {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) SetVSwitchId(v string) *GetGatewayResponseBodyDataZonesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) Validate() error {
	return dara.Validate(s)
}

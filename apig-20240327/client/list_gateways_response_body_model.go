// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewaysResponseBody
	GetCode() *string
	SetData(v *ListGatewaysResponseBodyData) *ListGatewaysResponseBody
	GetData() *ListGatewaysResponseBodyData
	SetMessage(v string) *ListGatewaysResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewaysResponseBody
	GetRequestId() *string
}

type ListGatewaysResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The instances.
	Data *ListGatewaysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5B626361-070A-56A7-B127-ADAC8F3655DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewaysResponseBody) GetData() *ListGatewaysResponseBodyData {
	return s.Data
}

func (s *ListGatewaysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewaysResponseBody) SetCode(v string) *ListGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewaysResponseBody) SetData(v *ListGatewaysResponseBodyData) *ListGatewaysResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewaysResponseBody) SetMessage(v string) *ListGatewaysResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewaysResponseBody) SetRequestId(v string) *ListGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewaysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewaysResponseBodyData struct {
	// The instances.
	Items []*ListGatewaysResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalSize *int64 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListGatewaysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyData) GetItems() []*ListGatewaysResponseBodyDataItems {
	return s.Items
}

func (s *ListGatewaysResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewaysResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewaysResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewaysResponseBodyData) SetItems(v []*ListGatewaysResponseBodyDataItems) *ListGatewaysResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewaysResponseBodyData) SetPageNumber(v int32) *ListGatewaysResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysResponseBodyData) SetPageSize(v int32) *ListGatewaysResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewaysResponseBodyData) SetTotalSize(v int64) *ListGatewaysResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewaysResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewaysResponseBodyDataItems struct {
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// 	- PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The creation source of the instance. Valid values:
	//
	// 	- Console
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// The time when the instance was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The time when the instance expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 172086834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// example:
	//
	// Serverless
	GatewayEdition *string `json:"gatewayEdition,omitempty" xml:"gatewayEdition,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gw-cpv54p5***
	GatewayId   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	Legacy      *bool   `json:"legacy,omitempty" xml:"legacy,omitempty"`
	// The ingress addresses of the instance.
	LoadBalancers []*ListGatewaysResponseBodyDataItemsLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	// The instance name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The node quantity of the instance.
	//
	// example:
	//
	// 2
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The security group information about the instance.
	SecurityGroup *ListGatewaysResponseBodyDataItemsSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// The instance specification. Valid values:
	//
	// 	- apigw.small.x1
	//
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The instance state. Valid values:
	//
	// 	- Running: The instance is running.
	//
	// 	- Creating: The instance is being created.
	//
	// 	- CreateFailed: The instance fails to be created.
	//
	// 	- Upgrading: The instance is being upgraded.
	//
	// 	- UpgradeFailed: The instance fails to be upgraded.
	//
	// 	- Restarting: The instance is being restarted.
	//
	// 	- RestartFailed: The instance fails to be restarted.
	//
	// 	- Deleting: The instance is being released.
	//
	// 	- DeleteFailed: The instance failed to be released.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The second-level domain names.
	SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	// The tags.
	Tags []*ListGatewaysResponseBodyDataItemsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The destination version of the instance. If the value is inconsistent with the current version, you can upgrade the instance.
	//
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// The time when the instance was last updated. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	// The vSwitch information.
	VSwitch *ListGatewaysResponseBodyDataItemsVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The instance version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The virtual private cloud (VPC) information of the instance.
	Vpc *ListGatewaysResponseBodyDataItemsVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	// The availability zones of the instance.
	Zones []*ListGatewaysResponseBodyDataItemsZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s ListGatewaysResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItems) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListGatewaysResponseBodyDataItems) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *ListGatewaysResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListGatewaysResponseBodyDataItems) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *ListGatewaysResponseBodyDataItems) GetGatewayEdition() *string {
	return s.GatewayEdition
}

func (s *ListGatewaysResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewaysResponseBodyDataItems) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewaysResponseBodyDataItems) GetLegacy() *bool {
	return s.Legacy
}

func (s *ListGatewaysResponseBodyDataItems) GetLoadBalancers() []*ListGatewaysResponseBodyDataItemsLoadBalancers {
	return s.LoadBalancers
}

func (s *ListGatewaysResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListGatewaysResponseBodyDataItems) GetReplicas() *string {
	return s.Replicas
}

func (s *ListGatewaysResponseBodyDataItems) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGatewaysResponseBodyDataItems) GetSecurityGroup() *ListGatewaysResponseBodyDataItemsSecurityGroup {
	return s.SecurityGroup
}

func (s *ListGatewaysResponseBodyDataItems) GetSpec() *string {
	return s.Spec
}

func (s *ListGatewaysResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListGatewaysResponseBodyDataItems) GetSubDomainInfos() []*SubDomainInfo {
	return s.SubDomainInfos
}

func (s *ListGatewaysResponseBodyDataItems) GetTags() []*ListGatewaysResponseBodyDataItemsTags {
	return s.Tags
}

func (s *ListGatewaysResponseBodyDataItems) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *ListGatewaysResponseBodyDataItems) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListGatewaysResponseBodyDataItems) GetVSwitch() *ListGatewaysResponseBodyDataItemsVSwitch {
	return s.VSwitch
}

func (s *ListGatewaysResponseBodyDataItems) GetVersion() *string {
	return s.Version
}

func (s *ListGatewaysResponseBodyDataItems) GetVpc() *ListGatewaysResponseBodyDataItemsVpc {
	return s.Vpc
}

func (s *ListGatewaysResponseBodyDataItems) GetZones() []*ListGatewaysResponseBodyDataItemsZones {
	return s.Zones
}

func (s *ListGatewaysResponseBodyDataItems) SetChargeType(v string) *ListGatewaysResponseBodyDataItems {
	s.ChargeType = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetCreateFrom(v string) *ListGatewaysResponseBodyDataItems {
	s.CreateFrom = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetCreateTimestamp(v int64) *ListGatewaysResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetExpireTimestamp(v int64) *ListGatewaysResponseBodyDataItems {
	s.ExpireTimestamp = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetGatewayEdition(v string) *ListGatewaysResponseBodyDataItems {
	s.GatewayEdition = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetGatewayId(v string) *ListGatewaysResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetGatewayType(v string) *ListGatewaysResponseBodyDataItems {
	s.GatewayType = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetLegacy(v bool) *ListGatewaysResponseBodyDataItems {
	s.Legacy = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetLoadBalancers(v []*ListGatewaysResponseBodyDataItemsLoadBalancers) *ListGatewaysResponseBodyDataItems {
	s.LoadBalancers = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetName(v string) *ListGatewaysResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetReplicas(v string) *ListGatewaysResponseBodyDataItems {
	s.Replicas = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetResourceGroupId(v string) *ListGatewaysResponseBodyDataItems {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetSecurityGroup(v *ListGatewaysResponseBodyDataItemsSecurityGroup) *ListGatewaysResponseBodyDataItems {
	s.SecurityGroup = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetSpec(v string) *ListGatewaysResponseBodyDataItems {
	s.Spec = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetStatus(v string) *ListGatewaysResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetSubDomainInfos(v []*SubDomainInfo) *ListGatewaysResponseBodyDataItems {
	s.SubDomainInfos = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetTags(v []*ListGatewaysResponseBodyDataItemsTags) *ListGatewaysResponseBodyDataItems {
	s.Tags = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetTargetVersion(v string) *ListGatewaysResponseBodyDataItems {
	s.TargetVersion = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListGatewaysResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetVSwitch(v *ListGatewaysResponseBodyDataItemsVSwitch) *ListGatewaysResponseBodyDataItems {
	s.VSwitch = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetVersion(v string) *ListGatewaysResponseBodyDataItems {
	s.Version = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetVpc(v *ListGatewaysResponseBodyDataItemsVpc) *ListGatewaysResponseBodyDataItems {
	s.Vpc = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetZones(v []*ListGatewaysResponseBodyDataItemsZones) *ListGatewaysResponseBodyDataItems {
	s.Zones = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) Validate() error {
	if s.LoadBalancers != nil {
		for _, item := range s.LoadBalancers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type ListGatewaysResponseBodyDataItemsLoadBalancers struct {
	// The load balancer IP address.
	//
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The IP version of the address. Valid values:
	//
	// 	- ipv4: IPv4
	//
	// 	- ipv6: IPv6
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// The address type. Valid values:
	//
	// 	- Internet
	//
	// 	- Intranet
	//
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// Indicates whether the address is the default ingress address of the instance.
	//
	// example:
	//
	// true
	GatewayDefault *bool     `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	Ipv4Addresses  []*string `json:"ipv4Addresses,omitempty" xml:"ipv4Addresses,omitempty" type:"Repeated"`
	Ipv6Addresses  []*string `json:"ipv6Addresses,omitempty" xml:"ipv6Addresses,omitempty" type:"Repeated"`
	// The load balancer ID.
	//
	// example:
	//
	// nlb-xqwioje1c91r***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The mode in which the load balancer is provided. Valid values:
	//
	// 	- Managed: Cloud-native API Gateway manages and provides the load balancer.
	//
	// example:
	//
	// Managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The list of listened ports.
	Ports []*ListGatewaysResponseBodyDataItemsLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// The load balancer status. Valid values:
	//
	// 	- Ready: The load balancer is available.
	//
	// 	- NotCreate: The load balancer is not associated with the instance.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The load balancer type. Valid values:
	//
	// 	- NLB: Network Load Balancer
	//
	// 	- CLB: Classic Load Balancer
	//
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetAddress() *string {
	return s.Address
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetAddressType() *string {
	return s.AddressType
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetGatewayDefault() *bool {
	return s.GatewayDefault
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetIpv4Addresses() []*string {
	return s.Ipv4Addresses
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetIpv6Addresses() []*string {
	return s.Ipv6Addresses
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetMode() *string {
	return s.Mode
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetPorts() []*ListGatewaysResponseBodyDataItemsLoadBalancersPorts {
	return s.Ports
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetStatus() *string {
	return s.Status
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) GetType() *string {
	return s.Type
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetAddress(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Address = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetAddressIpVersion(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetAddressType(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetGatewayDefault(v bool) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.GatewayDefault = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetIpv4Addresses(v []*string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Ipv4Addresses = v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetIpv6Addresses(v []*string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Ipv6Addresses = v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetLoadBalancerId(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetMode(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Mode = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetPorts(v []*ListGatewaysResponseBodyDataItemsLoadBalancersPorts) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Ports = v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetStatus(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Status = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetType(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Type = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) Validate() error {
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

type ListGatewaysResponseBodyDataItemsLoadBalancersPorts struct {
	// The port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancersPorts) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancersPorts) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) GetPort() *int32 {
	return s.Port
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) SetPort(v int32) *ListGatewaysResponseBodyDataItemsLoadBalancersPorts {
	s.Port = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) SetProtocol(v string) *ListGatewaysResponseBodyDataItemsLoadBalancersPorts {
	s.Protocol = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) Validate() error {
	return dara.Validate(s)
}

type ListGatewaysResponseBodyDataItemsSecurityGroup struct {
	// The security group ID.
	//
	// example:
	//
	// sg-xxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsSecurityGroup) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListGatewaysResponseBodyDataItemsSecurityGroup) SetSecurityGroupId(v string) *ListGatewaysResponseBodyDataItemsSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsSecurityGroup) Validate() error {
	return dara.Validate(s)
}

type ListGatewaysResponseBodyDataItemsTags struct {
	// The tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsTags) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsTags) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsTags) GetKey() *string {
	return s.Key
}

func (s *ListGatewaysResponseBodyDataItemsTags) GetValue() *string {
	return s.Value
}

func (s *ListGatewaysResponseBodyDataItemsTags) SetKey(v string) *ListGatewaysResponseBodyDataItemsTags {
	s.Key = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsTags) SetValue(v string) *ListGatewaysResponseBodyDataItemsTags {
	s.Value = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsTags) Validate() error {
	return dara.Validate(s)
}

type ListGatewaysResponseBodyDataItemsVSwitch struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsVSwitch) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsVSwitch) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGatewaysResponseBodyDataItemsVSwitch) SetVSwitchId(v string) *ListGatewaysResponseBodyDataItemsVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsVSwitch) Validate() error {
	return dara.Validate(s)
}

type ListGatewaysResponseBodyDataItemsVpc struct {
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsVpc) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsVpc) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewaysResponseBodyDataItemsVpc) SetVpcId(v string) *ListGatewaysResponseBodyDataItemsVpc {
	s.VpcId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsVpc) Validate() error {
	return dara.Validate(s)
}

type ListGatewaysResponseBodyDataItemsZones struct {
	// The vSwitch information.
	VSwitch *ListGatewaysResponseBodyDataItemsZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsZones) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsZones) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsZones) GetVSwitch() *ListGatewaysResponseBodyDataItemsZonesVSwitch {
	return s.VSwitch
}

func (s *ListGatewaysResponseBodyDataItemsZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListGatewaysResponseBodyDataItemsZones) SetVSwitch(v *ListGatewaysResponseBodyDataItemsZonesVSwitch) *ListGatewaysResponseBodyDataItemsZones {
	s.VSwitch = v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsZones) SetZoneId(v string) *ListGatewaysResponseBodyDataItemsZones {
	s.ZoneId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsZones) Validate() error {
	if s.VSwitch != nil {
		if err := s.VSwitch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewaysResponseBodyDataItemsZonesVSwitch struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsZonesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsZonesVSwitch) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsZonesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGatewaysResponseBodyDataItemsZonesVSwitch) SetVSwitchId(v string) *ListGatewaysResponseBodyDataItemsZonesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsZonesVSwitch) Validate() error {
	return dara.Validate(s)
}

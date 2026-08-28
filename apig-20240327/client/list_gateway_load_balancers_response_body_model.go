// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewayLoadBalancersResponseBody
	GetCode() *string
	SetData(v *ListGatewayLoadBalancersResponseBodyData) *ListGatewayLoadBalancersResponseBody
	GetData() *ListGatewayLoadBalancersResponseBodyData
	SetMessage(v string) *ListGatewayLoadBalancersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayLoadBalancersResponseBody
	GetRequestId() *string
}

type ListGatewayLoadBalancersResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListGatewayLoadBalancersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C7C7C3EB-00B6-509A-B6A3-5462EE759047
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewayLoadBalancersResponseBody) GetData() *ListGatewayLoadBalancersResponseBodyData {
	return s.Data
}

func (s *ListGatewayLoadBalancersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayLoadBalancersResponseBody) SetCode(v string) *ListGatewayLoadBalancersResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBody) SetData(v *ListGatewayLoadBalancersResponseBodyData) *ListGatewayLoadBalancersResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBody) SetMessage(v string) *ListGatewayLoadBalancersResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBody) SetRequestId(v string) *ListGatewayLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayLoadBalancersResponseBodyData struct {
	Items []*ListGatewayLoadBalancersResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListGatewayLoadBalancersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyData) GetItems() []*ListGatewayLoadBalancersResponseBodyDataItems {
	return s.Items
}

func (s *ListGatewayLoadBalancersResponseBodyData) SetItems(v []*ListGatewayLoadBalancersResponseBodyDataItems) *ListGatewayLoadBalancersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyData) Validate() error {
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

type ListGatewayLoadBalancersResponseBodyDataItems struct {
	// example:
	//
	// true
	DefaultGatewayIngress *bool `json:"defaultGatewayIngress,omitempty" xml:"defaultGatewayIngress,omitempty"`
	// example:
	//
	// false
	EditEnable *bool `json:"editEnable,omitempty" xml:"editEnable,omitempty"`
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// 47.x.x.x
	LoadBalancerAddress *string `json:"loadBalancerAddress,omitempty" xml:"loadBalancerAddress,omitempty"`
	// example:
	//
	// ipv4
	LoadBalancerAddressIpVersion *string `json:"loadBalancerAddressIpVersion,omitempty" xml:"loadBalancerAddressIpVersion,omitempty"`
	// example:
	//
	// Internet
	LoadBalancerAddressType *string `json:"loadBalancerAddressType,omitempty" xml:"loadBalancerAddressType,omitempty"`
	// example:
	//
	// Active
	LoadBalancerAvailableStatus *string `json:"loadBalancerAvailableStatus,omitempty" xml:"loadBalancerAvailableStatus,omitempty"`
	// example:
	//
	// lb-bp1xxxx / nlb-xxxx
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// example:
	//
	// Managed
	LoadBalancerMode *string `json:"loadBalancerMode,omitempty" xml:"loadBalancerMode,omitempty"`
	// example:
	//
	// my-nlb
	LoadBalancerName *string `json:"loadBalancerName,omitempty" xml:"loadBalancerName,omitempty"`
	// example:
	//
	// slb.s2.small
	LoadBalancerSpec *string `json:"loadBalancerSpec,omitempty" xml:"loadBalancerSpec,omitempty"`
	// example:
	//
	// Ready
	LoadBalancerStatus *string `json:"loadBalancerStatus,omitempty" xml:"loadBalancerStatus,omitempty"`
	// example:
	//
	// NLB
	LoadBalancerType         *string                                                                  `json:"loadBalancerType,omitempty" xml:"loadBalancerType,omitempty"`
	LoadBalancerZoneMappings []*ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings `json:"loadBalancerZoneMappings,omitempty" xml:"loadBalancerZoneMappings,omitempty" type:"Repeated"`
	Ports                    []*ListGatewayLoadBalancersResponseBodyDataItemsPorts                    `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	ServiceWeight *int64 `json:"serviceWeight,omitempty" xml:"serviceWeight,omitempty"`
	// example:
	//
	// -
	VServerGroupMetaInfo   *string                                                                `json:"vServerGroupMetaInfo,omitempty" xml:"vServerGroupMetaInfo,omitempty"`
	VirtualServerGroupList []*ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList `json:"virtualServerGroupList,omitempty" xml:"virtualServerGroupList,omitempty" type:"Repeated"`
}

func (s ListGatewayLoadBalancersResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetDefaultGatewayIngress() *bool {
	return s.DefaultGatewayIngress
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetEditEnable() *bool {
	return s.EditEnable
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerAddress() *string {
	return s.LoadBalancerAddress
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerAddressIpVersion() *string {
	return s.LoadBalancerAddressIpVersion
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerAddressType() *string {
	return s.LoadBalancerAddressType
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerAvailableStatus() *string {
	return s.LoadBalancerAvailableStatus
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerMode() *string {
	return s.LoadBalancerMode
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetLoadBalancerZoneMappings() []*ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings {
	return s.LoadBalancerZoneMappings
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetPorts() []*ListGatewayLoadBalancersResponseBodyDataItemsPorts {
	return s.Ports
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetServiceWeight() *int64 {
	return s.ServiceWeight
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetVServerGroupMetaInfo() *string {
	return s.VServerGroupMetaInfo
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) GetVirtualServerGroupList() []*ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList {
	return s.VirtualServerGroupList
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetDefaultGatewayIngress(v bool) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.DefaultGatewayIngress = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetEditEnable(v bool) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.EditEnable = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetGatewayId(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerAddress(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerAddress = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerAddressIpVersion(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerAddressIpVersion = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerAddressType(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerAddressType = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerAvailableStatus(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerAvailableStatus = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerId(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerMode(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerMode = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerName(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerName = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerSpec(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerSpec = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerStatus(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerType(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerType = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetLoadBalancerZoneMappings(v []*ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.LoadBalancerZoneMappings = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetPorts(v []*ListGatewayLoadBalancersResponseBodyDataItemsPorts) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.Ports = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetServiceWeight(v int64) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.ServiceWeight = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetVServerGroupMetaInfo(v string) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.VServerGroupMetaInfo = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) SetVirtualServerGroupList(v []*ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) *ListGatewayLoadBalancersResponseBodyDataItems {
	s.VirtualServerGroupList = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItems) Validate() error {
	if s.LoadBalancerZoneMappings != nil {
		for _, item := range s.LoadBalancerZoneMappings {
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
	if s.VirtualServerGroupList != nil {
		for _, item := range s.VirtualServerGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings struct {
	LoadBalancerAddresses []*ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses `json:"loadBalancerAddresses,omitempty" xml:"loadBalancerAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// vsw-bp1xxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) GetLoadBalancerAddresses() []*ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	return s.LoadBalancerAddresses
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) SetLoadBalancerAddresses(v []*ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) SetStatus(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings {
	s.Status = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) SetVSwitchId(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) SetZoneId(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappings) Validate() error {
	if s.LoadBalancerAddresses != nil {
		for _, item := range s.LoadBalancerAddresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses struct {
	// example:
	//
	// eip-bp1xxxx
	AllocationId *string `json:"allocationId,omitempty" xml:"allocationId,omitempty"`
	// example:
	//
	// eni-bp1xxxx
	EniId              *string   `json:"eniId,omitempty" xml:"eniId,omitempty"`
	Ipv4LocalAddresses []*string `json:"ipv4LocalAddresses,omitempty" xml:"ipv4LocalAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 2408:xxxx
	Ipv6Address        *string   `json:"ipv6Address,omitempty" xml:"ipv6Address,omitempty"`
	Ipv6LocalAddresses []*string `json:"ipv6LocalAddresses,omitempty" xml:"ipv6LocalAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 192.168.x.x
	PrivateIPv4Address *string `json:"privateIPv4Address,omitempty" xml:"privateIPv4Address,omitempty"`
	// example:
	//
	// Healthy
	PrivateIPv4HcStatus *string `json:"privateIPv4HcStatus,omitempty" xml:"privateIPv4HcStatus,omitempty"`
	// example:
	//
	// Healthy
	PrivateIPv6HcStatus *string `json:"privateIPv6HcStatus,omitempty" xml:"privateIPv6HcStatus,omitempty"`
	// example:
	//
	// 47.x.x.x
	PublicIPv4Address *string `json:"publicIPv4Address,omitempty" xml:"publicIPv4Address,omitempty"`
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetAllocationId() *string {
	return s.AllocationId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetEniId() *string {
	return s.EniId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetIpv4LocalAddresses() []*string {
	return s.Ipv4LocalAddresses
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetIpv6LocalAddresses() []*string {
	return s.Ipv6LocalAddresses
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetPrivateIPv4Address() *string {
	return s.PrivateIPv4Address
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetPrivateIPv4HcStatus() *string {
	return s.PrivateIPv4HcStatus
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetPrivateIPv6HcStatus() *string {
	return s.PrivateIPv6HcStatus
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) GetPublicIPv4Address() *string {
	return s.PublicIPv4Address
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetAllocationId(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.AllocationId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetEniId(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetIpv4LocalAddresses(v []*string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.Ipv4LocalAddresses = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetIpv6Address(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.Ipv6Address = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetIpv6LocalAddresses(v []*string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.Ipv6LocalAddresses = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetPrivateIPv4Address(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4Address = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetPrivateIPv4HcStatus(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4HcStatus = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetPrivateIPv6HcStatus(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv6HcStatus = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) SetPublicIPv4Address(v string) *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses {
	s.PublicIPv4Address = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsLoadBalancerZoneMappingsLoadBalancerAddresses) Validate() error {
	return dara.Validate(s)
}

type ListGatewayLoadBalancersResponseBodyDataItemsPorts struct {
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// https
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsPorts) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsPorts) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsPorts) GetPort() *int32 {
	return s.Port
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsPorts) SetPort(v int32) *ListGatewayLoadBalancersResponseBodyDataItemsPorts {
	s.Port = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsPorts) SetProtocol(v string) *ListGatewayLoadBalancersResponseBodyDataItemsPorts {
	s.Protocol = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsPorts) Validate() error {
	return dara.Validate(s)
}

type ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList struct {
	Listeners []*ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners `json:"listeners,omitempty" xml:"listeners,omitempty" type:"Repeated"`
	// example:
	//
	// rsp-xxxx
	VirtualServiceGroupId *string `json:"virtualServiceGroupId,omitempty" xml:"virtualServiceGroupId,omitempty"`
	// example:
	//
	// 80-tcp
	VirtualServiceGroupName *string `json:"virtualServiceGroupName,omitempty" xml:"virtualServiceGroupName,omitempty"`
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) GetListeners() []*ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners {
	return s.Listeners
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) GetVirtualServiceGroupId() *string {
	return s.VirtualServiceGroupId
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) GetVirtualServiceGroupName() *string {
	return s.VirtualServiceGroupName
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) SetListeners(v []*ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList {
	s.Listeners = v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) SetVirtualServiceGroupId(v string) *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList {
	s.VirtualServiceGroupId = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) SetVirtualServiceGroupName(v string) *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList {
	s.VirtualServiceGroupName = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupList) Validate() error {
	if s.Listeners != nil {
		for _, item := range s.Listeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners struct {
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// https
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) GetPort() *int32 {
	return s.Port
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) GetProtocol() *string {
	return s.Protocol
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) SetPort(v int32) *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners {
	s.Port = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) SetProtocol(v string) *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners {
	s.Protocol = &v
	return s
}

func (s *ListGatewayLoadBalancersResponseBodyDataItemsVirtualServerGroupListListeners) Validate() error {
	return dara.Validate(s)
}

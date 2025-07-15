// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6AddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Addresses(v *DescribeIpv6AddressesResponseBodyIpv6Addresses) *DescribeIpv6AddressesResponseBody
	GetIpv6Addresses() *DescribeIpv6AddressesResponseBodyIpv6Addresses
	SetPageNumber(v int32) *DescribeIpv6AddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6AddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIpv6AddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpv6AddressesResponseBody
	GetTotalCount() *int32
}

type DescribeIpv6AddressesResponseBody struct {
	// The details of the IPv6 address.
	Ipv6Addresses *DescribeIpv6AddressesResponseBodyIpv6Addresses `json:"Ipv6Addresses,omitempty" xml:"Ipv6Addresses,omitempty" type:"Struct"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA4486A8-B6AE-469E-AB09-820EF8ECFA2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpv6AddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBody) GetIpv6Addresses() *DescribeIpv6AddressesResponseBodyIpv6Addresses {
	return s.Ipv6Addresses
}

func (s *DescribeIpv6AddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6AddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6AddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpv6AddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpv6AddressesResponseBody) SetIpv6Addresses(v *DescribeIpv6AddressesResponseBodyIpv6Addresses) *DescribeIpv6AddressesResponseBody {
	s.Ipv6Addresses = v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetPageNumber(v int32) *DescribeIpv6AddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetPageSize(v int32) *DescribeIpv6AddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetRequestId(v string) *DescribeIpv6AddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetTotalCount(v int32) *DescribeIpv6AddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesResponseBodyIpv6Addresses struct {
	Ipv6Address []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6Addresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6Addresses) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6Addresses) GetIpv6Address() []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	return s.Ipv6Address
}

func (s *DescribeIpv6AddressesResponseBodyIpv6Addresses) SetIpv6Address(v []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) *DescribeIpv6AddressesResponseBodyIpv6Addresses {
	s.Ipv6Address = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6Addresses) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address struct {
	// The type of IPv6 address. Valid values:
	//
	// 	- IPv6Address (default): indicates a single IPv6 IP.
	//
	// 	- IPv6Prefix: indicates IPv6 CIDR.
	//
	// example:
	//
	// IPv6Address
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The time when the IPv6 address was created.
	//
	// example:
	//
	// 2020-12-20T14:56:09Z
	AllocationTime *string `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	// The ID of the instance associated with the IPv6 address.
	//
	// example:
	//
	// i-2ze72wuqj4y3jl4f****
	AssociatedInstanceId *string `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	// The type of instance associated with the IPv6 address.
	//
	// example:
	//
	// EcsInstance
	AssociatedInstanceType *string `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	// The IPv6 address.
	//
	// example:
	//
	// 2408:XXXX:153:3921:851c:c435:7b12:1c5f
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The description of the IPv6 address.
	//
	// example:
	//
	// test
	Ipv6AddressDescription *string `json:"Ipv6AddressDescription,omitempty" xml:"Ipv6AddressDescription,omitempty"`
	// The ID of the IPv6 address.
	//
	// example:
	//
	// ipv6-2zen5j4axcp5l5qyy****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	// The name of the IPv6 address.
	//
	// example:
	//
	// test
	Ipv6AddressName *string `json:"Ipv6AddressName,omitempty" xml:"Ipv6AddressName,omitempty"`
	// The ID of the IPv6 gateway to which the IPv6 address belongs.
	//
	// example:
	//
	// ipv6gw-2zewg0l66s73b4k2q****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The Internet bandwidth of the IPv6 address.
	Ipv6InternetBandwidth *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth `json:"Ipv6InternetBandwidth,omitempty" xml:"Ipv6InternetBandwidth,omitempty" type:"Struct"`
	// The ISP of the IPv6 address. Valid values:
	//
	// 	- **BGP*	- (default)
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaTelecom**
	//
	// example:
	//
	// BGP
	Ipv6Isp *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	// The type of connection supported by the IPv6 address. Valid values:
	//
	// 	- **Private**
	//
	// 	- **Public**
	//
	// example:
	//
	// Private
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The peak bandwidth of the IPv6 address.
	//
	// example:
	//
	// 5
	RealBandwidth *int32 `json:"RealBandwidth,omitempty" xml:"RealBandwidth,omitempty"`
	// The ID of the resource group to which the IPv6 gateway belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	ServiceManaged *int32 `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The status of the IPv6 address.
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the IPv6 address belongs.
	//
	// example:
	//
	// vsw-25navfgbue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the IPv6 address belongs.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAllocationTime() *string {
	return s.AllocationTime
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAssociatedInstanceId() *string {
	return s.AssociatedInstanceId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAssociatedInstanceType() *string {
	return s.AssociatedInstanceType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6AddressDescription() *string {
	return s.Ipv6AddressDescription
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6AddressName() *string {
	return s.Ipv6AddressName
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6InternetBandwidth() *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	return s.Ipv6InternetBandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetRealBandwidth() *int32 {
	return s.RealBandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetServiceManaged() *int32 {
	return s.ServiceManaged
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetTags() *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags {
	return s.Tags
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAddressType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AddressType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAllocationTime(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AllocationTime = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAssociatedInstanceId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AssociatedInstanceId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAssociatedInstanceType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6Address(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6AddressDescription(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6AddressDescription = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6AddressId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6AddressId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6AddressName(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6AddressName = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6GatewayId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6InternetBandwidth(v *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6InternetBandwidth = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6Isp(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6Isp = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetNetworkType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.NetworkType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetRealBandwidth(v int32) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.RealBandwidth = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetResourceGroupId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetServiceManaged(v int32) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetStatus(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Status = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetTags(v *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Tags = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetVSwitchId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.VSwitchId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetVpcId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.VpcId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth struct {
	// The dedicated Internet bandwidth of the IPv6 address. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The status of the Internet bandwidth of the IPv6 address. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// 	- **SecurityLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// Indicates whether renewal data is included. Valid values:
	//
	// 	- **false**
	//
	// 	- **true*	- **true*	- is returned only when **IncludeReservationData*	- is set to **true*	- and some orders have not taken effect.
	//
	// example:
	//
	// false
	HasReservationData *bool `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// The billing method of the Internet bandwidth of the IPv6 address. Valid values:
	//
	// Only **PostPaid*	- may be returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The billing method of the Internet bandwidth of the IPv6 address. Valid values:
	//
	// 	- **PayByTraffic**
	//
	// 	- **PayByBandwidth**
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The Internet bandwidth ID of the IPv6 address.
	//
	// example:
	//
	// ipv6bw-hp3b35oq1fj50kbv****
	Ipv6InternetBandwidthId *string `json:"Ipv6InternetBandwidthId,omitempty" xml:"Ipv6InternetBandwidthId,omitempty"`
	// The time when the renewal takes effect. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2021-05-23T16:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The maximum bandwidth after the renewal takes effect. Unit: Mbit/s.
	//
	// example:
	//
	// 12
	ReservationBandwidth *int64 `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The metering method that is used after the renewal takes effect.
	//
	// 	- **PayByTraffic**
	//
	// 	- **PayByBandwidth**
	//
	// example:
	//
	// PayByTraffic
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The type of the renewal order. Only **RENEW*	- may be returned, which indicates that the order is placed for service renewal.
	//
	// example:
	//
	// RENEW
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetHasReservationData() *bool {
	return s.HasReservationData
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetIpv6InternetBandwidthId() *string {
	return s.Ipv6InternetBandwidthId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationBandwidth() *int64 {
	return s.ReservationBandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetBandwidth(v int32) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.Bandwidth = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetBusinessStatus(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetHasReservationData(v bool) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.HasReservationData = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetInstanceChargeType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetInternetChargeType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetIpv6InternetBandwidthId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.Ipv6InternetBandwidthId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationActiveTime(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationBandwidth(v int64) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationInternetChargeType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationOrderType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags struct {
	Tag []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) GetTag() []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag {
	return s.Tag
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) SetTag(v []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags {
	s.Tag = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length. It can be an empty string. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify at most 20 tag values at a time.
	//
	// example:
	//
	// yunke
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) SetKey(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) SetValue(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) Validate() error {
	return dara.Validate(s)
}

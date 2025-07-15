// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6AddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *DescribeIpv6AddressesRequest
	GetAddressType() *string
	SetAssociatedInstanceId(v string) *DescribeIpv6AddressesRequest
	GetAssociatedInstanceId() *string
	SetAssociatedInstanceType(v string) *DescribeIpv6AddressesRequest
	GetAssociatedInstanceType() *string
	SetIncludeReservationData(v bool) *DescribeIpv6AddressesRequest
	GetIncludeReservationData() *bool
	SetIpv6Address(v string) *DescribeIpv6AddressesRequest
	GetIpv6Address() *string
	SetIpv6AddressId(v string) *DescribeIpv6AddressesRequest
	GetIpv6AddressId() *string
	SetIpv6InternetBandwidthId(v string) *DescribeIpv6AddressesRequest
	GetIpv6InternetBandwidthId() *string
	SetName(v string) *DescribeIpv6AddressesRequest
	GetName() *string
	SetNetworkType(v string) *DescribeIpv6AddressesRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeIpv6AddressesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIpv6AddressesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIpv6AddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6AddressesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeIpv6AddressesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeIpv6AddressesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeIpv6AddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIpv6AddressesRequest
	GetResourceOwnerId() *int64
	SetServiceManaged(v bool) *DescribeIpv6AddressesRequest
	GetServiceManaged() *bool
	SetTag(v []*DescribeIpv6AddressesRequestTag) *DescribeIpv6AddressesRequest
	GetTag() []*DescribeIpv6AddressesRequestTag
	SetVSwitchId(v string) *DescribeIpv6AddressesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeIpv6AddressesRequest
	GetVpcId() *string
}

type DescribeIpv6AddressesRequest struct {
	// The type of IP address. Valid values:
	//
	// - IPv6Address (default): indicates an IPv6 instance used to query a single IPv6 address.
	//
	// - IPv6Prefix: indicates an IPv6 instance used to query prefix CIDR blocks.
	//
	// example:
	//
	// IPv6Address
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the instance that is assigned the IPv6 address.
	//
	// example:
	//
	// i-2ze72wuqj4y3jl4f****
	AssociatedInstanceId *string `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	// The type of instance associated with the IPv6 address. Valid values:
	//
	// 	- **EcsInstance**: Elastic Compute Service (ECS) instance in a virtual private cloud (VPC)
	//
	// 	- **NetworkInterface**: secondary elastic network interface (ENI)
	//
	// example:
	//
	// EcsInstance
	AssociatedInstanceType *string `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	// Specifies whether to return information about pending orders. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IncludeReservationData *bool `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	// The IPv6 address that you want to query.
	//
	// example:
	//
	// 2408:XXXX:153:3921:851c:c435:7b12:1c5f
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The ID of the IPv6 address that you want to query. You can enter at most 20 IPv6 address IDs in each API request. Separate IPv6 address IDs with commas (,).
	//
	// example:
	//
	// ipv6-2zen5j4axcp5l5qyy****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	// The ID of the Internet bandwidth that you purchased for the IPv6 address.
	//
	// example:
	//
	// ipv6bw-uf6hcyzu65v98v3du****
	Ipv6InternetBandwidthId *string `json:"Ipv6InternetBandwidthId,omitempty" xml:"Ipv6InternetBandwidthId,omitempty"`
	// The name of the IPv6 address that you want to query.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of communication supported by the IPv6 address. Valid values:
	//
	// 	- **Private**
	//
	// 	- **Public**
	//
	// example:
	//
	// Private
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which you want to query IPv6 addresses. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPv6 gateway belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// If you do not specify this parameter, all instances are queried.
	//
	// example:
	//
	// false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The tag list.
	Tag []*DescribeIpv6AddressesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s DescribeIpv6AddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeIpv6AddressesRequest) GetAssociatedInstanceId() *string {
	return s.AssociatedInstanceId
}

func (s *DescribeIpv6AddressesRequest) GetAssociatedInstanceType() *string {
	return s.AssociatedInstanceType
}

func (s *DescribeIpv6AddressesRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeIpv6AddressesRequest) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeIpv6AddressesRequest) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *DescribeIpv6AddressesRequest) GetIpv6InternetBandwidthId() *string {
	return s.Ipv6InternetBandwidthId
}

func (s *DescribeIpv6AddressesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeIpv6AddressesRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeIpv6AddressesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIpv6AddressesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIpv6AddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6AddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6AddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpv6AddressesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpv6AddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIpv6AddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIpv6AddressesRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeIpv6AddressesRequest) GetTag() []*DescribeIpv6AddressesRequestTag {
	return s.Tag
}

func (s *DescribeIpv6AddressesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeIpv6AddressesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpv6AddressesRequest) SetAddressType(v string) *DescribeIpv6AddressesRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetAssociatedInstanceId(v string) *DescribeIpv6AddressesRequest {
	s.AssociatedInstanceId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetAssociatedInstanceType(v string) *DescribeIpv6AddressesRequest {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetIncludeReservationData(v bool) *DescribeIpv6AddressesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetIpv6Address(v string) *DescribeIpv6AddressesRequest {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetIpv6AddressId(v string) *DescribeIpv6AddressesRequest {
	s.Ipv6AddressId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetIpv6InternetBandwidthId(v string) *DescribeIpv6AddressesRequest {
	s.Ipv6InternetBandwidthId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetName(v string) *DescribeIpv6AddressesRequest {
	s.Name = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetNetworkType(v string) *DescribeIpv6AddressesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetOwnerAccount(v string) *DescribeIpv6AddressesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetOwnerId(v int64) *DescribeIpv6AddressesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetPageNumber(v int32) *DescribeIpv6AddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetPageSize(v int32) *DescribeIpv6AddressesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetRegionId(v string) *DescribeIpv6AddressesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetResourceGroupId(v string) *DescribeIpv6AddressesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetResourceOwnerAccount(v string) *DescribeIpv6AddressesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetResourceOwnerId(v int64) *DescribeIpv6AddressesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetServiceManaged(v bool) *DescribeIpv6AddressesRequest {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetTag(v []*DescribeIpv6AddressesRequestTag) *DescribeIpv6AddressesRequest {
	s.Tag = v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetVSwitchId(v string) *DescribeIpv6AddressesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) SetVpcId(v string) *DescribeIpv6AddressesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeIpv6AddressesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesRequestTag struct {
	// The key of tag N. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// The tag value can be up to 128 characters in length. It can be an empty string. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each tag key corresponds to one tag value. You can specify at most 20 tag values at a time.
	//
	// example:
	//
	// yunke
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIpv6AddressesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeIpv6AddressesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeIpv6AddressesRequestTag) SetKey(v string) *DescribeIpv6AddressesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeIpv6AddressesRequestTag) SetValue(v string) *DescribeIpv6AddressesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeIpv6AddressesRequestTag) Validate() error {
	return dara.Validate(s)
}

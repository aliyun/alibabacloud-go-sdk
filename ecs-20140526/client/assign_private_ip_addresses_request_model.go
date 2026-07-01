// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssignPrivateIpAddressesRequest
	GetClientToken() *string
	SetIpv4Prefix(v []*string) *AssignPrivateIpAddressesRequest
	GetIpv4Prefix() []*string
	SetIpv4PrefixCount(v int32) *AssignPrivateIpAddressesRequest
	GetIpv4PrefixCount() *int32
	SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *AssignPrivateIpAddressesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssignPrivateIpAddressesRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v []*string) *AssignPrivateIpAddressesRequest
	GetPrivateIpAddress() []*string
	SetRegionId(v string) *AssignPrivateIpAddressesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssignPrivateIpAddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssignPrivateIpAddressesRequest
	GetResourceOwnerId() *int64
	SetSecondaryPrivateIpAddressCount(v int32) *AssignPrivateIpAddressesRequest
	GetSecondaryPrivateIpAddressCount() *int32
}

type AssignPrivateIpAddressesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// One or more IPv4 prefixes to allocate to the network interface controller (NIC). Valid values of N: 1 to 10.
	//
	// > To configure IPv4 prefixes for the ENI, you must specify either the Ipv4Prefix.N parameter or the Ipv4PrefixCount parameter, but not both.
	Ipv4Prefix []*string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty" type:"Repeated"`
	// The number of randomly generated IPv4 prefixes to allocate to the network interface controller (NIC). Valid values: 1 to 10.
	//
	// > To configure IPv4 prefixes for the ENI, you must specify either the Ipv4Prefix.N parameter or the Ipv4PrefixCount parameter, but not both.
	//
	// example:
	//
	// 1
	Ipv4PrefixCount *int32 `json:"Ipv4PrefixCount,omitempty" xml:"Ipv4PrefixCount,omitempty"`
	// The ID of the network interface controller (NIC).
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp67acfmxazb4p****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// One or more secondary private IP addresses selected from the available IP addresses of the vSwitch to which the network interface controller (NIC) belongs. Valid values of N:
	//
	// - When the ENI is in the Available (`Available`) state: 1 to 32.
	//
	// - When the ENI is in the `InUse` state: limited by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// When you allocate secondary private IP addresses, you cannot specify both PrivateIpAddress.N and SecondaryPrivateIpAddressCount.
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The region ID of the network interface controller (NIC). You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of private IP addresses to be automatically assigned from the available IP addresses of the vSwitch.
	//
	// When you assign secondary private IP addresses, you cannot specify both PrivateIpAddress.N and SecondaryPrivateIpAddressCount.
	//
	// example:
	//
	// 1
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
}

func (s AssignPrivateIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssignPrivateIpAddressesRequest) GetIpv4Prefix() []*string {
	return s.Ipv4Prefix
}

func (s *AssignPrivateIpAddressesRequest) GetIpv4PrefixCount() *int32 {
	return s.Ipv4PrefixCount
}

func (s *AssignPrivateIpAddressesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignPrivateIpAddressesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssignPrivateIpAddressesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssignPrivateIpAddressesRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *AssignPrivateIpAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssignPrivateIpAddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssignPrivateIpAddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssignPrivateIpAddressesRequest) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *AssignPrivateIpAddressesRequest) SetClientToken(v string) *AssignPrivateIpAddressesRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetIpv4Prefix(v []*string) *AssignPrivateIpAddressesRequest {
	s.Ipv4Prefix = v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetIpv4PrefixCount(v int32) *AssignPrivateIpAddressesRequest {
	s.Ipv4PrefixCount = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetOwnerAccount(v string) *AssignPrivateIpAddressesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetOwnerId(v int64) *AssignPrivateIpAddressesRequest {
	s.OwnerId = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetPrivateIpAddress(v []*string) *AssignPrivateIpAddressesRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetRegionId(v string) *AssignPrivateIpAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetResourceOwnerAccount(v string) *AssignPrivateIpAddressesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetResourceOwnerId(v int64) *AssignPrivateIpAddressesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetSecondaryPrivateIpAddressCount(v int32) *AssignPrivateIpAddressesRequest {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}

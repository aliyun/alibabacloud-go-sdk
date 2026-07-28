// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCanAllocateVpcPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetIpVersion() *string
	SetOwnerAccount(v string) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v string) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *CheckCanAllocateVpcPrivateIpAddressRequest
	GetVSwitchId() *string
}

type CheckCanAllocateVpcPrivateIpAddressRequest struct {
	// The type of the private IP address to check. Valid values:
	//
	// - **ipv4**: IPv4. This parameter is optional when the private IP address is of the IPv4 type.
	//
	// - **ipv6**: IPv6. This parameter is required when the private IP address is of the IPv6 type.
	//
	// example:
	//
	// ipv4
	IpVersion    *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private IP address to check. The private IP address must belong to the vSwitch specified by the **VSwitchId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.7
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID of the vSwitch to which the private IP address belongs.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch to which the private IP address belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-m5ew3t46z2drmifnt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CheckCanAllocateVpcPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCanAllocateVpcPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetIpVersion(v string) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.IpVersion = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetOwnerAccount(v string) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetOwnerId(v int64) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetPrivateIpAddress(v string) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetRegionId(v string) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetResourceOwnerAccount(v string) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetResourceOwnerId(v int64) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) SetVSwitchId(v string) *CheckCanAllocateVpcPrivateIpAddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

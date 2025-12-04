// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssignPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnAssignPrivateIpAddressRequest
	GetClientToken() *string
	SetIpName(v string) *UnAssignPrivateIpAddressRequest
	GetIpName() *string
	SetNetworkInterfaceId(v string) *UnAssignPrivateIpAddressRequest
	GetNetworkInterfaceId() *string
	SetPrivateIpAddress(v string) *UnAssignPrivateIpAddressRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *UnAssignPrivateIpAddressRequest
	GetRegionId() *string
	SetSubnetId(v string) *UnAssignPrivateIpAddressRequest
	GetSubnetId() *string
}

type UnAssignPrivateIpAddressRequest struct {
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// 141cccd6-dfbd-11ec-b8e8-0242ac110003
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IP unique identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-xxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 10.209.75.242
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Region
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Subnet
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s UnAssignPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UnAssignPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnAssignPrivateIpAddressRequest) GetIpName() *string {
	return s.IpName
}

func (s *UnAssignPrivateIpAddressRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *UnAssignPrivateIpAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *UnAssignPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnAssignPrivateIpAddressRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *UnAssignPrivateIpAddressRequest) SetClientToken(v string) *UnAssignPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetIpName(v string) *UnAssignPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *UnAssignPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetPrivateIpAddress(v string) *UnAssignPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetRegionId(v string) *UnAssignPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetSubnetId(v string) *UnAssignPrivateIpAddressRequest {
	s.SubnetId = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

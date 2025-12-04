// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignMac(v bool) *AssignPrivateIpAddressRequest
	GetAssignMac() *bool
	SetClientToken(v string) *AssignPrivateIpAddressRequest
	GetClientToken() *string
	SetDescription(v string) *AssignPrivateIpAddressRequest
	GetDescription() *string
	SetNetworkInterfaceId(v string) *AssignPrivateIpAddressRequest
	GetNetworkInterfaceId() *string
	SetPrivateIpAddress(v string) *AssignPrivateIpAddressRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *AssignPrivateIpAddressRequest
	GetRegionId() *string
	SetSkipConfig(v bool) *AssignPrivateIpAddressRequest
	GetSkipConfig() *bool
	SetSubnetId(v string) *AssignPrivateIpAddressRequest
	GetSubnetId() *string
}

type AssignPrivateIpAddressRequest struct {
	// Specifies whether to assign a mac address.
	//
	// example:
	//
	// true
	AssignMac *bool `json:"AssignMac,omitempty" xml:"AssignMac,omitempty"`
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the variable.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the network interface controller.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The secondary private IP address.
	//
	// example:
	//
	// 10.0.6.194
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The default value is false. If you set the value to true, the secondary private IP address application process can be accelerated.
	//
	// >  For more information, submit a ticket.
	//
	// example:
	//
	// false
	SkipConfig *bool `json:"SkipConfig,omitempty" xml:"SkipConfig,omitempty"`
	// It belongs to the Lingjun subnet.
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s AssignPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressRequest) GetAssignMac() *bool {
	return s.AssignMac
}

func (s *AssignPrivateIpAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssignPrivateIpAddressRequest) GetDescription() *string {
	return s.Description
}

func (s *AssignPrivateIpAddressRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignPrivateIpAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *AssignPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssignPrivateIpAddressRequest) GetSkipConfig() *bool {
	return s.SkipConfig
}

func (s *AssignPrivateIpAddressRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *AssignPrivateIpAddressRequest) SetAssignMac(v bool) *AssignPrivateIpAddressRequest {
	s.AssignMac = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetClientToken(v string) *AssignPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetDescription(v string) *AssignPrivateIpAddressRequest {
	s.Description = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetPrivateIpAddress(v string) *AssignPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetRegionId(v string) *AssignPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetSkipConfig(v bool) *AssignPrivateIpAddressRequest {
	s.SkipConfig = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetSubnetId(v string) *AssignPrivateIpAddressRequest {
	s.SubnetId = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

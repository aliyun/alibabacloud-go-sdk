// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignPrivateIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *UnassignPrivateIpAddressesRequest
	GetNetworkInterfaceId() *string
	SetPrivateIpAddress(v []*string) *UnassignPrivateIpAddressesRequest
	GetPrivateIpAddress() []*string
}

type UnassignPrivateIpAddressesRequest struct {
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-f8z57orgmt6d144t****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The secondary private IP addresses to unassign.
	//
	// This parameter is required.
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
}

func (s UnassignPrivateIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassignPrivateIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *UnassignPrivateIpAddressesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *UnassignPrivateIpAddressesRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *UnassignPrivateIpAddressesRequest) SetNetworkInterfaceId(v string) *UnassignPrivateIpAddressesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetPrivateIpAddress(v []*string) *UnassignPrivateIpAddressesRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}

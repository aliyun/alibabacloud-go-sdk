// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesRequest
	GetNetworkInterfaceId() *string
	SetVSwitchId(v string) *AssignPrivateIpAddressesRequest
	GetVSwitchId() *string
}

type AssignPrivateIpAddressesRequest struct {
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-uf6533jbifugr5fo2j1w
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5rllcjb3ol6duzjdnbm1ombn7
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AssignPrivateIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignPrivateIpAddressesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AssignPrivateIpAddressesRequest) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) SetVSwitchId(v string) *AssignPrivateIpAddressesRequest {
	s.VSwitchId = &v
	return s
}

func (s *AssignPrivateIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}

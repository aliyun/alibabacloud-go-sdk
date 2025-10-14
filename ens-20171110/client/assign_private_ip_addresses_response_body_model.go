// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedPrivateIpAddressesSet(v *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) *AssignPrivateIpAddressesResponseBody
	GetAssignedPrivateIpAddressesSet() *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet
	SetRequestId(v string) *AssignPrivateIpAddressesResponseBody
	GetRequestId() *string
}

type AssignPrivateIpAddressesResponseBody struct {
	// Details about the ENI and the secondary private IP addresses that are assigned to the ENI.
	AssignedPrivateIpAddressesSet *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet `json:"AssignedPrivateIpAddressesSet,omitempty" xml:"AssignedPrivateIpAddressesSet,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignPrivateIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponseBody) GetAssignedPrivateIpAddressesSet() *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	return s.AssignedPrivateIpAddressesSet
}

func (s *AssignPrivateIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignPrivateIpAddressesResponseBody) SetAssignedPrivateIpAddressesSet(v *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) *AssignPrivateIpAddressesResponseBody {
	s.AssignedPrivateIpAddressesSet = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBody) SetRequestId(v string) *AssignPrivateIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignPrivateIpAddressesResponseBody) Validate() error {
	if s.AssignedPrivateIpAddressesSet != nil {
		if err := s.AssignedPrivateIpAddressesSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet struct {
	// The ID of the ENI.
	//
	// example:
	//
	// eni-uf620pb4d19ljnu4a64m
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The assigned private IP addresses.
	PrivateIpSet []*string `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GetPrivateIpSet() []*string {
	return s.PrivateIpSet
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) SetPrivateIpSet(v []*string) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	s.PrivateIpSet = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) Validate() error {
	return dara.Validate(s)
}

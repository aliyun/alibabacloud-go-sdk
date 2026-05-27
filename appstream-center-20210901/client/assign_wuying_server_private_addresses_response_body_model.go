// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignWuyingServerPrivateAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedPrivateIpAddresses(v []*string) *AssignWuyingServerPrivateAddressesResponseBody
	GetAssignedPrivateIpAddresses() []*string
	SetRequestId(v string) *AssignWuyingServerPrivateAddressesResponseBody
	GetRequestId() *string
}

type AssignWuyingServerPrivateAddressesResponseBody struct {
	AssignedPrivateIpAddresses []*string `json:"AssignedPrivateIpAddresses,omitempty" xml:"AssignedPrivateIpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignWuyingServerPrivateAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignWuyingServerPrivateAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *AssignWuyingServerPrivateAddressesResponseBody) GetAssignedPrivateIpAddresses() []*string {
	return s.AssignedPrivateIpAddresses
}

func (s *AssignWuyingServerPrivateAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignWuyingServerPrivateAddressesResponseBody) SetAssignedPrivateIpAddresses(v []*string) *AssignWuyingServerPrivateAddressesResponseBody {
	s.AssignedPrivateIpAddresses = v
	return s
}

func (s *AssignWuyingServerPrivateAddressesResponseBody) SetRequestId(v string) *AssignWuyingServerPrivateAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignWuyingServerPrivateAddressesResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignWuyingServerPrivateAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateIpAddresses(v []*string) *UnassignWuyingServerPrivateAddressesRequest
	GetPrivateIpAddresses() []*string
	SetWuyingServerId(v string) *UnassignWuyingServerPrivateAddressesRequest
	GetWuyingServerId() *string
}

type UnassignWuyingServerPrivateAddressesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["10.0.0.2","10.0.0.3"]
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// aig-bp1234567890abcde
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s UnassignWuyingServerPrivateAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassignWuyingServerPrivateAddressesRequest) GoString() string {
	return s.String()
}

func (s *UnassignWuyingServerPrivateAddressesRequest) GetPrivateIpAddresses() []*string {
	return s.PrivateIpAddresses
}

func (s *UnassignWuyingServerPrivateAddressesRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *UnassignWuyingServerPrivateAddressesRequest) SetPrivateIpAddresses(v []*string) *UnassignWuyingServerPrivateAddressesRequest {
	s.PrivateIpAddresses = v
	return s
}

func (s *UnassignWuyingServerPrivateAddressesRequest) SetWuyingServerId(v string) *UnassignWuyingServerPrivateAddressesRequest {
	s.WuyingServerId = &v
	return s
}

func (s *UnassignWuyingServerPrivateAddressesRequest) Validate() error {
	return dara.Validate(s)
}

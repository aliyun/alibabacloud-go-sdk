// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *AllocateIpAddressRequest
	GetNetworkInterfaceId() *string
	SetOfficeSiteId(v string) *AllocateIpAddressRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *AllocateIpAddressRequest
	GetRegionId() *string
}

type AllocateIpAddressRequest struct {
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OfficeSiteId       *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateIpAddressRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AllocateIpAddressRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *AllocateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateIpAddressRequest) SetNetworkInterfaceId(v string) *AllocateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AllocateIpAddressRequest) SetOfficeSiteId(v string) *AllocateIpAddressRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *AllocateIpAddressRequest) SetRegionId(v string) *AllocateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

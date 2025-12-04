// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLniPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpName(v string) *GetLniPrivateIpAddressRequest
	GetIpName() *string
	SetNetworkInterfaceId(v string) *GetLniPrivateIpAddressRequest
	GetNetworkInterfaceId() *string
	SetRegionId(v string) *GetLniPrivateIpAddressRequest
	GetRegionId() *string
}

type GetLniPrivateIpAddressRequest struct {
	// IP unique identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-xxxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLniPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressRequest) GetIpName() *string {
	return s.IpName
}

func (s *GetLniPrivateIpAddressRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *GetLniPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLniPrivateIpAddressRequest) SetIpName(v string) *GetLniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *GetLniPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *GetLniPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetLniPrivateIpAddressRequest) SetRegionId(v string) *GetLniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *GetLniPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

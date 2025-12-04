// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *GetNetworkInterfaceRequest
	GetNetworkInterfaceId() *string
	SetRegionId(v string) *GetNetworkInterfaceRequest
	GetRegionId() *string
	SetSubnetId(v string) *GetNetworkInterfaceRequest
	GetSubnetId() *string
}

type GetNetworkInterfaceRequest struct {
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
	// Subnet of Lingjun
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s GetNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *GetNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNetworkInterfaceRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *GetNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *GetNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetNetworkInterfaceRequest) SetRegionId(v string) *GetNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *GetNetworkInterfaceRequest) SetSubnetId(v string) *GetNetworkInterfaceRequest {
	s.SubnetId = &v
	return s
}

func (s *GetNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLeniPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticNetworkInterfaceId(v string) *GetLeniPrivateIpAddressRequest
	GetElasticNetworkInterfaceId() *string
	SetIpName(v string) *GetLeniPrivateIpAddressRequest
	GetIpName() *string
	SetRegionId(v string) *GetLeniPrivateIpAddressRequest
	GetRegionId() *string
}

type GetLeniPrivateIpAddressRequest struct {
	// Lingjun Elastic Network Interface ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLeniPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *GetLeniPrivateIpAddressRequest) GetIpName() *string {
	return s.IpName
}

func (s *GetLeniPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *GetLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetLeniPrivateIpAddressRequest) SetIpName(v string) *GetLeniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *GetLeniPrivateIpAddressRequest) SetRegionId(v string) *GetLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *GetLeniPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

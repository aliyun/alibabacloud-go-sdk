// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignLeniPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnassignLeniPrivateIpAddressRequest
	GetClientToken() *string
	SetElasticNetworkInterfaceId(v string) *UnassignLeniPrivateIpAddressRequest
	GetElasticNetworkInterfaceId() *string
	SetIpName(v string) *UnassignLeniPrivateIpAddressRequest
	GetIpName() *string
	SetRegionId(v string) *UnassignLeniPrivateIpAddressRequest
	GetRegionId() *string
}

type UnassignLeniPrivateIpAddressRequest struct {
	// The idempotent identifier.
	//
	// example:
	//
	// 967e77a2-b61d-11ec-a147-0242c0a80504
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s UnassignLeniPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassignLeniPrivateIpAddressRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *UnassignLeniPrivateIpAddressRequest) GetIpName() *string {
	return s.IpName
}

func (s *UnassignLeniPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassignLeniPrivateIpAddressRequest) SetClientToken(v string) *UnassignLeniPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *UnassignLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) SetIpName(v string) *UnassignLeniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) SetRegionId(v string) *UnassignLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

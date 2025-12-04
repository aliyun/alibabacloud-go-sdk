// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLeniPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLeniPrivateIpAddressRequest
	GetDescription() *string
	SetElasticNetworkInterfaceId(v string) *UpdateLeniPrivateIpAddressRequest
	GetElasticNetworkInterfaceId() *string
	SetIpName(v string) *UpdateLeniPrivateIpAddressRequest
	GetIpName() *string
	SetRegionId(v string) *UpdateLeniPrivateIpAddressRequest
	GetRegionId() *string
}

type UpdateLeniPrivateIpAddressRequest struct {
	// The description of the ECS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s UpdateLeniPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLeniPrivateIpAddressRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *UpdateLeniPrivateIpAddressRequest) GetIpName() *string {
	return s.IpName
}

func (s *UpdateLeniPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLeniPrivateIpAddressRequest) SetDescription(v string) *UpdateLeniPrivateIpAddressRequest {
	s.Description = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *UpdateLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) SetIpName(v string) *UpdateLeniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) SetRegionId(v string) *UpdateLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

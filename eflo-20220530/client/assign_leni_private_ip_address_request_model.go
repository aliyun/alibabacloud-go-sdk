// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignLeniPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssignLeniPrivateIpAddressRequest
	GetClientToken() *string
	SetDescription(v string) *AssignLeniPrivateIpAddressRequest
	GetDescription() *string
	SetElasticNetworkInterfaceId(v string) *AssignLeniPrivateIpAddressRequest
	GetElasticNetworkInterfaceId() *string
	SetPrivateIpAddress(v string) *AssignLeniPrivateIpAddressRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *AssignLeniPrivateIpAddressRequest
	GetRegionId() *string
}

type AssignLeniPrivateIpAddressRequest struct {
	// The idempotent identifier.
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the response code.
	//
	// example:
	//
	// wuhuaiyu
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private network IP (automatically assigned by default).
	//
	// example:
	//
	// 10.0.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssignLeniPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssignLeniPrivateIpAddressRequest) GetDescription() *string {
	return s.Description
}

func (s *AssignLeniPrivateIpAddressRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *AssignLeniPrivateIpAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *AssignLeniPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssignLeniPrivateIpAddressRequest) SetClientToken(v string) *AssignLeniPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetDescription(v string) *AssignLeniPrivateIpAddressRequest {
	s.Description = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *AssignLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetPrivateIpAddress(v string) *AssignLeniPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetRegionId(v string) *AssignLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}

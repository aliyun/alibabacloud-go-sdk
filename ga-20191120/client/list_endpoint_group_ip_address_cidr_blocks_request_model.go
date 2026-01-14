// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointGroupIpAddressCidrBlocksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListEndpointGroupIpAddressCidrBlocksRequest
	GetAcceleratorId() *string
	SetEndpointGroupRegion(v string) *ListEndpointGroupIpAddressCidrBlocksRequest
	GetEndpointGroupRegion() *string
	SetRegionId(v string) *ListEndpointGroupIpAddressCidrBlocksRequest
	GetRegionId() *string
}

type ListEndpointGroupIpAddressCidrBlocksRequest struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4q****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The region ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEndpointGroupIpAddressCidrBlocksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupIpAddressCidrBlocksRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) SetAcceleratorId(v string) *ListEndpointGroupIpAddressCidrBlocksRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) SetEndpointGroupRegion(v string) *ListEndpointGroupIpAddressCidrBlocksRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) SetRegionId(v string) *ListEndpointGroupIpAddressCidrBlocksRequest {
	s.RegionId = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksRequest) Validate() error {
	return dara.Validate(s)
}

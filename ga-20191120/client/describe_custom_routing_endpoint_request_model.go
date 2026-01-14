// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroup(v string) *DescribeCustomRoutingEndpointRequest
	GetEndpointGroup() *string
	SetEndpointId(v string) *DescribeCustomRoutingEndpointRequest
	GetEndpointId() *string
	SetRegionId(v string) *DescribeCustomRoutingEndpointRequest
	GetRegionId() *string
}

type DescribeCustomRoutingEndpointRequest struct {
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp1j184jhb9i9ubwf****
	EndpointGroup *string `json:"EndpointGroup,omitempty" xml:"EndpointGroup,omitempty"`
	// The ID of the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp1dmlohjjz4kqaun****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointRequest) GetEndpointGroup() *string {
	return s.EndpointGroup
}

func (s *DescribeCustomRoutingEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeCustomRoutingEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomRoutingEndpointRequest) SetEndpointGroup(v string) *DescribeCustomRoutingEndpointRequest {
	s.EndpointGroup = &v
	return s
}

func (s *DescribeCustomRoutingEndpointRequest) SetEndpointId(v string) *DescribeCustomRoutingEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointRequest) SetRegionId(v string) *DescribeCustomRoutingEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointRequest) Validate() error {
	return dara.Validate(s)
}

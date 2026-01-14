// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreOfEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeLogStoreOfEndpointGroupRequest
	GetAcceleratorId() *string
	SetEndpointGroupId(v string) *DescribeLogStoreOfEndpointGroupRequest
	GetEndpointGroupId() *string
	SetListenerId(v string) *DescribeLogStoreOfEndpointGroupRequest
	GetListenerId() *string
	SetRegionId(v string) *DescribeLogStoreOfEndpointGroupRequest
	GetRegionId() *string
}

type DescribeLogStoreOfEndpointGroupRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-xxxxxxxxxxxxx
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-xxxxxxxxxxxxxxx
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-xxxxxxxxxxxxxxx
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogStoreOfEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreOfEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreOfEndpointGroupRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeLogStoreOfEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeLogStoreOfEndpointGroupRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeLogStoreOfEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogStoreOfEndpointGroupRequest) SetAcceleratorId(v string) *DescribeLogStoreOfEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupRequest) SetEndpointGroupId(v string) *DescribeLogStoreOfEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupRequest) SetListenerId(v string) *DescribeLogStoreOfEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupRequest) SetRegionId(v string) *DescribeLogStoreOfEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}

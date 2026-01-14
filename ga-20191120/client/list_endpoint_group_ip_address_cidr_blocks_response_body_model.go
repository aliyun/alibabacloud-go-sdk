// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointGroupIpAddressCidrBlocksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupRegion(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody
	GetEndpointGroupRegion() *string
	SetIpAddressCidrBlocks(v []*string) *ListEndpointGroupIpAddressCidrBlocksResponseBody
	GetIpAddressCidrBlocks() []*string
	SetRequestId(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody
	GetResourceGroupId() *string
	SetState(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody
	GetState() *string
}

type ListEndpointGroupIpAddressCidrBlocksResponseBody struct {
	// The region ID of the endpoint group.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The CIDR blocks.
	IpAddressCidrBlocks []*string `json:"IpAddressCidrBlocks,omitempty" xml:"IpAddressCidrBlocks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4B6DBBB0-2D01-4C6A-A384-4129266E6B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the endpoint group belongs.
	//
	// example:
	//
	// rg-aekztkx4zwc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the endpoint group.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListEndpointGroupIpAddressCidrBlocksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupIpAddressCidrBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) GetIpAddressCidrBlocks() []*string {
	return s.IpAddressCidrBlocks
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) GetState() *string {
	return s.State
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) SetEndpointGroupRegion(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) SetIpAddressCidrBlocks(v []*string) *ListEndpointGroupIpAddressCidrBlocksResponseBody {
	s.IpAddressCidrBlocks = v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) SetRequestId(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) SetResourceGroupId(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) SetState(v string) *ListEndpointGroupIpAddressCidrBlocksResponseBody {
	s.State = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponseBody) Validate() error {
	return dara.Validate(s)
}

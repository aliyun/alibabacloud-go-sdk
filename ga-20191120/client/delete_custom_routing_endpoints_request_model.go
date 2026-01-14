// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCustomRoutingEndpointsRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *DeleteCustomRoutingEndpointsRequest
	GetEndpointGroupId() *string
	SetEndpointIds(v []*string) *DeleteCustomRoutingEndpointsRequest
	GetEndpointIds() []*string
	SetRegionId(v string) *DeleteCustomRoutingEndpointsRequest
	GetRegionId() *string
}

type DeleteCustomRoutingEndpointsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the endpoint group to which the endpoint that you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1bpn0kn908w4nb****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The IDs of endpoints to be deleted.
	//
	// If you do not set this parameter, all the endpoints in the specified endpoint group are deleted.
	//
	// You can specify at most 10 endpoint IDs.
	//
	// This parameter is required.
	EndpointIds []*string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty" type:"Repeated"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomRoutingEndpointsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DeleteCustomRoutingEndpointsRequest) GetEndpointIds() []*string {
	return s.EndpointIds
}

func (s *DeleteCustomRoutingEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomRoutingEndpointsRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *DeleteCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) SetEndpointIds(v []*string) *DeleteCustomRoutingEndpointsRequest {
	s.EndpointIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) Validate() error {
	return dara.Validate(s)
}

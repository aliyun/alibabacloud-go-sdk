// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteCustomRoutingEndpointGroupsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteCustomRoutingEndpointGroupsRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteCustomRoutingEndpointGroupsRequest
	GetDryRun() *bool
	SetEndpointGroupIds(v []*string) *DeleteCustomRoutingEndpointGroupsRequest
	GetEndpointGroupIds() []*string
	SetRegionId(v string) *DeleteCustomRoutingEndpointGroupsRequest
	GetRegionId() *string
}

type DeleteCustomRoutingEndpointGroupsRequest struct {
	// The ID of the GA instance that you want to query.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true:*	- performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false:*	- performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of the endpoint groups that you want to delete.
	//
	// You can specify up to 10 endpoint group IDs.
	//
	// This parameter is required.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetAcceleratorId(v string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetDryRun(v bool) *DeleteCustomRoutingEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetEndpointGroupIds(v []*string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) Validate() error {
	return dara.Validate(s)
}

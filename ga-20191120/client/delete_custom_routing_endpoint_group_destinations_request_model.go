// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointGroupDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest
	GetClientToken() *string
	SetDestinationIds(v []*string) *DeleteCustomRoutingEndpointGroupDestinationsRequest
	GetDestinationIds() []*string
	SetDryRun(v bool) *DeleteCustomRoutingEndpointGroupDestinationsRequest
	GetDryRun() *bool
	SetEndpointGroupId(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest
	GetRegionId() *string
}

type DeleteCustomRoutingEndpointGroupDestinationsRequest struct {
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
	// The IDs of the endpoint group mappings.
	DestinationIds []*string `json:"DestinationIds,omitempty" xml:"DestinationIds,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint group ID.
	//
	// **
	//
	// ****
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) GetDestinationIds() []*string {
	return s.DestinationIds
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetDestinationIds(v []*string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetDryRun(v bool) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) Validate() error {
	return dara.Validate(s)
}

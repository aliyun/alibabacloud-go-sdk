// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointTrafficPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest
	GetClientToken() *string
	SetEndpointId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest
	GetEndpointId() *string
	SetPolicyIds(v []*string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest
	GetPolicyIds() []*string
	SetRegionId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest
	GetRegionId() *string
}

type DeleteCustomRoutingEndpointTrafficPoliciesRequest struct {
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
	// The ID of the endpoint for which you want to delete traffic destinations.
	//
	// >  This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-2zewuzypq5e6r3pfh****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The IDs of the traffic destinations.
	//
	// You can specify the IDs of up to 9,000 traffic destinations.
	//
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetPolicyIds(v []*string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) Validate() error {
	return dara.Validate(s)
}

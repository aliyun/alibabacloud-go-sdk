// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetBasicEndpointGroupRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *GetBasicEndpointGroupRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *GetBasicEndpointGroupRequest
	GetRegionId() *string
}

type GetBasicEndpointGroupRequest struct {
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
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where the basic GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetBasicEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *GetBasicEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicEndpointGroupRequest) SetClientToken(v string) *GetBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicEndpointGroupRequest) SetEndpointGroupId(v string) *GetBasicEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *GetBasicEndpointGroupRequest) SetRegionId(v string) *GetBasicEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}

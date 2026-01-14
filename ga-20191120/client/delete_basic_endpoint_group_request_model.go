// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteBasicEndpointGroupRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *DeleteBasicEndpointGroupRequest
	GetEndpointGroupId() *string
}

type DeleteBasicEndpointGroupRequest struct {
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
	// The ID of the endpoint group that is associated with the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s DeleteBasicEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBasicEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DeleteBasicEndpointGroupRequest) SetClientToken(v string) *DeleteBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicEndpointGroupRequest) SetEndpointGroupId(v string) *DeleteBasicEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteBasicEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}

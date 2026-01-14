// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteEndpointGroupRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteEndpointGroupRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *DeleteEndpointGroupRequest
	GetEndpointGroupId() *string
}

type DeleteEndpointGroupRequest struct {
	// The GA instance ID.
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
	// The ID of the endpoint group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s DeleteEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DeleteEndpointGroupRequest) SetAcceleratorId(v string) *DeleteEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteEndpointGroupRequest) SetClientToken(v string) *DeleteEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEndpointGroupRequest) SetEndpointGroupId(v string) *DeleteEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}

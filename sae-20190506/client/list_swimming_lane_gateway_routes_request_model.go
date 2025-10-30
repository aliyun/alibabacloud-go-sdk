// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGatewayRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayUniqueId(v string) *ListSwimmingLaneGatewayRoutesRequest
	GetGatewayUniqueId() *string
	SetNamespaceId(v string) *ListSwimmingLaneGatewayRoutesRequest
	GetNamespaceId() *string
}

type ListSwimmingLaneGatewayRoutesRequest struct {
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-277c0727535f4aae917e48de0f******
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListSwimmingLaneGatewayRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGatewayRoutesRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListSwimmingLaneGatewayRoutesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSwimmingLaneGatewayRoutesRequest) SetGatewayUniqueId(v string) *ListSwimmingLaneGatewayRoutesRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesRequest) SetNamespaceId(v string) *ListSwimmingLaneGatewayRoutesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTablesWithVpcGatewayEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateRouteTablesWithVpcGatewayEndpointResponseBody
	GetRequestId() *string
}

type AssociateRouteTablesWithVpcGatewayEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 59BDDA2D-FB52-59F9-9DC5-5EA7D6808B8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateRouteTablesWithVpcGatewayEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTablesWithVpcGatewayEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponseBody) SetRequestId(v string) *AssociateRouteTablesWithVpcGatewayEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateRouteTablesFromVpcGatewayEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateRouteTablesFromVpcGatewayEndpointResponseBody
	GetRequestId() *string
}

type DissociateRouteTablesFromVpcGatewayEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 59BDDA2D-FB52-59F9-9DC5-5EA7D6808B8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateRouteTablesFromVpcGatewayEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateRouteTablesFromVpcGatewayEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponseBody) SetRequestId(v string) *DissociateRouteTablesFromVpcGatewayEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

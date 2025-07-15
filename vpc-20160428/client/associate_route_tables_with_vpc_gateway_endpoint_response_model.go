// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTablesWithVpcGatewayEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateRouteTablesWithVpcGatewayEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateRouteTablesWithVpcGatewayEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AssociateRouteTablesWithVpcGatewayEndpointResponseBody) *AssociateRouteTablesWithVpcGatewayEndpointResponse
	GetBody() *AssociateRouteTablesWithVpcGatewayEndpointResponseBody
}

type AssociateRouteTablesWithVpcGatewayEndpointResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateRouteTablesWithVpcGatewayEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateRouteTablesWithVpcGatewayEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTablesWithVpcGatewayEndpointResponse) GoString() string {
	return s.String()
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) GetBody() *AssociateRouteTablesWithVpcGatewayEndpointResponseBody {
	return s.Body
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) SetHeaders(v map[string]*string) *AssociateRouteTablesWithVpcGatewayEndpointResponse {
	s.Headers = v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) SetStatusCode(v int32) *AssociateRouteTablesWithVpcGatewayEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) SetBody(v *AssociateRouteTablesWithVpcGatewayEndpointResponseBody) *AssociateRouteTablesWithVpcGatewayEndpointResponse {
	s.Body = v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointResponse) Validate() error {
	return dara.Validate(s)
}

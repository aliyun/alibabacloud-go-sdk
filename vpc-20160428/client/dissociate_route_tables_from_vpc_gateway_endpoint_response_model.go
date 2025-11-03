// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateRouteTablesFromVpcGatewayEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateRouteTablesFromVpcGatewayEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateRouteTablesFromVpcGatewayEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DissociateRouteTablesFromVpcGatewayEndpointResponseBody) *DissociateRouteTablesFromVpcGatewayEndpointResponse
	GetBody() *DissociateRouteTablesFromVpcGatewayEndpointResponseBody
}

type DissociateRouteTablesFromVpcGatewayEndpointResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateRouteTablesFromVpcGatewayEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateRouteTablesFromVpcGatewayEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateRouteTablesFromVpcGatewayEndpointResponse) GoString() string {
	return s.String()
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) GetBody() *DissociateRouteTablesFromVpcGatewayEndpointResponseBody {
	return s.Body
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) SetHeaders(v map[string]*string) *DissociateRouteTablesFromVpcGatewayEndpointResponse {
	s.Headers = v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) SetStatusCode(v int32) *DissociateRouteTablesFromVpcGatewayEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) SetBody(v *DissociateRouteTablesFromVpcGatewayEndpointResponseBody) *DissociateRouteTablesFromVpcGatewayEndpointResponse {
	s.Body = v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

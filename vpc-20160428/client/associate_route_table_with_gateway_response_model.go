// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableWithGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateRouteTableWithGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateRouteTableWithGatewayResponse
	GetStatusCode() *int32
	SetBody(v *AssociateRouteTableWithGatewayResponseBody) *AssociateRouteTableWithGatewayResponse
	GetBody() *AssociateRouteTableWithGatewayResponseBody
}

type AssociateRouteTableWithGatewayResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateRouteTableWithGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateRouteTableWithGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableWithGatewayResponse) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableWithGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateRouteTableWithGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateRouteTableWithGatewayResponse) GetBody() *AssociateRouteTableWithGatewayResponseBody {
	return s.Body
}

func (s *AssociateRouteTableWithGatewayResponse) SetHeaders(v map[string]*string) *AssociateRouteTableWithGatewayResponse {
	s.Headers = v
	return s
}

func (s *AssociateRouteTableWithGatewayResponse) SetStatusCode(v int32) *AssociateRouteTableWithGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateRouteTableWithGatewayResponse) SetBody(v *AssociateRouteTableWithGatewayResponseBody) *AssociateRouteTableWithGatewayResponse {
	s.Body = v
	return s
}

func (s *AssociateRouteTableWithGatewayResponse) Validate() error {
	return dara.Validate(s)
}

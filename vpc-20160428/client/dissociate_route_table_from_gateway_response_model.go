// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateRouteTableFromGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateRouteTableFromGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateRouteTableFromGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DissociateRouteTableFromGatewayResponseBody) *DissociateRouteTableFromGatewayResponse
	GetBody() *DissociateRouteTableFromGatewayResponseBody
}

type DissociateRouteTableFromGatewayResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateRouteTableFromGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateRouteTableFromGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateRouteTableFromGatewayResponse) GoString() string {
	return s.String()
}

func (s *DissociateRouteTableFromGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateRouteTableFromGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateRouteTableFromGatewayResponse) GetBody() *DissociateRouteTableFromGatewayResponseBody {
	return s.Body
}

func (s *DissociateRouteTableFromGatewayResponse) SetHeaders(v map[string]*string) *DissociateRouteTableFromGatewayResponse {
	s.Headers = v
	return s
}

func (s *DissociateRouteTableFromGatewayResponse) SetStatusCode(v int32) *DissociateRouteTableFromGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateRouteTableFromGatewayResponse) SetBody(v *DissociateRouteTableFromGatewayResponseBody) *DissociateRouteTableFromGatewayResponse {
	s.Body = v
	return s
}

func (s *DissociateRouteTableFromGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

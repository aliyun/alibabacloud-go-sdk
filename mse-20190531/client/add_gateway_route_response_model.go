// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayRouteResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayRouteResponseBody) *AddGatewayRouteResponse
	GetBody() *AddGatewayRouteResponseBody
}

type AddGatewayRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayRouteResponse) GetBody() *AddGatewayRouteResponseBody {
	return s.Body
}

func (s *AddGatewayRouteResponse) SetHeaders(v map[string]*string) *AddGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayRouteResponse) SetStatusCode(v int32) *AddGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayRouteResponse) SetBody(v *AddGatewayRouteResponseBody) *AddGatewayRouteResponse {
	s.Body = v
	return s
}

func (s *AddGatewayRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

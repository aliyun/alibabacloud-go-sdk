// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyGatewayRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyGatewayRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyGatewayRouteResponse
	GetStatusCode() *int32
	SetBody(v *ApplyGatewayRouteResponseBody) *ApplyGatewayRouteResponse
	GetBody() *ApplyGatewayRouteResponseBody
}

type ApplyGatewayRouteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyGatewayRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *ApplyGatewayRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyGatewayRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyGatewayRouteResponse) GetBody() *ApplyGatewayRouteResponseBody {
	return s.Body
}

func (s *ApplyGatewayRouteResponse) SetHeaders(v map[string]*string) *ApplyGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *ApplyGatewayRouteResponse) SetStatusCode(v int32) *ApplyGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyGatewayRouteResponse) SetBody(v *ApplyGatewayRouteResponseBody) *ApplyGatewayRouteResponse {
	s.Body = v
	return s
}

func (s *ApplyGatewayRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

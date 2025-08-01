// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayRouteResponseBody) *DeleteGatewayRouteResponse
	GetBody() *DeleteGatewayRouteResponseBody
}

type DeleteGatewayRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayRouteResponse) GetBody() *DeleteGatewayRouteResponseBody {
	return s.Body
}

func (s *DeleteGatewayRouteResponse) SetHeaders(v map[string]*string) *DeleteGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayRouteResponse) SetStatusCode(v int32) *DeleteGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayRouteResponse) SetBody(v *DeleteGatewayRouteResponseBody) *DeleteGatewayRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayRouteResponse) Validate() error {
	return dara.Validate(s)
}

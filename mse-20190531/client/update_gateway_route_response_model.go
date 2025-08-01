// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteResponseBody) *UpdateGatewayRouteResponse
	GetBody() *UpdateGatewayRouteResponseBody
}

type UpdateGatewayRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteResponse) GetBody() *UpdateGatewayRouteResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteResponse) SetStatusCode(v int32) *UpdateGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteResponse) SetBody(v *UpdateGatewayRouteResponseBody) *UpdateGatewayRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteResponse) Validate() error {
	return dara.Validate(s)
}

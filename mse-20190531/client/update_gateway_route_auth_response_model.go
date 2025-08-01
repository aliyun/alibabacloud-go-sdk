// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteAuthResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteAuthResponseBody) *UpdateGatewayRouteAuthResponse
	GetBody() *UpdateGatewayRouteAuthResponseBody
}

type UpdateGatewayRouteAuthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteAuthResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteAuthResponse) GetBody() *UpdateGatewayRouteAuthResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteAuthResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteAuthResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteAuthResponse) SetStatusCode(v int32) *UpdateGatewayRouteAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponse) SetBody(v *UpdateGatewayRouteAuthResponseBody) *UpdateGatewayRouteAuthResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteAuthResponse) Validate() error {
	return dara.Validate(s)
}

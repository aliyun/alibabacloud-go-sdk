// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteOnAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayRouteOnAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayRouteOnAuthResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayRouteOnAuthResponseBody) *ListGatewayRouteOnAuthResponse
	GetBody() *ListGatewayRouteOnAuthResponseBody
}

type ListGatewayRouteOnAuthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayRouteOnAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayRouteOnAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteOnAuthResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteOnAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayRouteOnAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayRouteOnAuthResponse) GetBody() *ListGatewayRouteOnAuthResponseBody {
	return s.Body
}

func (s *ListGatewayRouteOnAuthResponse) SetHeaders(v map[string]*string) *ListGatewayRouteOnAuthResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayRouteOnAuthResponse) SetStatusCode(v int32) *ListGatewayRouteOnAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponse) SetBody(v *ListGatewayRouteOnAuthResponseBody) *ListGatewayRouteOnAuthResponse {
	s.Body = v
	return s
}

func (s *ListGatewayRouteOnAuthResponse) Validate() error {
	return dara.Validate(s)
}

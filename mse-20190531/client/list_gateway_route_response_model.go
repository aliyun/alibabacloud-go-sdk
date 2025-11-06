// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayRouteResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayRouteResponseBody) *ListGatewayRouteResponse
	GetBody() *ListGatewayRouteResponseBody
}

type ListGatewayRouteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayRouteResponse) GetBody() *ListGatewayRouteResponseBody {
	return s.Body
}

func (s *ListGatewayRouteResponse) SetHeaders(v map[string]*string) *ListGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayRouteResponse) SetStatusCode(v int32) *ListGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayRouteResponse) SetBody(v *ListGatewayRouteResponseBody) *ListGatewayRouteResponse {
	s.Body = v
	return s
}

func (s *ListGatewayRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

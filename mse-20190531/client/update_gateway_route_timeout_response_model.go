// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTimeoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteTimeoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteTimeoutResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteTimeoutResponseBody) *UpdateGatewayRouteTimeoutResponse
	GetBody() *UpdateGatewayRouteTimeoutResponseBody
}

type UpdateGatewayRouteTimeoutResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteTimeoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteTimeoutResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTimeoutResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTimeoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteTimeoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteTimeoutResponse) GetBody() *UpdateGatewayRouteTimeoutResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteTimeoutResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteTimeoutResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponse) SetStatusCode(v int32) *UpdateGatewayRouteTimeoutResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponse) SetBody(v *UpdateGatewayRouteTimeoutResponseBody) *UpdateGatewayRouteTimeoutResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

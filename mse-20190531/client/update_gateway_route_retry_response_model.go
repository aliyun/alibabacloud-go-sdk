// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteRetryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteRetryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteRetryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteRetryResponseBody) *UpdateGatewayRouteRetryResponse
	GetBody() *UpdateGatewayRouteRetryResponseBody
}

type UpdateGatewayRouteRetryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteRetryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteRetryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRetryResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRetryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteRetryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteRetryResponse) GetBody() *UpdateGatewayRouteRetryResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteRetryResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteRetryResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteRetryResponse) SetStatusCode(v int32) *UpdateGatewayRouteRetryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponse) SetBody(v *UpdateGatewayRouteRetryResponseBody) *UpdateGatewayRouteRetryResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteRetryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

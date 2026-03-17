// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayBgpRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayBgpRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayBgpRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayBgpRouteResponseBody) *UpdateSmartAccessGatewayBgpRouteResponse
	GetBody() *UpdateSmartAccessGatewayBgpRouteResponseBody
}

type UpdateSmartAccessGatewayBgpRouteResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayBgpRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayBgpRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayBgpRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) GetBody() *UpdateSmartAccessGatewayBgpRouteResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayBgpRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayBgpRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) SetBody(v *UpdateSmartAccessGatewayBgpRouteResponseBody) *UpdateSmartAccessGatewayBgpRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

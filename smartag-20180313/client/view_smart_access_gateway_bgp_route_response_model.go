// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayBgpRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayBgpRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayBgpRouteResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayBgpRouteResponseBody) *ViewSmartAccessGatewayBgpRouteResponse
	GetBody() *ViewSmartAccessGatewayBgpRouteResponseBody
}

type ViewSmartAccessGatewayBgpRouteResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayBgpRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayBgpRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayBgpRouteResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) GetBody() *ViewSmartAccessGatewayBgpRouteResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayBgpRouteResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayBgpRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) SetBody(v *ViewSmartAccessGatewayBgpRouteResponseBody) *ViewSmartAccessGatewayBgpRouteResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

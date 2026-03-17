// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayOspfRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayOspfRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayOspfRouteResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayOspfRouteResponseBody) *ViewSmartAccessGatewayOspfRouteResponse
	GetBody() *ViewSmartAccessGatewayOspfRouteResponseBody
}

type ViewSmartAccessGatewayOspfRouteResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayOspfRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayOspfRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayOspfRouteResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) GetBody() *ViewSmartAccessGatewayOspfRouteResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayOspfRouteResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayOspfRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) SetBody(v *ViewSmartAccessGatewayOspfRouteResponseBody) *ViewSmartAccessGatewayOspfRouteResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

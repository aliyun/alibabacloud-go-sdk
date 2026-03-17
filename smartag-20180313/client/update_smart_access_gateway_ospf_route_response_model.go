// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayOspfRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayOspfRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayOspfRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayOspfRouteResponseBody) *UpdateSmartAccessGatewayOspfRouteResponse
	GetBody() *UpdateSmartAccessGatewayOspfRouteResponseBody
}

type UpdateSmartAccessGatewayOspfRouteResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayOspfRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayOspfRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayOspfRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) GetBody() *UpdateSmartAccessGatewayOspfRouteResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayOspfRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayOspfRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) SetBody(v *UpdateSmartAccessGatewayOspfRouteResponseBody) *UpdateSmartAccessGatewayOspfRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

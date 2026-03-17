// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayGlobalRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayGlobalRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) *UpdateSmartAccessGatewayGlobalRouteProtocolResponse
	GetBody() *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody
}

type UpdateSmartAccessGatewayGlobalRouteProtocolResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayGlobalRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayGlobalRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) GetBody() *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayGlobalRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) SetBody(v *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) *UpdateSmartAccessGatewayGlobalRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

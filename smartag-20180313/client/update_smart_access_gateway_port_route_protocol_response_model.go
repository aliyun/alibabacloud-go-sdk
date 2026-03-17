// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayPortRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayPortRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayPortRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) *UpdateSmartAccessGatewayPortRouteProtocolResponse
	GetBody() *UpdateSmartAccessGatewayPortRouteProtocolResponseBody
}

type UpdateSmartAccessGatewayPortRouteProtocolResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayPortRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayPortRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayPortRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) GetBody() *UpdateSmartAccessGatewayPortRouteProtocolResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayPortRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayPortRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) SetBody(v *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) *UpdateSmartAccessGatewayPortRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

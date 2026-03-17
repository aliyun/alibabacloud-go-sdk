// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayGlobalRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayGlobalRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayGlobalRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) *ViewSmartAccessGatewayGlobalRouteProtocolResponse
	GetBody() *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody
}

type ViewSmartAccessGatewayGlobalRouteProtocolResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) GetBody() *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayGlobalRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayGlobalRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) SetBody(v *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) *ViewSmartAccessGatewayGlobalRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

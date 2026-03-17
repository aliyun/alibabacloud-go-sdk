// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayPortRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayPortRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayPortRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayPortRouteProtocolResponseBody) *ViewSmartAccessGatewayPortRouteProtocolResponse
	GetBody() *ViewSmartAccessGatewayPortRouteProtocolResponseBody
}

type ViewSmartAccessGatewayPortRouteProtocolResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayPortRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) GetBody() *ViewSmartAccessGatewayPortRouteProtocolResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayPortRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayPortRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) SetBody(v *ViewSmartAccessGatewayPortRouteProtocolResponseBody) *ViewSmartAccessGatewayPortRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

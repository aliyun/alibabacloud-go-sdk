// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnhancedVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnhancedVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnhancedVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnhancedVpnGatewayResponseBody) *DeleteEnhancedVpnGatewayResponse
	GetBody() *DeleteEnhancedVpnGatewayResponseBody
}

type DeleteEnhancedVpnGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnhancedVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnhancedVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnhancedVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnhancedVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnhancedVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnhancedVpnGatewayResponse) GetBody() *DeleteEnhancedVpnGatewayResponseBody {
	return s.Body
}

func (s *DeleteEnhancedVpnGatewayResponse) SetHeaders(v map[string]*string) *DeleteEnhancedVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnhancedVpnGatewayResponse) SetStatusCode(v int32) *DeleteEnhancedVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayResponse) SetBody(v *DeleteEnhancedVpnGatewayResponseBody) *DeleteEnhancedVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteEnhancedVpnGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

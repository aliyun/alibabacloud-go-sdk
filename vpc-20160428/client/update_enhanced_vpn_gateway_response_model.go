// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnhancedVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnhancedVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnhancedVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnhancedVpnGatewayResponseBody) *UpdateEnhancedVpnGatewayResponse
	GetBody() *UpdateEnhancedVpnGatewayResponseBody
}

type UpdateEnhancedVpnGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnhancedVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnhancedVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnhancedVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnhancedVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnhancedVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnhancedVpnGatewayResponse) GetBody() *UpdateEnhancedVpnGatewayResponseBody {
	return s.Body
}

func (s *UpdateEnhancedVpnGatewayResponse) SetHeaders(v map[string]*string) *UpdateEnhancedVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponse) SetStatusCode(v int32) *UpdateEnhancedVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponse) SetBody(v *UpdateEnhancedVpnGatewayResponseBody) *UpdateEnhancedVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

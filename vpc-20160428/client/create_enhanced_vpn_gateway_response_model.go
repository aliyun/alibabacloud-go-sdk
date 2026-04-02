// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnhancedVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnhancedVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnhancedVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnhancedVpnGatewayResponseBody) *CreateEnhancedVpnGatewayResponse
	GetBody() *CreateEnhancedVpnGatewayResponseBody
}

type CreateEnhancedVpnGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnhancedVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnhancedVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnhancedVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateEnhancedVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnhancedVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnhancedVpnGatewayResponse) GetBody() *CreateEnhancedVpnGatewayResponseBody {
	return s.Body
}

func (s *CreateEnhancedVpnGatewayResponse) SetHeaders(v map[string]*string) *CreateEnhancedVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateEnhancedVpnGatewayResponse) SetStatusCode(v int32) *CreateEnhancedVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnhancedVpnGatewayResponse) SetBody(v *CreateEnhancedVpnGatewayResponseBody) *CreateEnhancedVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateEnhancedVpnGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

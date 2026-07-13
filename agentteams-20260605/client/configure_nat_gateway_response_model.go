// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureNatGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureNatGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureNatGatewayResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureNatGatewayResponseBody) *ConfigureNatGatewayResponse
	GetBody() *ConfigureNatGatewayResponseBody
}

type ConfigureNatGatewayResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureNatGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureNatGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureNatGatewayResponse) GoString() string {
	return s.String()
}

func (s *ConfigureNatGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureNatGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureNatGatewayResponse) GetBody() *ConfigureNatGatewayResponseBody {
	return s.Body
}

func (s *ConfigureNatGatewayResponse) SetHeaders(v map[string]*string) *ConfigureNatGatewayResponse {
	s.Headers = v
	return s
}

func (s *ConfigureNatGatewayResponse) SetStatusCode(v int32) *ConfigureNatGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureNatGatewayResponse) SetBody(v *ConfigureNatGatewayResponseBody) *ConfigureNatGatewayResponse {
	s.Body = v
	return s
}

func (s *ConfigureNatGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

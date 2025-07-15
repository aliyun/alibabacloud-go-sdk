// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpnGatewayResponseBody) *CreateVpnGatewayResponse
	GetBody() *CreateVpnGatewayResponseBody
}

type CreateVpnGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpnGatewayResponse) GetBody() *CreateVpnGatewayResponseBody {
	return s.Body
}

func (s *CreateVpnGatewayResponse) SetHeaders(v map[string]*string) *CreateVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateVpnGatewayResponse) SetStatusCode(v int32) *CreateVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpnGatewayResponse) SetBody(v *CreateVpnGatewayResponseBody) *CreateVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateVpnGatewayResponse) Validate() error {
	return dara.Validate(s)
}

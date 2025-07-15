// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv6GatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpv6GatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpv6GatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpv6GatewayResponseBody) *CreateIpv6GatewayResponse
	GetBody() *CreateIpv6GatewayResponseBody
}

type CreateIpv6GatewayResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpv6GatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpv6GatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6GatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateIpv6GatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpv6GatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpv6GatewayResponse) GetBody() *CreateIpv6GatewayResponseBody {
	return s.Body
}

func (s *CreateIpv6GatewayResponse) SetHeaders(v map[string]*string) *CreateIpv6GatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateIpv6GatewayResponse) SetStatusCode(v int32) *CreateIpv6GatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpv6GatewayResponse) SetBody(v *CreateIpv6GatewayResponseBody) *CreateIpv6GatewayResponse {
	s.Body = v
	return s
}

func (s *CreateIpv6GatewayResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv4GatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpv4GatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpv4GatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpv4GatewayResponseBody) *CreateIpv4GatewayResponse
	GetBody() *CreateIpv4GatewayResponseBody
}

type CreateIpv4GatewayResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpv4GatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpv4GatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv4GatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateIpv4GatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpv4GatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpv4GatewayResponse) GetBody() *CreateIpv4GatewayResponseBody {
	return s.Body
}

func (s *CreateIpv4GatewayResponse) SetHeaders(v map[string]*string) *CreateIpv4GatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateIpv4GatewayResponse) SetStatusCode(v int32) *CreateIpv4GatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpv4GatewayResponse) SetBody(v *CreateIpv4GatewayResponseBody) *CreateIpv4GatewayResponse {
	s.Body = v
	return s
}

func (s *CreateIpv4GatewayResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6GatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpv6GatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpv6GatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpv6GatewayResponseBody) *DeleteIpv6GatewayResponse
	GetBody() *DeleteIpv6GatewayResponseBody
}

type DeleteIpv6GatewayResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpv6GatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpv6GatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6GatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpv6GatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpv6GatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpv6GatewayResponse) GetBody() *DeleteIpv6GatewayResponseBody {
	return s.Body
}

func (s *DeleteIpv6GatewayResponse) SetHeaders(v map[string]*string) *DeleteIpv6GatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpv6GatewayResponse) SetStatusCode(v int32) *DeleteIpv6GatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpv6GatewayResponse) SetBody(v *DeleteIpv6GatewayResponseBody) *DeleteIpv6GatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteIpv6GatewayResponse) Validate() error {
	return dara.Validate(s)
}

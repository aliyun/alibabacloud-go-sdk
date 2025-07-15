// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv4GatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpv4GatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpv4GatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpv4GatewayResponseBody) *DeleteIpv4GatewayResponse
	GetBody() *DeleteIpv4GatewayResponseBody
}

type DeleteIpv4GatewayResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpv4GatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpv4GatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv4GatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpv4GatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpv4GatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpv4GatewayResponse) GetBody() *DeleteIpv4GatewayResponseBody {
	return s.Body
}

func (s *DeleteIpv4GatewayResponse) SetHeaders(v map[string]*string) *DeleteIpv4GatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpv4GatewayResponse) SetStatusCode(v int32) *DeleteIpv4GatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpv4GatewayResponse) SetBody(v *DeleteIpv4GatewayResponseBody) *DeleteIpv4GatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteIpv4GatewayResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpnGatewayResponseBody) *DeleteVpnGatewayResponse
	GetBody() *DeleteVpnGatewayResponseBody
}

type DeleteVpnGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpnGatewayResponse) GetBody() *DeleteVpnGatewayResponseBody {
	return s.Body
}

func (s *DeleteVpnGatewayResponse) SetHeaders(v map[string]*string) *DeleteVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpnGatewayResponse) SetStatusCode(v int32) *DeleteVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpnGatewayResponse) SetBody(v *DeleteVpnGatewayResponseBody) *DeleteVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteVpnGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

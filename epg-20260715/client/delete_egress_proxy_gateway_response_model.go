// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEgressProxyGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEgressProxyGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEgressProxyGatewayResponseBody) *DeleteEgressProxyGatewayResponse
	GetBody() *DeleteEgressProxyGatewayResponseBody
}

type DeleteEgressProxyGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEgressProxyGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEgressProxyGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEgressProxyGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEgressProxyGatewayResponse) GetBody() *DeleteEgressProxyGatewayResponseBody {
	return s.Body
}

func (s *DeleteEgressProxyGatewayResponse) SetHeaders(v map[string]*string) *DeleteEgressProxyGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteEgressProxyGatewayResponse) SetStatusCode(v int32) *DeleteEgressProxyGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEgressProxyGatewayResponse) SetBody(v *DeleteEgressProxyGatewayResponseBody) *DeleteEgressProxyGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteEgressProxyGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSslVpnServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSslVpnServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSslVpnServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSslVpnServerResponseBody) *DeleteSslVpnServerResponse
	GetBody() *DeleteSslVpnServerResponseBody
}

type DeleteSslVpnServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSslVpnServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSslVpnServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSslVpnServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteSslVpnServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSslVpnServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSslVpnServerResponse) GetBody() *DeleteSslVpnServerResponseBody {
	return s.Body
}

func (s *DeleteSslVpnServerResponse) SetHeaders(v map[string]*string) *DeleteSslVpnServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteSslVpnServerResponse) SetStatusCode(v int32) *DeleteSslVpnServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSslVpnServerResponse) SetBody(v *DeleteSslVpnServerResponseBody) *DeleteSslVpnServerResponse {
	s.Body = v
	return s
}

func (s *DeleteSslVpnServerResponse) Validate() error {
	return dara.Validate(s)
}

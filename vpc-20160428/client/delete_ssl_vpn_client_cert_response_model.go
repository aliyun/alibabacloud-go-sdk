// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSslVpnClientCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSslVpnClientCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSslVpnClientCertResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSslVpnClientCertResponseBody) *DeleteSslVpnClientCertResponse
	GetBody() *DeleteSslVpnClientCertResponseBody
}

type DeleteSslVpnClientCertResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSslVpnClientCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSslVpnClientCertResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSslVpnClientCertResponse) GoString() string {
	return s.String()
}

func (s *DeleteSslVpnClientCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSslVpnClientCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSslVpnClientCertResponse) GetBody() *DeleteSslVpnClientCertResponseBody {
	return s.Body
}

func (s *DeleteSslVpnClientCertResponse) SetHeaders(v map[string]*string) *DeleteSslVpnClientCertResponse {
	s.Headers = v
	return s
}

func (s *DeleteSslVpnClientCertResponse) SetStatusCode(v int32) *DeleteSslVpnClientCertResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSslVpnClientCertResponse) SetBody(v *DeleteSslVpnClientCertResponseBody) *DeleteSslVpnClientCertResponse {
	s.Body = v
	return s
}

func (s *DeleteSslVpnClientCertResponse) Validate() error {
	return dara.Validate(s)
}

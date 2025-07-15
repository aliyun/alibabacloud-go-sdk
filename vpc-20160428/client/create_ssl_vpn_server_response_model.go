// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSslVpnServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSslVpnServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSslVpnServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateSslVpnServerResponseBody) *CreateSslVpnServerResponse
	GetBody() *CreateSslVpnServerResponseBody
}

type CreateSslVpnServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSslVpnServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSslVpnServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSslVpnServerResponse) GoString() string {
	return s.String()
}

func (s *CreateSslVpnServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSslVpnServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSslVpnServerResponse) GetBody() *CreateSslVpnServerResponseBody {
	return s.Body
}

func (s *CreateSslVpnServerResponse) SetHeaders(v map[string]*string) *CreateSslVpnServerResponse {
	s.Headers = v
	return s
}

func (s *CreateSslVpnServerResponse) SetStatusCode(v int32) *CreateSslVpnServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSslVpnServerResponse) SetBody(v *CreateSslVpnServerResponseBody) *CreateSslVpnServerResponse {
	s.Body = v
	return s
}

func (s *CreateSslVpnServerResponse) Validate() error {
	return dara.Validate(s)
}

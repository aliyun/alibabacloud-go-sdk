// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSslVpnClientCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSslVpnClientCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSslVpnClientCertResponse
	GetStatusCode() *int32
	SetBody(v *CreateSslVpnClientCertResponseBody) *CreateSslVpnClientCertResponse
	GetBody() *CreateSslVpnClientCertResponseBody
}

type CreateSslVpnClientCertResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSslVpnClientCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSslVpnClientCertResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSslVpnClientCertResponse) GoString() string {
	return s.String()
}

func (s *CreateSslVpnClientCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSslVpnClientCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSslVpnClientCertResponse) GetBody() *CreateSslVpnClientCertResponseBody {
	return s.Body
}

func (s *CreateSslVpnClientCertResponse) SetHeaders(v map[string]*string) *CreateSslVpnClientCertResponse {
	s.Headers = v
	return s
}

func (s *CreateSslVpnClientCertResponse) SetStatusCode(v int32) *CreateSslVpnClientCertResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSslVpnClientCertResponse) SetBody(v *CreateSslVpnClientCertResponseBody) *CreateSslVpnClientCertResponse {
	s.Body = v
	return s
}

func (s *CreateSslVpnClientCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

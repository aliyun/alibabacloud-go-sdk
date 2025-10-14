// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewSSLCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewSSLCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewSSLCertificateResponse
	GetStatusCode() *int32
	SetBody(v *RenewSSLCertificateResponseBody) *RenewSSLCertificateResponse
	GetBody() *RenewSSLCertificateResponseBody
}

type RenewSSLCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewSSLCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *RenewSSLCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewSSLCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewSSLCertificateResponse) GetBody() *RenewSSLCertificateResponseBody {
	return s.Body
}

func (s *RenewSSLCertificateResponse) SetHeaders(v map[string]*string) *RenewSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *RenewSSLCertificateResponse) SetStatusCode(v int32) *RenewSSLCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewSSLCertificateResponse) SetBody(v *RenewSSLCertificateResponseBody) *RenewSSLCertificateResponse {
	s.Body = v
	return s
}

func (s *RenewSSLCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

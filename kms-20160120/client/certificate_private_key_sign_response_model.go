// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePrivateKeySignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CertificatePrivateKeySignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CertificatePrivateKeySignResponse
	GetStatusCode() *int32
	SetBody(v *CertificatePrivateKeySignResponseBody) *CertificatePrivateKeySignResponse
	GetBody() *CertificatePrivateKeySignResponseBody
}

type CertificatePrivateKeySignResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePrivateKeySignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePrivateKeySignResponse) String() string {
	return dara.Prettify(s)
}

func (s CertificatePrivateKeySignResponse) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CertificatePrivateKeySignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CertificatePrivateKeySignResponse) GetBody() *CertificatePrivateKeySignResponseBody {
	return s.Body
}

func (s *CertificatePrivateKeySignResponse) SetHeaders(v map[string]*string) *CertificatePrivateKeySignResponse {
	s.Headers = v
	return s
}

func (s *CertificatePrivateKeySignResponse) SetStatusCode(v int32) *CertificatePrivateKeySignResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePrivateKeySignResponse) SetBody(v *CertificatePrivateKeySignResponseBody) *CertificatePrivateKeySignResponse {
	s.Body = v
	return s
}

func (s *CertificatePrivateKeySignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

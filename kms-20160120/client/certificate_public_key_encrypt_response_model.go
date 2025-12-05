// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePublicKeyEncryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CertificatePublicKeyEncryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CertificatePublicKeyEncryptResponse
	GetStatusCode() *int32
	SetBody(v *CertificatePublicKeyEncryptResponseBody) *CertificatePublicKeyEncryptResponse
	GetBody() *CertificatePublicKeyEncryptResponseBody
}

type CertificatePublicKeyEncryptResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePublicKeyEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePublicKeyEncryptResponse) String() string {
	return dara.Prettify(s)
}

func (s CertificatePublicKeyEncryptResponse) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CertificatePublicKeyEncryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CertificatePublicKeyEncryptResponse) GetBody() *CertificatePublicKeyEncryptResponseBody {
	return s.Body
}

func (s *CertificatePublicKeyEncryptResponse) SetHeaders(v map[string]*string) *CertificatePublicKeyEncryptResponse {
	s.Headers = v
	return s
}

func (s *CertificatePublicKeyEncryptResponse) SetStatusCode(v int32) *CertificatePublicKeyEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponse) SetBody(v *CertificatePublicKeyEncryptResponseBody) *CertificatePublicKeyEncryptResponse {
	s.Body = v
	return s
}

func (s *CertificatePublicKeyEncryptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

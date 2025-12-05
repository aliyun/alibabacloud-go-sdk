// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePublicKeyVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CertificatePublicKeyVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CertificatePublicKeyVerifyResponse
	GetStatusCode() *int32
	SetBody(v *CertificatePublicKeyVerifyResponseBody) *CertificatePublicKeyVerifyResponse
	GetBody() *CertificatePublicKeyVerifyResponseBody
}

type CertificatePublicKeyVerifyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePublicKeyVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePublicKeyVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s CertificatePublicKeyVerifyResponse) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CertificatePublicKeyVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CertificatePublicKeyVerifyResponse) GetBody() *CertificatePublicKeyVerifyResponseBody {
	return s.Body
}

func (s *CertificatePublicKeyVerifyResponse) SetHeaders(v map[string]*string) *CertificatePublicKeyVerifyResponse {
	s.Headers = v
	return s
}

func (s *CertificatePublicKeyVerifyResponse) SetStatusCode(v int32) *CertificatePublicKeyVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponse) SetBody(v *CertificatePublicKeyVerifyResponseBody) *CertificatePublicKeyVerifyResponse {
	s.Body = v
	return s
}

func (s *CertificatePublicKeyVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

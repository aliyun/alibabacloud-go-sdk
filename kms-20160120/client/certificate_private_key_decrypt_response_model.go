// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePrivateKeyDecryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CertificatePrivateKeyDecryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CertificatePrivateKeyDecryptResponse
	GetStatusCode() *int32
	SetBody(v *CertificatePrivateKeyDecryptResponseBody) *CertificatePrivateKeyDecryptResponse
	GetBody() *CertificatePrivateKeyDecryptResponseBody
}

type CertificatePrivateKeyDecryptResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertificatePrivateKeyDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertificatePrivateKeyDecryptResponse) String() string {
	return dara.Prettify(s)
}

func (s CertificatePrivateKeyDecryptResponse) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CertificatePrivateKeyDecryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CertificatePrivateKeyDecryptResponse) GetBody() *CertificatePrivateKeyDecryptResponseBody {
	return s.Body
}

func (s *CertificatePrivateKeyDecryptResponse) SetHeaders(v map[string]*string) *CertificatePrivateKeyDecryptResponse {
	s.Headers = v
	return s
}

func (s *CertificatePrivateKeyDecryptResponse) SetStatusCode(v int32) *CertificatePrivateKeyDecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponse) SetBody(v *CertificatePrivateKeyDecryptResponseBody) *CertificatePrivateKeyDecryptResponse {
	s.Body = v
	return s
}

func (s *CertificatePrivateKeyDecryptResponse) Validate() error {
	return dara.Validate(s)
}

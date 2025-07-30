// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificatePrivateKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCertificatePrivateKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCertificatePrivateKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCertificatePrivateKeyResponseBody) *DescribeCertificatePrivateKeyResponse
	GetBody() *DescribeCertificatePrivateKeyResponseBody
}

type DescribeCertificatePrivateKeyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificatePrivateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificatePrivateKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificatePrivateKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificatePrivateKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCertificatePrivateKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCertificatePrivateKeyResponse) GetBody() *DescribeCertificatePrivateKeyResponseBody {
	return s.Body
}

func (s *DescribeCertificatePrivateKeyResponse) SetHeaders(v map[string]*string) *DescribeCertificatePrivateKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificatePrivateKeyResponse) SetStatusCode(v int32) *DescribeCertificatePrivateKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificatePrivateKeyResponse) SetBody(v *DescribeCertificatePrivateKeyResponseBody) *DescribeCertificatePrivateKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeCertificatePrivateKeyResponse) Validate() error {
	return dara.Validate(s)
}

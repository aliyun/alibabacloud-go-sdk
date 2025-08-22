// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSSLCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSSLCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSSLCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSSLCertificateListResponseBody) *DescribeDcdnSSLCertificateListResponse
	GetBody() *DescribeDcdnSSLCertificateListResponseBody
}

type DescribeDcdnSSLCertificateListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSSLCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSSLCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSSLCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSSLCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSSLCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSSLCertificateListResponse) GetBody() *DescribeDcdnSSLCertificateListResponseBody {
	return s.Body
}

func (s *DescribeDcdnSSLCertificateListResponse) SetHeaders(v map[string]*string) *DescribeDcdnSSLCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponse) SetStatusCode(v int32) *DescribeDcdnSSLCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponse) SetBody(v *DescribeDcdnSSLCertificateListResponseBody) *DescribeDcdnSSLCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

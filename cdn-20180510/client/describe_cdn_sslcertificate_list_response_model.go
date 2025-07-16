// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSSLCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnSSLCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnSSLCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnSSLCertificateListResponseBody) *DescribeCdnSSLCertificateListResponse
	GetBody() *DescribeCdnSSLCertificateListResponseBody
}

type DescribeCdnSSLCertificateListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnSSLCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnSSLCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSSLCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnSSLCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnSSLCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnSSLCertificateListResponse) GetBody() *DescribeCdnSSLCertificateListResponseBody {
	return s.Body
}

func (s *DescribeCdnSSLCertificateListResponse) SetHeaders(v map[string]*string) *DescribeCdnSSLCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnSSLCertificateListResponse) SetStatusCode(v int32) *DescribeCdnSSLCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponse) SetBody(v *DescribeCdnSSLCertificateListResponseBody) *DescribeCdnSSLCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnSSLCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

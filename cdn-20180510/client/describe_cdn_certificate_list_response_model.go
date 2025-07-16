// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnCertificateListResponseBody) *DescribeCdnCertificateListResponse
	GetBody() *DescribeCdnCertificateListResponseBody
}

type DescribeCdnCertificateListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnCertificateListResponse) GetBody() *DescribeCdnCertificateListResponseBody {
	return s.Body
}

func (s *DescribeCdnCertificateListResponse) SetHeaders(v map[string]*string) *DescribeCdnCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnCertificateListResponse) SetStatusCode(v int32) *DescribeCdnCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnCertificateListResponse) SetBody(v *DescribeCdnCertificateListResponseBody) *DescribeCdnCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

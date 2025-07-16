// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnCertificateDetailResponseBody) *DescribeCdnCertificateDetailResponse
	GetBody() *DescribeCdnCertificateDetailResponseBody
}

type DescribeCdnCertificateDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnCertificateDetailResponse) GetBody() *DescribeCdnCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeCdnCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeCdnCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnCertificateDetailResponse) SetStatusCode(v int32) *DescribeCdnCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponse) SetBody(v *DescribeCdnCertificateDetailResponseBody) *DescribeCdnCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnCertificateDetailResponse) Validate() error {
	return dara.Validate(s)
}

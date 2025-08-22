// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnCertificateDetailResponseBody) *DescribeDcdnCertificateDetailResponse
	GetBody() *DescribeDcdnCertificateDetailResponseBody
}

type DescribeDcdnCertificateDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnCertificateDetailResponse) GetBody() *DescribeDcdnCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeDcdnCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnCertificateDetailResponse) SetStatusCode(v int32) *DescribeDcdnCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponse) SetBody(v *DescribeDcdnCertificateDetailResponseBody) *DescribeDcdnCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnCertificateDetailResponse) Validate() error {
	return dara.Validate(s)
}

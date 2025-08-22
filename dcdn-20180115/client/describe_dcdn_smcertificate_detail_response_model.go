// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSMCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSMCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSMCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSMCertificateDetailResponseBody) *DescribeDcdnSMCertificateDetailResponse
	GetBody() *DescribeDcdnSMCertificateDetailResponseBody
}

type DescribeDcdnSMCertificateDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSMCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSMCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSMCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSMCertificateDetailResponse) GetBody() *DescribeDcdnSMCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeDcdnSMCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnSMCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponse) SetStatusCode(v int32) *DescribeDcdnSMCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponse) SetBody(v *DescribeDcdnSMCertificateDetailResponseBody) *DescribeDcdnSMCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveCertificateDetailResponseBody) *DescribeLiveCertificateDetailResponse
	GetBody() *DescribeLiveCertificateDetailResponseBody
}

type DescribeLiveCertificateDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveCertificateDetailResponse) GetBody() *DescribeLiveCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeLiveCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeLiveCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) SetStatusCode(v int32) *DescribeLiveCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) SetBody(v *DescribeLiveCertificateDetailResponseBody) *DescribeLiveCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) Validate() error {
	return dara.Validate(s)
}

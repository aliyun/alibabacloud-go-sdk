// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsCertificateDetailResponseBody) *DescribeVsCertificateDetailResponse
	GetBody() *DescribeVsCertificateDetailResponseBody
}

type DescribeVsCertificateDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsCertificateDetailResponse) GetBody() *DescribeVsCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeVsCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeVsCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsCertificateDetailResponse) SetStatusCode(v int32) *DescribeVsCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsCertificateDetailResponse) SetBody(v *DescribeVsCertificateDetailResponseBody) *DescribeVsCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVsCertificateDetailResponse) Validate() error {
	return dara.Validate(s)
}

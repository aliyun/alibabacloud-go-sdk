// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodCertificateDetailResponseBody) *DescribeVodCertificateDetailResponse
	GetBody() *DescribeVodCertificateDetailResponseBody
}

type DescribeVodCertificateDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodCertificateDetailResponse) GetBody() *DescribeVodCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeVodCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeVodCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodCertificateDetailResponse) SetStatusCode(v int32) *DescribeVodCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodCertificateDetailResponse) SetBody(v *DescribeVodCertificateDetailResponseBody) *DescribeVodCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVodCertificateDetailResponse) Validate() error {
	return dara.Validate(s)
}

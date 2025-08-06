// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodCertificateListResponseBody) *DescribeVodCertificateListResponse
	GetBody() *DescribeVodCertificateListResponseBody
}

type DescribeVodCertificateListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodCertificateListResponse) GetBody() *DescribeVodCertificateListResponseBody {
	return s.Body
}

func (s *DescribeVodCertificateListResponse) SetHeaders(v map[string]*string) *DescribeVodCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodCertificateListResponse) SetStatusCode(v int32) *DescribeVodCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodCertificateListResponse) SetBody(v *DescribeVodCertificateListResponseBody) *DescribeVodCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeVodCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

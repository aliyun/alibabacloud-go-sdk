// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodSSLCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodSSLCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodSSLCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodSSLCertificateListResponseBody) *DescribeVodSSLCertificateListResponse
	GetBody() *DescribeVodSSLCertificateListResponseBody
}

type DescribeVodSSLCertificateListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodSSLCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodSSLCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodSSLCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodSSLCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodSSLCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodSSLCertificateListResponse) GetBody() *DescribeVodSSLCertificateListResponseBody {
	return s.Body
}

func (s *DescribeVodSSLCertificateListResponse) SetHeaders(v map[string]*string) *DescribeVodSSLCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodSSLCertificateListResponse) SetStatusCode(v int32) *DescribeVodSSLCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponse) SetBody(v *DescribeVodSSLCertificateListResponseBody) *DescribeVodSSLCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeVodSSLCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsCertificateListResponseBody) *DescribeVsCertificateListResponse
	GetBody() *DescribeVsCertificateListResponseBody
}

type DescribeVsCertificateListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsCertificateListResponse) GetBody() *DescribeVsCertificateListResponseBody {
	return s.Body
}

func (s *DescribeVsCertificateListResponse) SetHeaders(v map[string]*string) *DescribeVsCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsCertificateListResponse) SetStatusCode(v int32) *DescribeVsCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsCertificateListResponse) SetBody(v *DescribeVsCertificateListResponseBody) *DescribeVsCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeVsCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

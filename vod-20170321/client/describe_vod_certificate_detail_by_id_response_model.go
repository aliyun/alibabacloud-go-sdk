// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateDetailByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodCertificateDetailByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodCertificateDetailByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodCertificateDetailByIdResponseBody) *DescribeVodCertificateDetailByIdResponse
	GetBody() *DescribeVodCertificateDetailByIdResponseBody
}

type DescribeVodCertificateDetailByIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodCertificateDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodCertificateDetailByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateDetailByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodCertificateDetailByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodCertificateDetailByIdResponse) GetBody() *DescribeVodCertificateDetailByIdResponseBody {
	return s.Body
}

func (s *DescribeVodCertificateDetailByIdResponse) SetHeaders(v map[string]*string) *DescribeVodCertificateDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponse) SetStatusCode(v int32) *DescribeVodCertificateDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponse) SetBody(v *DescribeVodCertificateDetailByIdResponseBody) *DescribeVodCertificateDetailByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponse) Validate() error {
	return dara.Validate(s)
}

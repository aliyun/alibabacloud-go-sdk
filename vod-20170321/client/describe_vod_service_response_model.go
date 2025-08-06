// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodServiceResponseBody) *DescribeVodServiceResponse
	GetBody() *DescribeVodServiceResponseBody
}

type DescribeVodServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodServiceResponse) GetBody() *DescribeVodServiceResponseBody {
	return s.Body
}

func (s *DescribeVodServiceResponse) SetHeaders(v map[string]*string) *DescribeVodServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodServiceResponse) SetStatusCode(v int32) *DescribeVodServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodServiceResponse) SetBody(v *DescribeVodServiceResponseBody) *DescribeVodServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeVodServiceResponse) Validate() error {
	return dara.Validate(s)
}

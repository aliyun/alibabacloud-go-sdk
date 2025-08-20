// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterResourceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterResourceDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterResourceDetailResponseBody) *DescribeClusterResourceDetailResponse
	GetBody() *DescribeClusterResourceDetailResponseBody
}

type DescribeClusterResourceDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResourceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResourceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterResourceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterResourceDetailResponse) GetBody() *DescribeClusterResourceDetailResponseBody {
	return s.Body
}

func (s *DescribeClusterResourceDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterResourceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResourceDetailResponse) SetStatusCode(v int32) *DescribeClusterResourceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResourceDetailResponse) SetBody(v *DescribeClusterResourceDetailResponseBody) *DescribeClusterResourceDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterResourceDetailResponse) Validate() error {
	return dara.Validate(s)
}

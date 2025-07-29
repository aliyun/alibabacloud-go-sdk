// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterNodePoolDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterNodePoolDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterNodePoolDetailResponseBody) *DescribeClusterNodePoolDetailResponse
	GetBody() *DescribeClusterNodePoolDetailResponseBody
}

type DescribeClusterNodePoolDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterNodePoolDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterNodePoolDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterNodePoolDetailResponse) GetBody() *DescribeClusterNodePoolDetailResponseBody {
	return s.Body
}

func (s *DescribeClusterNodePoolDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterNodePoolDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponse) SetStatusCode(v int32) *DescribeClusterNodePoolDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponse) SetBody(v *DescribeClusterNodePoolDetailResponseBody) *DescribeClusterNodePoolDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponse) Validate() error {
	return dara.Validate(s)
}

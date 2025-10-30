// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterNodePoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterNodePoolsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterNodePoolsResponseBody) *DescribeClusterNodePoolsResponse
	GetBody() *DescribeClusterNodePoolsResponseBody
}

type DescribeClusterNodePoolsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterNodePoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterNodePoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterNodePoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterNodePoolsResponse) GetBody() *DescribeClusterNodePoolsResponseBody {
	return s.Body
}

func (s *DescribeClusterNodePoolsResponse) SetHeaders(v map[string]*string) *DescribeClusterNodePoolsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodePoolsResponse) SetStatusCode(v int32) *DescribeClusterNodePoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterNodePoolsResponse) SetBody(v *DescribeClusterNodePoolsResponseBody) *DescribeClusterNodePoolsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterNodePoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

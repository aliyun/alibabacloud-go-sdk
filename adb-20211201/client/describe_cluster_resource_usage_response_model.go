// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterResourceUsageResponseBody) *DescribeClusterResourceUsageResponse
	GetBody() *DescribeClusterResourceUsageResponseBody
}

type DescribeClusterResourceUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterResourceUsageResponse) GetBody() *DescribeClusterResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeClusterResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeClusterResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResourceUsageResponse) SetStatusCode(v int32) *DescribeClusterResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResourceUsageResponse) SetBody(v *DescribeClusterResourceUsageResponseBody) *DescribeClusterResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterResourceUsageResponse) Validate() error {
	return dara.Validate(s)
}

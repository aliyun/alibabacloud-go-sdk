// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterVulStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterVulStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterVulStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterVulStatisticsResponseBody) *DescribeClusterVulStatisticsResponse
	GetBody() *DescribeClusterVulStatisticsResponseBody
}

type DescribeClusterVulStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterVulStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterVulStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterVulStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterVulStatisticsResponse) GetBody() *DescribeClusterVulStatisticsResponseBody {
	return s.Body
}

func (s *DescribeClusterVulStatisticsResponse) SetHeaders(v map[string]*string) *DescribeClusterVulStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterVulStatisticsResponse) SetStatusCode(v int32) *DescribeClusterVulStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterVulStatisticsResponse) SetBody(v *DescribeClusterVulStatisticsResponseBody) *DescribeClusterVulStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterVulStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulFixStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulFixStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulFixStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulFixStatisticsResponseBody) *DescribeVulFixStatisticsResponse
	GetBody() *DescribeVulFixStatisticsResponseBody
}

type DescribeVulFixStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulFixStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulFixStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulFixStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulFixStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulFixStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulFixStatisticsResponse) GetBody() *DescribeVulFixStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVulFixStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVulFixStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulFixStatisticsResponse) SetStatusCode(v int32) *DescribeVulFixStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulFixStatisticsResponse) SetBody(v *DescribeVulFixStatisticsResponseBody) *DescribeVulFixStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulFixStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

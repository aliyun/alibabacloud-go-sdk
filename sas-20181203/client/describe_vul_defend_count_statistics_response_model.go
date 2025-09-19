// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDefendCountStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulDefendCountStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulDefendCountStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulDefendCountStatisticsResponseBody) *DescribeVulDefendCountStatisticsResponse
	GetBody() *DescribeVulDefendCountStatisticsResponseBody
}

type DescribeVulDefendCountStatisticsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulDefendCountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulDefendCountStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDefendCountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDefendCountStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulDefendCountStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulDefendCountStatisticsResponse) GetBody() *DescribeVulDefendCountStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVulDefendCountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVulDefendCountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulDefendCountStatisticsResponse) SetStatusCode(v int32) *DescribeVulDefendCountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulDefendCountStatisticsResponse) SetBody(v *DescribeVulDefendCountStatisticsResponseBody) *DescribeVulDefendCountStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulDefendCountStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

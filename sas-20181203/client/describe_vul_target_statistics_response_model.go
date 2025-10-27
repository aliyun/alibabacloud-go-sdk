// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulTargetStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulTargetStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulTargetStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulTargetStatisticsResponseBody) *DescribeVulTargetStatisticsResponse
	GetBody() *DescribeVulTargetStatisticsResponseBody
}

type DescribeVulTargetStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulTargetStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulTargetStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulTargetStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulTargetStatisticsResponse) GetBody() *DescribeVulTargetStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVulTargetStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVulTargetStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulTargetStatisticsResponse) SetStatusCode(v int32) *DescribeVulTargetStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponse) SetBody(v *DescribeVulTargetStatisticsResponseBody) *DescribeVulTargetStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulTargetStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

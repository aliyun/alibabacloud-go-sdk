// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulNumStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulNumStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulNumStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulNumStatisticsResponseBody) *DescribeVulNumStatisticsResponse
	GetBody() *DescribeVulNumStatisticsResponseBody
}

type DescribeVulNumStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulNumStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulNumStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulNumStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulNumStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulNumStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulNumStatisticsResponse) GetBody() *DescribeVulNumStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVulNumStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVulNumStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulNumStatisticsResponse) SetStatusCode(v int32) *DescribeVulNumStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulNumStatisticsResponse) SetBody(v *DescribeVulNumStatisticsResponseBody) *DescribeVulNumStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulNumStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

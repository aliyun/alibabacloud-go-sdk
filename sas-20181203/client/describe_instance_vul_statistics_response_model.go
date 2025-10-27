// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVulStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceVulStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceVulStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceVulStatisticsResponseBody) *DescribeInstanceVulStatisticsResponse
	GetBody() *DescribeInstanceVulStatisticsResponseBody
}

type DescribeInstanceVulStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceVulStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceVulStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVulStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVulStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceVulStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceVulStatisticsResponse) GetBody() *DescribeInstanceVulStatisticsResponseBody {
	return s.Body
}

func (s *DescribeInstanceVulStatisticsResponse) SetHeaders(v map[string]*string) *DescribeInstanceVulStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceVulStatisticsResponse) SetStatusCode(v int32) *DescribeInstanceVulStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceVulStatisticsResponse) SetBody(v *DescribeInstanceVulStatisticsResponseBody) *DescribeInstanceVulStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceVulStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

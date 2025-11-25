// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceStatisticsResponseBody) *DescribeInstanceStatisticsResponse
	GetBody() *DescribeInstanceStatisticsResponseBody
}

type DescribeInstanceStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceStatisticsResponse) GetBody() *DescribeInstanceStatisticsResponseBody {
	return s.Body
}

func (s *DescribeInstanceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetStatusCode(v int32) *DescribeInstanceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetBody(v *DescribeInstanceStatisticsResponseBody) *DescribeInstanceStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

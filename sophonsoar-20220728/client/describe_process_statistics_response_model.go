// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProcessStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProcessStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProcessStatisticsResponseBody) *DescribeProcessStatisticsResponse
	GetBody() *DescribeProcessStatisticsResponseBody
}

type DescribeProcessStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProcessStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProcessStatisticsResponse) GetBody() *DescribeProcessStatisticsResponseBody {
	return s.Body
}

func (s *DescribeProcessStatisticsResponse) SetHeaders(v map[string]*string) *DescribeProcessStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessStatisticsResponse) SetStatusCode(v int32) *DescribeProcessStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessStatisticsResponse) SetBody(v *DescribeProcessStatisticsResponseBody) *DescribeProcessStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeProcessStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

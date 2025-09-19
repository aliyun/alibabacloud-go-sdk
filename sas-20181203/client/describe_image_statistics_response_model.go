// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageStatisticsResponseBody) *DescribeImageStatisticsResponse
	GetBody() *DescribeImageStatisticsResponseBody
}

type DescribeImageStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageStatisticsResponse) GetBody() *DescribeImageStatisticsResponseBody {
	return s.Body
}

func (s *DescribeImageStatisticsResponse) SetHeaders(v map[string]*string) *DescribeImageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageStatisticsResponse) SetStatusCode(v int32) *DescribeImageStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageStatisticsResponse) SetBody(v *DescribeImageStatisticsResponseBody) *DescribeImageStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeImageStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

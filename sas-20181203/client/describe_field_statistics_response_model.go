// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFieldStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFieldStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFieldStatisticsResponseBody) *DescribeFieldStatisticsResponse
	GetBody() *DescribeFieldStatisticsResponseBody
}

type DescribeFieldStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFieldStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFieldStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFieldStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFieldStatisticsResponse) GetBody() *DescribeFieldStatisticsResponseBody {
	return s.Body
}

func (s *DescribeFieldStatisticsResponse) SetHeaders(v map[string]*string) *DescribeFieldStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldStatisticsResponse) SetStatusCode(v int32) *DescribeFieldStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldStatisticsResponse) SetBody(v *DescribeFieldStatisticsResponseBody) *DescribeFieldStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeFieldStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTableStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTableStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTableStatisticsResponseBody) *DescribeTableStatisticsResponse
	GetBody() *DescribeTableStatisticsResponseBody
}

type DescribeTableStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTableStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTableStatisticsResponse) GetBody() *DescribeTableStatisticsResponseBody {
	return s.Body
}

func (s *DescribeTableStatisticsResponse) SetHeaders(v map[string]*string) *DescribeTableStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableStatisticsResponse) SetStatusCode(v int32) *DescribeTableStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableStatisticsResponse) SetBody(v *DescribeTableStatisticsResponseBody) *DescribeTableStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeTableStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

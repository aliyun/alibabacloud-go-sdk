// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExposedStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExposedStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExposedStatisticsResponseBody) *DescribeExposedStatisticsResponse
	GetBody() *DescribeExposedStatisticsResponseBody
}

type DescribeExposedStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExposedStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExposedStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExposedStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExposedStatisticsResponse) GetBody() *DescribeExposedStatisticsResponseBody {
	return s.Body
}

func (s *DescribeExposedStatisticsResponse) SetHeaders(v map[string]*string) *DescribeExposedStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedStatisticsResponse) SetStatusCode(v int32) *DescribeExposedStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExposedStatisticsResponse) SetBody(v *DescribeExposedStatisticsResponseBody) *DescribeExposedStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeExposedStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

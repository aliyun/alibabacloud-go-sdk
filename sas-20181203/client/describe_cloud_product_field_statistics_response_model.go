// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudProductFieldStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudProductFieldStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudProductFieldStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudProductFieldStatisticsResponseBody) *DescribeCloudProductFieldStatisticsResponse
	GetBody() *DescribeCloudProductFieldStatisticsResponseBody
}

type DescribeCloudProductFieldStatisticsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudProductFieldStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudProductFieldStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudProductFieldStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudProductFieldStatisticsResponse) GetBody() *DescribeCloudProductFieldStatisticsResponseBody {
	return s.Body
}

func (s *DescribeCloudProductFieldStatisticsResponse) SetHeaders(v map[string]*string) *DescribeCloudProductFieldStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponse) SetStatusCode(v int32) *DescribeCloudProductFieldStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponse) SetBody(v *DescribeCloudProductFieldStatisticsResponseBody) *DescribeCloudProductFieldStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

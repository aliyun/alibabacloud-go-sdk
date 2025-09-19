// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedStatisticsDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExposedStatisticsDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExposedStatisticsDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExposedStatisticsDetailResponseBody) *DescribeExposedStatisticsDetailResponse
	GetBody() *DescribeExposedStatisticsDetailResponseBody
}

type DescribeExposedStatisticsDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExposedStatisticsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExposedStatisticsDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExposedStatisticsDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExposedStatisticsDetailResponse) GetBody() *DescribeExposedStatisticsDetailResponseBody {
	return s.Body
}

func (s *DescribeExposedStatisticsDetailResponse) SetHeaders(v map[string]*string) *DescribeExposedStatisticsDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedStatisticsDetailResponse) SetStatusCode(v int32) *DescribeExposedStatisticsDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponse) SetBody(v *DescribeExposedStatisticsDetailResponseBody) *DescribeExposedStatisticsDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeExposedStatisticsDetailResponse) Validate() error {
	return dara.Validate(s)
}

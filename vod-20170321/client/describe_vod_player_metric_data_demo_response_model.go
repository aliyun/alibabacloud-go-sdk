// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataDemoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodPlayerMetricDataDemoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodPlayerMetricDataDemoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodPlayerMetricDataDemoResponseBody) *DescribeVodPlayerMetricDataDemoResponse
	GetBody() *DescribeVodPlayerMetricDataDemoResponseBody
}

type DescribeVodPlayerMetricDataDemoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodPlayerMetricDataDemoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodPlayerMetricDataDemoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataDemoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataDemoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodPlayerMetricDataDemoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodPlayerMetricDataDemoResponse) GetBody() *DescribeVodPlayerMetricDataDemoResponseBody {
	return s.Body
}

func (s *DescribeVodPlayerMetricDataDemoResponse) SetHeaders(v map[string]*string) *DescribeVodPlayerMetricDataDemoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponse) SetStatusCode(v int32) *DescribeVodPlayerMetricDataDemoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponse) SetBody(v *DescribeVodPlayerMetricDataDemoResponseBody) *DescribeVodPlayerMetricDataDemoResponse {
	s.Body = v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoResponse) Validate() error {
	return dara.Validate(s)
}

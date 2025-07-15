// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobMetricDataResponseBody) *DescribeJobMetricDataResponse
	GetBody() *DescribeJobMetricDataResponseBody
}

type DescribeJobMetricDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobMetricDataResponse) GetBody() *DescribeJobMetricDataResponseBody {
	return s.Body
}

func (s *DescribeJobMetricDataResponse) SetHeaders(v map[string]*string) *DescribeJobMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobMetricDataResponse) SetStatusCode(v int32) *DescribeJobMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobMetricDataResponse) SetBody(v *DescribeJobMetricDataResponseBody) *DescribeJobMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeJobMetricDataResponse) Validate() error {
	return dara.Validate(s)
}

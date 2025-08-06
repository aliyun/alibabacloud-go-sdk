// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayMetricDataResponseBody) *DescribePlayMetricDataResponse
	GetBody() *DescribePlayMetricDataResponseBody
}

type DescribePlayMetricDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayMetricDataResponse) GetBody() *DescribePlayMetricDataResponseBody {
	return s.Body
}

func (s *DescribePlayMetricDataResponse) SetHeaders(v map[string]*string) *DescribePlayMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayMetricDataResponse) SetStatusCode(v int32) *DescribePlayMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayMetricDataResponse) SetBody(v *DescribePlayMetricDataResponseBody) *DescribePlayMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribePlayMetricDataResponse) Validate() error {
	return dara.Validate(s)
}

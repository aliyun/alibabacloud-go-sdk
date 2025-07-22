// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQoeMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQoeMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQoeMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQoeMetricDataResponseBody) *DescribeQoeMetricDataResponse
	GetBody() *DescribeQoeMetricDataResponseBody
}

type DescribeQoeMetricDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQoeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQoeMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQoeMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQoeMetricDataResponse) GetBody() *DescribeQoeMetricDataResponseBody {
	return s.Body
}

func (s *DescribeQoeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeQoeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQoeMetricDataResponse) SetStatusCode(v int32) *DescribeQoeMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQoeMetricDataResponse) SetBody(v *DescribeQoeMetricDataResponseBody) *DescribeQoeMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeQoeMetricDataResponse) Validate() error {
	return dara.Validate(s)
}

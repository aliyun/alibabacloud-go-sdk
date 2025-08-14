// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventResultBarChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventResultBarChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventResultBarChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventResultBarChartResponseBody) *DescribeEventResultBarChartResponse
	GetBody() *DescribeEventResultBarChartResponseBody
}

type DescribeEventResultBarChartResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventResultBarChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventResultBarChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventResultBarChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventResultBarChartResponse) GetBody() *DescribeEventResultBarChartResponseBody {
	return s.Body
}

func (s *DescribeEventResultBarChartResponse) SetHeaders(v map[string]*string) *DescribeEventResultBarChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventResultBarChartResponse) SetStatusCode(v int32) *DescribeEventResultBarChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventResultBarChartResponse) SetBody(v *DescribeEventResultBarChartResponseBody) *DescribeEventResultBarChartResponse {
	s.Body = v
	return s
}

func (s *DescribeEventResultBarChartResponse) Validate() error {
	return dara.Validate(s)
}

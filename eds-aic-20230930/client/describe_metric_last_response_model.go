// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricLastResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricLastResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricLastResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricLastResponseBody) *DescribeMetricLastResponse
	GetBody() *DescribeMetricLastResponseBody
}

type DescribeMetricLastResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricLastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricLastResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricLastResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricLastResponse) GetBody() *DescribeMetricLastResponseBody {
	return s.Body
}

func (s *DescribeMetricLastResponse) SetHeaders(v map[string]*string) *DescribeMetricLastResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricLastResponse) SetStatusCode(v int32) *DescribeMetricLastResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricLastResponse) SetBody(v *DescribeMetricLastResponseBody) *DescribeMetricLastResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricLastResponse) Validate() error {
	return dara.Validate(s)
}

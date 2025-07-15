// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricLastResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobMetricLastResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobMetricLastResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobMetricLastResponseBody) *DescribeJobMetricLastResponse
	GetBody() *DescribeJobMetricLastResponseBody
}

type DescribeJobMetricLastResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobMetricLastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobMetricLastResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricLastResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobMetricLastResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobMetricLastResponse) GetBody() *DescribeJobMetricLastResponseBody {
	return s.Body
}

func (s *DescribeJobMetricLastResponse) SetHeaders(v map[string]*string) *DescribeJobMetricLastResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobMetricLastResponse) SetStatusCode(v int32) *DescribeJobMetricLastResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobMetricLastResponse) SetBody(v *DescribeJobMetricLastResponseBody) *DescribeJobMetricLastResponse {
	s.Body = v
	return s
}

func (s *DescribeJobMetricLastResponse) Validate() error {
	return dara.Validate(s)
}

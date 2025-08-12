// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricTopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricTopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricTopResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricTopResponseBody) *DescribeMetricTopResponse
	GetBody() *DescribeMetricTopResponseBody
}

type DescribeMetricTopResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricTopResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricTopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricTopResponse) GetBody() *DescribeMetricTopResponseBody {
	return s.Body
}

func (s *DescribeMetricTopResponse) SetHeaders(v map[string]*string) *DescribeMetricTopResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricTopResponse) SetStatusCode(v int32) *DescribeMetricTopResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricTopResponse) SetBody(v *DescribeMetricTopResponseBody) *DescribeMetricTopResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricTopResponse) Validate() error {
	return dara.Validate(s)
}

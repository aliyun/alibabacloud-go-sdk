// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayMetricAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayMetricAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayMetricAuthResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayMetricAuthResponseBody) *DescribePlayMetricAuthResponse
	GetBody() *DescribePlayMetricAuthResponseBody
}

type DescribePlayMetricAuthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayMetricAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayMetricAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricAuthResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayMetricAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayMetricAuthResponse) GetBody() *DescribePlayMetricAuthResponseBody {
	return s.Body
}

func (s *DescribePlayMetricAuthResponse) SetHeaders(v map[string]*string) *DescribePlayMetricAuthResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayMetricAuthResponse) SetStatusCode(v int32) *DescribePlayMetricAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayMetricAuthResponse) SetBody(v *DescribePlayMetricAuthResponseBody) *DescribePlayMetricAuthResponse {
	s.Body = v
	return s
}

func (s *DescribePlayMetricAuthResponse) Validate() error {
	return dara.Validate(s)
}

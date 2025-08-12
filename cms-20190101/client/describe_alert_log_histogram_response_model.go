// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogHistogramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertLogHistogramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertLogHistogramResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertLogHistogramResponseBody) *DescribeAlertLogHistogramResponse
	GetBody() *DescribeAlertLogHistogramResponseBody
}

type DescribeAlertLogHistogramResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertLogHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertLogHistogramResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertLogHistogramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertLogHistogramResponse) GetBody() *DescribeAlertLogHistogramResponseBody {
	return s.Body
}

func (s *DescribeAlertLogHistogramResponse) SetHeaders(v map[string]*string) *DescribeAlertLogHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertLogHistogramResponse) SetStatusCode(v int32) *DescribeAlertLogHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertLogHistogramResponse) SetBody(v *DescribeAlertLogHistogramResponseBody) *DescribeAlertLogHistogramResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertLogHistogramResponse) Validate() error {
	return dara.Validate(s)
}

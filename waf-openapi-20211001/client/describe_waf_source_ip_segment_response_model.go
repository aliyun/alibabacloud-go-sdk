// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafSourceIpSegmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWafSourceIpSegmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWafSourceIpSegmentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWafSourceIpSegmentResponseBody) *DescribeWafSourceIpSegmentResponse
	GetBody() *DescribeWafSourceIpSegmentResponseBody
}

type DescribeWafSourceIpSegmentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWafSourceIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWafSourceIpSegmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWafSourceIpSegmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWafSourceIpSegmentResponse) GetBody() *DescribeWafSourceIpSegmentResponseBody {
	return s.Body
}

func (s *DescribeWafSourceIpSegmentResponse) SetHeaders(v map[string]*string) *DescribeWafSourceIpSegmentResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) SetStatusCode(v int32) *DescribeWafSourceIpSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) SetBody(v *DescribeWafSourceIpSegmentResponseBody) *DescribeWafSourceIpSegmentResponse {
	s.Body = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcMPUEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcMPUEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcMPUEventSubResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcMPUEventSubResponseBody) *DescribeRtcMPUEventSubResponse
	GetBody() *DescribeRtcMPUEventSubResponseBody
}

type DescribeRtcMPUEventSubResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcMPUEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcMPUEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcMPUEventSubResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcMPUEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcMPUEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcMPUEventSubResponse) GetBody() *DescribeRtcMPUEventSubResponseBody {
	return s.Body
}

func (s *DescribeRtcMPUEventSubResponse) SetHeaders(v map[string]*string) *DescribeRtcMPUEventSubResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcMPUEventSubResponse) SetStatusCode(v int32) *DescribeRtcMPUEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponse) SetBody(v *DescribeRtcMPUEventSubResponseBody) *DescribeRtcMPUEventSubResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcMPUEventSubResponse) Validate() error {
	return dara.Validate(s)
}

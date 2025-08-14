// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcRobotInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcRobotInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcRobotInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcRobotInstanceResponseBody) *DescribeRtcRobotInstanceResponse
	GetBody() *DescribeRtcRobotInstanceResponseBody
}

type DescribeRtcRobotInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcRobotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcRobotInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcRobotInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcRobotInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcRobotInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcRobotInstanceResponse) GetBody() *DescribeRtcRobotInstanceResponseBody {
	return s.Body
}

func (s *DescribeRtcRobotInstanceResponse) SetHeaders(v map[string]*string) *DescribeRtcRobotInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcRobotInstanceResponse) SetStatusCode(v int32) *DescribeRtcRobotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponse) SetBody(v *DescribeRtcRobotInstanceResponseBody) *DescribeRtcRobotInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcRobotInstanceResponse) Validate() error {
	return dara.Validate(s)
}

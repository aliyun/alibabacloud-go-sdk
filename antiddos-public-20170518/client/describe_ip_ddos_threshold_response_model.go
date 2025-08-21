// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpDdosThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpDdosThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpDdosThresholdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpDdosThresholdResponseBody) *DescribeIpDdosThresholdResponse
	GetBody() *DescribeIpDdosThresholdResponseBody
}

type DescribeIpDdosThresholdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpDdosThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpDdosThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpDdosThresholdResponse) GetBody() *DescribeIpDdosThresholdResponseBody {
	return s.Body
}

func (s *DescribeIpDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeIpDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpDdosThresholdResponse) SetStatusCode(v int32) *DescribeIpDdosThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpDdosThresholdResponse) SetBody(v *DescribeIpDdosThresholdResponseBody) *DescribeIpDdosThresholdResponse {
	s.Body = v
	return s
}

func (s *DescribeIpDdosThresholdResponse) Validate() error {
	return dara.Validate(s)
}

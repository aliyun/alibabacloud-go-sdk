// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpTrafficResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpTrafficResponseBody) *DescribeIpTrafficResponse
	GetBody() *DescribeIpTrafficResponseBody
}

type DescribeIpTrafficResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpTrafficResponse) GetBody() *DescribeIpTrafficResponseBody {
	return s.Body
}

func (s *DescribeIpTrafficResponse) SetHeaders(v map[string]*string) *DescribeIpTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpTrafficResponse) SetStatusCode(v int32) *DescribeIpTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetBody(v *DescribeIpTrafficResponseBody) *DescribeIpTrafficResponse {
	s.Body = v
	return s
}

func (s *DescribeIpTrafficResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpStatusResponseBody) *DescribeIpStatusResponse
	GetBody() *DescribeIpStatusResponseBody
}

type DescribeIpStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpStatusResponse) GetBody() *DescribeIpStatusResponseBody {
	return s.Body
}

func (s *DescribeIpStatusResponse) SetHeaders(v map[string]*string) *DescribeIpStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpStatusResponse) SetStatusCode(v int32) *DescribeIpStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpStatusResponse) SetBody(v *DescribeIpStatusResponseBody) *DescribeIpStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeIpStatusResponse) Validate() error {
	return dara.Validate(s)
}

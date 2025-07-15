// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByIpControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisByIpControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisByIpControlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisByIpControlResponseBody) *DescribeApisByIpControlResponse
	GetBody() *DescribeApisByIpControlResponseBody
}

type DescribeApisByIpControlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByIpControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByIpControlResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisByIpControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisByIpControlResponse) GetBody() *DescribeApisByIpControlResponseBody {
	return s.Body
}

func (s *DescribeApisByIpControlResponse) SetHeaders(v map[string]*string) *DescribeApisByIpControlResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByIpControlResponse) SetStatusCode(v int32) *DescribeApisByIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByIpControlResponse) SetBody(v *DescribeApisByIpControlResponseBody) *DescribeApisByIpControlResponse {
	s.Body = v
	return s
}

func (s *DescribeApisByIpControlResponse) Validate() error {
	return dara.Validate(s)
}

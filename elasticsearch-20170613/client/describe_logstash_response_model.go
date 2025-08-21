// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogstashResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogstashResponseBody) *DescribeLogstashResponse
	GetBody() *DescribeLogstashResponseBody
}

type DescribeLogstashResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogstashResponse) GetBody() *DescribeLogstashResponseBody {
	return s.Body
}

func (s *DescribeLogstashResponse) SetHeaders(v map[string]*string) *DescribeLogstashResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogstashResponse) SetStatusCode(v int32) *DescribeLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogstashResponse) SetBody(v *DescribeLogstashResponseBody) *DescribeLogstashResponse {
	s.Body = v
	return s
}

func (s *DescribeLogstashResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByTrafficControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisByTrafficControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisByTrafficControlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisByTrafficControlResponseBody) *DescribeApisByTrafficControlResponse
	GetBody() *DescribeApisByTrafficControlResponseBody
}

type DescribeApisByTrafficControlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByTrafficControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisByTrafficControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisByTrafficControlResponse) GetBody() *DescribeApisByTrafficControlResponseBody {
	return s.Body
}

func (s *DescribeApisByTrafficControlResponse) SetHeaders(v map[string]*string) *DescribeApisByTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByTrafficControlResponse) SetStatusCode(v int32) *DescribeApisByTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByTrafficControlResponse) SetBody(v *DescribeApisByTrafficControlResponseBody) *DescribeApisByTrafficControlResponse {
	s.Body = v
	return s
}

func (s *DescribeApisByTrafficControlResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiGroupResponseBody) *DescribeApiGroupResponse
	GetBody() *DescribeApiGroupResponseBody
}

type DescribeApiGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiGroupResponse) GetBody() *DescribeApiGroupResponseBody {
	return s.Body
}

func (s *DescribeApiGroupResponse) SetHeaders(v map[string]*string) *DescribeApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupResponse) SetStatusCode(v int32) *DescribeApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGroupResponse) SetBody(v *DescribeApiGroupResponseBody) *DescribeApiGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeApiGroupResponse) Validate() error {
	return dara.Validate(s)
}

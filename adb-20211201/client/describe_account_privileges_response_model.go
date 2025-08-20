// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountPrivilegesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountPrivilegesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountPrivilegesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountPrivilegesResponseBody) *DescribeAccountPrivilegesResponse
	GetBody() *DescribeAccountPrivilegesResponseBody
}

type DescribeAccountPrivilegesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountPrivilegesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountPrivilegesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountPrivilegesResponse) GetBody() *DescribeAccountPrivilegesResponseBody {
	return s.Body
}

func (s *DescribeAccountPrivilegesResponse) SetHeaders(v map[string]*string) *DescribeAccountPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountPrivilegesResponse) SetStatusCode(v int32) *DescribeAccountPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountPrivilegesResponse) SetBody(v *DescribeAccountPrivilegesResponseBody) *DescribeAccountPrivilegesResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountPrivilegesResponse) Validate() error {
	return dara.Validate(s)
}

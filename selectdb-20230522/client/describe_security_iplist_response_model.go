// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityIPListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityIPListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityIPListResponseBody) *DescribeSecurityIPListResponse
	GetBody() *DescribeSecurityIPListResponseBody
}

type DescribeSecurityIPListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityIPListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityIPListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityIPListResponse) GetBody() *DescribeSecurityIPListResponseBody {
	return s.Body
}

func (s *DescribeSecurityIPListResponse) SetHeaders(v map[string]*string) *DescribeSecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIPListResponse) SetStatusCode(v int32) *DescribeSecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIPListResponse) SetBody(v *DescribeSecurityIPListResponseBody) *DescribeSecurityIPListResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityIPListResponse) Validate() error {
	return dara.Validate(s)
}

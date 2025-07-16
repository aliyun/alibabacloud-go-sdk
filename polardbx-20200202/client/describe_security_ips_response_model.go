// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityIpsResponseBody) *DescribeSecurityIpsResponse
	GetBody() *DescribeSecurityIpsResponseBody
}

type DescribeSecurityIpsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityIpsResponse) GetBody() *DescribeSecurityIpsResponseBody {
	return s.Body
}

func (s *DescribeSecurityIpsResponse) SetHeaders(v map[string]*string) *DescribeSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIpsResponse) SetStatusCode(v int32) *DescribeSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIpsResponse) SetBody(v *DescribeSecurityIpsResponseBody) *DescribeSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityIpsResponse) Validate() error {
	return dara.Validate(s)
}

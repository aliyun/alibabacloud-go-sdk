// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnabledPrivilegesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnabledPrivilegesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnabledPrivilegesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnabledPrivilegesResponseBody) *DescribeEnabledPrivilegesResponse
	GetBody() *DescribeEnabledPrivilegesResponseBody
}

type DescribeEnabledPrivilegesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnabledPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnabledPrivilegesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnabledPrivilegesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnabledPrivilegesResponse) GetBody() *DescribeEnabledPrivilegesResponseBody {
	return s.Body
}

func (s *DescribeEnabledPrivilegesResponse) SetHeaders(v map[string]*string) *DescribeEnabledPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnabledPrivilegesResponse) SetStatusCode(v int32) *DescribeEnabledPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponse) SetBody(v *DescribeEnabledPrivilegesResponseBody) *DescribeEnabledPrivilegesResponse {
	s.Body = v
	return s
}

func (s *DescribeEnabledPrivilegesResponse) Validate() error {
	return dara.Validate(s)
}

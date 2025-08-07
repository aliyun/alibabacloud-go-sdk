// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalSecurityIPGroupResponseBody) *DescribeGlobalSecurityIPGroupResponse
	GetBody() *DescribeGlobalSecurityIPGroupResponseBody
}

type DescribeGlobalSecurityIPGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalSecurityIPGroupResponse) GetBody() *DescribeGlobalSecurityIPGroupResponseBody {
	return s.Body
}

func (s *DescribeGlobalSecurityIPGroupResponse) SetHeaders(v map[string]*string) *DescribeGlobalSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponse) SetStatusCode(v int32) *DescribeGlobalSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponse) SetBody(v *DescribeGlobalSecurityIPGroupResponseBody) *DescribeGlobalSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponse) Validate() error {
	return dara.Validate(s)
}

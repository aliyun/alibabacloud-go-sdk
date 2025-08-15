// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceLinkedRoleStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceLinkedRoleStatusResponseBody) *DescribeServiceLinkedRoleStatusResponse
	GetBody() *DescribeServiceLinkedRoleStatusResponseBody
}

type DescribeServiceLinkedRoleStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceLinkedRoleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceLinkedRoleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceLinkedRoleStatusResponse) GetBody() *DescribeServiceLinkedRoleStatusResponseBody {
	return s.Body
}

func (s *DescribeServiceLinkedRoleStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponse) SetStatusCode(v int32) *DescribeServiceLinkedRoleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponse) SetBody(v *DescribeServiceLinkedRoleStatusResponseBody) *DescribeServiceLinkedRoleStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponse) Validate() error {
	return dara.Validate(s)
}

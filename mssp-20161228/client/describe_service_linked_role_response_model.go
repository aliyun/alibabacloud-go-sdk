// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceLinkedRoleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceLinkedRoleResponseBody) *DescribeServiceLinkedRoleResponse
	GetBody() *DescribeServiceLinkedRoleResponseBody
}

type DescribeServiceLinkedRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceLinkedRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceLinkedRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceLinkedRoleResponse) GetBody() *DescribeServiceLinkedRoleResponseBody {
	return s.Body
}

func (s *DescribeServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLinkedRoleResponse) SetStatusCode(v int32) *DescribeServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLinkedRoleResponse) SetBody(v *DescribeServiceLinkedRoleResponseBody) *DescribeServiceLinkedRoleResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceLinkedRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

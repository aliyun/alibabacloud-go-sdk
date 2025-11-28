// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountRoleResponseBody) *DescribeAccountRoleResponse
	GetBody() *DescribeAccountRoleResponseBody
}

type DescribeAccountRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountRoleResponse) GetBody() *DescribeAccountRoleResponseBody {
	return s.Body
}

func (s *DescribeAccountRoleResponse) SetHeaders(v map[string]*string) *DescribeAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountRoleResponse) SetStatusCode(v int32) *DescribeAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountRoleResponse) SetBody(v *DescribeAccountRoleResponseBody) *DescribeAccountRoleResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

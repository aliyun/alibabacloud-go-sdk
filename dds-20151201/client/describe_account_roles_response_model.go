// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountRolesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountRolesResponseBody) *DescribeAccountRolesResponse
	GetBody() *DescribeAccountRolesResponseBody
}

type DescribeAccountRolesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRolesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountRolesResponse) GetBody() *DescribeAccountRolesResponseBody {
	return s.Body
}

func (s *DescribeAccountRolesResponse) SetHeaders(v map[string]*string) *DescribeAccountRolesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountRolesResponse) SetStatusCode(v int32) *DescribeAccountRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountRolesResponse) SetBody(v *DescribeAccountRolesResponseBody) *DescribeAccountRolesResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

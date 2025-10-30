// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRolesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRolesResponseBody) *DescribeRolesResponse
	GetBody() *DescribeRolesResponseBody
}

type DescribeRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRolesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRolesResponse) GetBody() *DescribeRolesResponseBody {
	return s.Body
}

func (s *DescribeRolesResponse) SetHeaders(v map[string]*string) *DescribeRolesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRolesResponse) SetStatusCode(v int32) *DescribeRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRolesResponse) SetBody(v *DescribeRolesResponseBody) *DescribeRolesResponse {
	s.Body = v
	return s
}

func (s *DescribeRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

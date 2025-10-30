// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleAuthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoleAuthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoleAuthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoleAuthStatusResponseBody) *DescribeRoleAuthStatusResponse
	GetBody() *DescribeRoleAuthStatusResponseBody
}

type DescribeRoleAuthStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoleAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoleAuthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoleAuthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoleAuthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoleAuthStatusResponse) GetBody() *DescribeRoleAuthStatusResponseBody {
	return s.Body
}

func (s *DescribeRoleAuthStatusResponse) SetHeaders(v map[string]*string) *DescribeRoleAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoleAuthStatusResponse) SetStatusCode(v int32) *DescribeRoleAuthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoleAuthStatusResponse) SetBody(v *DescribeRoleAuthStatusResponseBody) *DescribeRoleAuthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRoleAuthStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

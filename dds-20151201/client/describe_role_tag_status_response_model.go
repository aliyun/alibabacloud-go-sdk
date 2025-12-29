// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleTagStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoleTagStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoleTagStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoleTagStatusResponseBody) *DescribeRoleTagStatusResponse
	GetBody() *DescribeRoleTagStatusResponseBody
}

type DescribeRoleTagStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoleTagStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoleTagStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleTagStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoleTagStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoleTagStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoleTagStatusResponse) GetBody() *DescribeRoleTagStatusResponseBody {
	return s.Body
}

func (s *DescribeRoleTagStatusResponse) SetHeaders(v map[string]*string) *DescribeRoleTagStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoleTagStatusResponse) SetStatusCode(v int32) *DescribeRoleTagStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoleTagStatusResponse) SetBody(v *DescribeRoleTagStatusResponseBody) *DescribeRoleTagStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRoleTagStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceLinkedRoleExistsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceLinkedRoleExistsResponseBody) *DescribeServiceLinkedRoleExistsResponse
	GetBody() *DescribeServiceLinkedRoleExistsResponseBody
}

type DescribeServiceLinkedRoleExistsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceLinkedRoleExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceLinkedRoleExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleExistsResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceLinkedRoleExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceLinkedRoleExistsResponse) GetBody() *DescribeServiceLinkedRoleExistsResponseBody {
	return s.Body
}

func (s *DescribeServiceLinkedRoleExistsResponse) SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleExistsResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLinkedRoleExistsResponse) SetStatusCode(v int32) *DescribeServiceLinkedRoleExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsResponse) SetBody(v *DescribeServiceLinkedRoleExistsResponseBody) *DescribeServiceLinkedRoleExistsResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceLinkedRoleExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityIPGroupResponseBody) *DescribeSecurityIPGroupResponse
	GetBody() *DescribeSecurityIPGroupResponseBody
}

type DescribeSecurityIPGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityIPGroupResponse) GetBody() *DescribeSecurityIPGroupResponseBody {
	return s.Body
}

func (s *DescribeSecurityIPGroupResponse) SetHeaders(v map[string]*string) *DescribeSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIPGroupResponse) SetStatusCode(v int32) *DescribeSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIPGroupResponse) SetBody(v *DescribeSecurityIPGroupResponseBody) *DescribeSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityIPGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

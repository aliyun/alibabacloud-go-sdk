// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclResourceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclResourceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclResourceNameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclResourceNameResponseBody) *DescribeAclResourceNameResponse
	GetBody() *DescribeAclResourceNameResponseBody
}

type DescribeAclResourceNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclResourceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclResourceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclResourceNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclResourceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclResourceNameResponse) GetBody() *DescribeAclResourceNameResponseBody {
	return s.Body
}

func (s *DescribeAclResourceNameResponse) SetHeaders(v map[string]*string) *DescribeAclResourceNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclResourceNameResponse) SetStatusCode(v int32) *DescribeAclResourceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclResourceNameResponse) SetBody(v *DescribeAclResourceNameResponseBody) *DescribeAclResourceNameResponse {
	s.Body = v
	return s
}

func (s *DescribeAclResourceNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

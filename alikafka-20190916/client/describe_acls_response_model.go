// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclsResponseBody) *DescribeAclsResponse
	GetBody() *DescribeAclsResponseBody
}

type DescribeAclsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclsResponse) GetBody() *DescribeAclsResponseBody {
	return s.Body
}

func (s *DescribeAclsResponse) SetHeaders(v map[string]*string) *DescribeAclsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclsResponse) SetStatusCode(v int32) *DescribeAclsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclsResponse) SetBody(v *DescribeAclsResponseBody) *DescribeAclsResponse {
	s.Body = v
	return s
}

func (s *DescribeAclsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

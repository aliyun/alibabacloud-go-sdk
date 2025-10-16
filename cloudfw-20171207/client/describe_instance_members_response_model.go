// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceMembersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceMembersResponseBody) *DescribeInstanceMembersResponse
	GetBody() *DescribeInstanceMembersResponseBody
}

type DescribeInstanceMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceMembersResponse) GetBody() *DescribeInstanceMembersResponseBody {
	return s.Body
}

func (s *DescribeInstanceMembersResponse) SetHeaders(v map[string]*string) *DescribeInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMembersResponse) SetStatusCode(v int32) *DescribeInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMembersResponse) SetBody(v *DescribeInstanceMembersResponseBody) *DescribeInstanceMembersResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

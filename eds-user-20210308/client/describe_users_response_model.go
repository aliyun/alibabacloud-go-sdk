// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsersResponseBody) *DescribeUsersResponse
	GetBody() *DescribeUsersResponseBody
}

type DescribeUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsersResponse) GetBody() *DescribeUsersResponseBody {
	return s.Body
}

func (s *DescribeUsersResponse) SetHeaders(v map[string]*string) *DescribeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersResponse) SetStatusCode(v int32) *DescribeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersResponse) SetBody(v *DescribeUsersResponseBody) *DescribeUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

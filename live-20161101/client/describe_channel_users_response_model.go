// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelUsersResponseBody) *DescribeChannelUsersResponse
	GetBody() *DescribeChannelUsersResponseBody
}

type DescribeChannelUsersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelUsersResponse) GetBody() *DescribeChannelUsersResponseBody {
	return s.Body
}

func (s *DescribeChannelUsersResponse) SetHeaders(v map[string]*string) *DescribeChannelUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUsersResponse) SetStatusCode(v int32) *DescribeChannelUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUsersResponse) SetBody(v *DescribeChannelUsersResponseBody) *DescribeChannelUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelUsersResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAllUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelAllUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelAllUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelAllUsersResponseBody) *DescribeChannelAllUsersResponse
	GetBody() *DescribeChannelAllUsersResponseBody
}

type DescribeChannelAllUsersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelAllUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelAllUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAllUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelAllUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelAllUsersResponse) GetBody() *DescribeChannelAllUsersResponseBody {
	return s.Body
}

func (s *DescribeChannelAllUsersResponse) SetHeaders(v map[string]*string) *DescribeChannelAllUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelAllUsersResponse) SetStatusCode(v int32) *DescribeChannelAllUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelAllUsersResponse) SetBody(v *DescribeChannelAllUsersResponseBody) *DescribeChannelAllUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelAllUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

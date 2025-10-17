// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelUserResponseBody) *DescribeChannelUserResponse
	GetBody() *DescribeChannelUserResponseBody
}

type DescribeChannelUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelUserResponse) GetBody() *DescribeChannelUserResponseBody {
	return s.Body
}

func (s *DescribeChannelUserResponse) SetHeaders(v map[string]*string) *DescribeChannelUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUserResponse) SetStatusCode(v int32) *DescribeChannelUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUserResponse) SetBody(v *DescribeChannelUserResponseBody) *DescribeChannelUserResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

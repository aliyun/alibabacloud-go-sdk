// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserInfoInChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserInfoInChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserInfoInChannelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserInfoInChannelResponseBody) *DescribeUserInfoInChannelResponse
	GetBody() *DescribeUserInfoInChannelResponseBody
}

type DescribeUserInfoInChannelResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserInfoInChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserInfoInChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoInChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserInfoInChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserInfoInChannelResponse) GetBody() *DescribeUserInfoInChannelResponseBody {
	return s.Body
}

func (s *DescribeUserInfoInChannelResponse) SetHeaders(v map[string]*string) *DescribeUserInfoInChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserInfoInChannelResponse) SetStatusCode(v int32) *DescribeUserInfoInChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserInfoInChannelResponse) SetBody(v *DescribeUserInfoInChannelResponseBody) *DescribeUserInfoInChannelResponse {
	s.Body = v
	return s
}

func (s *DescribeUserInfoInChannelResponse) Validate() error {
	return dara.Validate(s)
}

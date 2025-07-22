// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelResponseBody) *DescribeChannelResponse
	GetBody() *DescribeChannelResponseBody
}

type DescribeChannelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelResponse) GetBody() *DescribeChannelResponseBody {
	return s.Body
}

func (s *DescribeChannelResponse) SetHeaders(v map[string]*string) *DescribeChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelResponse) SetStatusCode(v int32) *DescribeChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelResponse) SetBody(v *DescribeChannelResponseBody) *DescribeChannelResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelResponse) Validate() error {
	return dara.Validate(s)
}

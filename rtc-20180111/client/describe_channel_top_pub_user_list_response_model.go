// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelTopPubUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelTopPubUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelTopPubUserListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelTopPubUserListResponseBody) *DescribeChannelTopPubUserListResponse
	GetBody() *DescribeChannelTopPubUserListResponseBody
}

type DescribeChannelTopPubUserListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelTopPubUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelTopPubUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelTopPubUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelTopPubUserListResponse) GetBody() *DescribeChannelTopPubUserListResponseBody {
	return s.Body
}

func (s *DescribeChannelTopPubUserListResponse) SetHeaders(v map[string]*string) *DescribeChannelTopPubUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) SetStatusCode(v int32) *DescribeChannelTopPubUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) SetBody(v *DescribeChannelTopPubUserListResponseBody) *DescribeChannelTopPubUserListResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) Validate() error {
	return dara.Validate(s)
}

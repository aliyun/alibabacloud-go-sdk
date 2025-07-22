// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcChannelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcChannelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcChannelListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcChannelListResponseBody) *DescribeRtcChannelListResponse
	GetBody() *DescribeRtcChannelListResponseBody
}

type DescribeRtcChannelListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcChannelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcChannelListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcChannelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcChannelListResponse) GetBody() *DescribeRtcChannelListResponseBody {
	return s.Body
}

func (s *DescribeRtcChannelListResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelListResponse) SetStatusCode(v int32) *DescribeRtcChannelListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcChannelListResponse) SetBody(v *DescribeRtcChannelListResponseBody) *DescribeRtcChannelListResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcChannelListResponse) Validate() error {
	return dara.Validate(s)
}

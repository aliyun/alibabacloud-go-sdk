// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelUserMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelUserMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelUserMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelUserMetricsResponseBody) *DescribeChannelUserMetricsResponse
	GetBody() *DescribeChannelUserMetricsResponseBody
}

type DescribeChannelUserMetricsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUserMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelUserMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelUserMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelUserMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelUserMetricsResponse) GetBody() *DescribeChannelUserMetricsResponseBody {
	return s.Body
}

func (s *DescribeChannelUserMetricsResponse) SetHeaders(v map[string]*string) *DescribeChannelUserMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUserMetricsResponse) SetStatusCode(v int32) *DescribeChannelUserMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUserMetricsResponse) SetBody(v *DescribeChannelUserMetricsResponseBody) *DescribeChannelUserMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelUserMetricsResponse) Validate() error {
	return dara.Validate(s)
}

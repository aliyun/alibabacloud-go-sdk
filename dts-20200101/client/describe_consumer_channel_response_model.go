// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConsumerChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConsumerChannelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConsumerChannelResponseBody) *DescribeConsumerChannelResponse
	GetBody() *DescribeConsumerChannelResponseBody
}

type DescribeConsumerChannelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConsumerChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConsumerChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConsumerChannelResponse) GetBody() *DescribeConsumerChannelResponseBody {
	return s.Body
}

func (s *DescribeConsumerChannelResponse) SetHeaders(v map[string]*string) *DescribeConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsumerChannelResponse) SetStatusCode(v int32) *DescribeConsumerChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConsumerChannelResponse) SetBody(v *DescribeConsumerChannelResponseBody) *DescribeConsumerChannelResponse {
	s.Body = v
	return s
}

func (s *DescribeConsumerChannelResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConsumerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConsumerGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConsumerGroupResponseBody) *DescribeConsumerGroupResponse
	GetBody() *DescribeConsumerGroupResponseBody
}

type DescribeConsumerGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConsumerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConsumerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConsumerGroupResponse) GetBody() *DescribeConsumerGroupResponseBody {
	return s.Body
}

func (s *DescribeConsumerGroupResponse) SetHeaders(v map[string]*string) *DescribeConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsumerGroupResponse) SetStatusCode(v int32) *DescribeConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConsumerGroupResponse) SetBody(v *DescribeConsumerGroupResponseBody) *DescribeConsumerGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeConsumerGroupResponse) Validate() error {
	return dara.Validate(s)
}

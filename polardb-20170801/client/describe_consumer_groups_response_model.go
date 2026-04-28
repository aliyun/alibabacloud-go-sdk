// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConsumerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConsumerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConsumerGroupsResponseBody) *DescribeConsumerGroupsResponse
	GetBody() *DescribeConsumerGroupsResponseBody
}

type DescribeConsumerGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConsumerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConsumerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConsumerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConsumerGroupsResponse) GetBody() *DescribeConsumerGroupsResponseBody {
	return s.Body
}

func (s *DescribeConsumerGroupsResponse) SetHeaders(v map[string]*string) *DescribeConsumerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsumerGroupsResponse) SetStatusCode(v int32) *DescribeConsumerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConsumerGroupsResponse) SetBody(v *DescribeConsumerGroupsResponseBody) *DescribeConsumerGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeConsumerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

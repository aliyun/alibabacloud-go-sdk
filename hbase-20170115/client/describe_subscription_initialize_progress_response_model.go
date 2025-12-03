// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInitializeProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionInitializeProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionInitializeProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionInitializeProgressResponseBody) *DescribeSubscriptionInitializeProgressResponse
	GetBody() *DescribeSubscriptionInitializeProgressResponseBody
}

type DescribeSubscriptionInitializeProgressResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionInitializeProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionInitializeProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionInitializeProgressResponse) GetBody() *DescribeSubscriptionInitializeProgressResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionInitializeProgressResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInitializeProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponse) SetStatusCode(v int32) *DescribeSubscriptionInitializeProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponse) SetBody(v *DescribeSubscriptionInitializeProgressResponseBody) *DescribeSubscriptionInitializeProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

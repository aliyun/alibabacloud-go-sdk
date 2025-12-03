// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionPerformanceResponseBody) *DescribeSubscriptionPerformanceResponse
	GetBody() *DescribeSubscriptionPerformanceResponseBody
}

type DescribeSubscriptionPerformanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionPerformanceResponse) GetBody() *DescribeSubscriptionPerformanceResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionPerformanceResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponse) SetStatusCode(v int32) *DescribeSubscriptionPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponse) SetBody(v *DescribeSubscriptionPerformanceResponseBody) *DescribeSubscriptionPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

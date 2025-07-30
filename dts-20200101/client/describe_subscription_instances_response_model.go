// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionInstancesResponseBody) *DescribeSubscriptionInstancesResponse
	GetBody() *DescribeSubscriptionInstancesResponseBody
}

type DescribeSubscriptionInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionInstancesResponse) GetBody() *DescribeSubscriptionInstancesResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionInstancesResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInstancesResponse) SetStatusCode(v int32) *DescribeSubscriptionInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponse) SetBody(v *DescribeSubscriptionInstancesResponseBody) *DescribeSubscriptionInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionInstancesResponse) Validate() error {
	return dara.Validate(s)
}

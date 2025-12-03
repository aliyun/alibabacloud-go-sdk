// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionsResponseBody) *DescribeSubscriptionsResponse
	GetBody() *DescribeSubscriptionsResponseBody
}

type DescribeSubscriptionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionsResponse) GetBody() *DescribeSubscriptionsResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionsResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionsResponse) SetStatusCode(v int32) *DescribeSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionsResponse) SetBody(v *DescribeSubscriptionsResponseBody) *DescribeSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionInstanceStatusResponseBody) *DescribeSubscriptionInstanceStatusResponse
	GetBody() *DescribeSubscriptionInstanceStatusResponseBody
}

type DescribeSubscriptionInstanceStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionInstanceStatusResponse) GetBody() *DescribeSubscriptionInstanceStatusResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponse) SetStatusCode(v int32) *DescribeSubscriptionInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponse) SetBody(v *DescribeSubscriptionInstanceStatusResponseBody) *DescribeSubscriptionInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

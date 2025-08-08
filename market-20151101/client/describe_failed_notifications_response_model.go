// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailedNotificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFailedNotificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFailedNotificationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFailedNotificationsResponseBody) *DescribeFailedNotificationsResponse
	GetBody() *DescribeFailedNotificationsResponseBody
}

type DescribeFailedNotificationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFailedNotificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFailedNotificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailedNotificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFailedNotificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFailedNotificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFailedNotificationsResponse) GetBody() *DescribeFailedNotificationsResponseBody {
	return s.Body
}

func (s *DescribeFailedNotificationsResponse) SetHeaders(v map[string]*string) *DescribeFailedNotificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFailedNotificationsResponse) SetStatusCode(v int32) *DescribeFailedNotificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFailedNotificationsResponse) SetBody(v *DescribeFailedNotificationsResponseBody) *DescribeFailedNotificationsResponse {
	s.Body = v
	return s
}

func (s *DescribeFailedNotificationsResponse) Validate() error {
	return dara.Validate(s)
}

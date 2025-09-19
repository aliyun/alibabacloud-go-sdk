// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockFileEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockFileEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockFileEventsResponseBody) *DescribeWebLockFileEventsResponse
	GetBody() *DescribeWebLockFileEventsResponseBody
}

type DescribeWebLockFileEventsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockFileEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockFileEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockFileEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockFileEventsResponse) GetBody() *DescribeWebLockFileEventsResponseBody {
	return s.Body
}

func (s *DescribeWebLockFileEventsResponse) SetHeaders(v map[string]*string) *DescribeWebLockFileEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockFileEventsResponse) SetStatusCode(v int32) *DescribeWebLockFileEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockFileEventsResponse) SetBody(v *DescribeWebLockFileEventsResponseBody) *DescribeWebLockFileEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockFileEventsResponse) Validate() error {
	return dara.Validate(s)
}

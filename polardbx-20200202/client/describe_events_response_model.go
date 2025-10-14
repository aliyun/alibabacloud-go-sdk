// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse
	GetBody() *DescribeEventsResponseBody
}

type DescribeEventsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventsResponse) GetBody() *DescribeEventsResponseBody {
	return s.Body
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetStatusCode(v int32) *DescribeEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

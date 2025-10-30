// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterEventsResponseBody) *DescribeClusterEventsResponse
	GetBody() *DescribeClusterEventsResponseBody
}

type DescribeClusterEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterEventsResponse) GetBody() *DescribeClusterEventsResponseBody {
	return s.Body
}

func (s *DescribeClusterEventsResponse) SetHeaders(v map[string]*string) *DescribeClusterEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterEventsResponse) SetStatusCode(v int32) *DescribeClusterEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterEventsResponse) SetBody(v *DescribeClusterEventsResponseBody) *DescribeClusterEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

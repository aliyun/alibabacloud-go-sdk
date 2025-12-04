// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLifeCycleEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceLifeCycleEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceLifeCycleEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceLifeCycleEventsResponseBody) *DescribeResourceLifeCycleEventsResponse
	GetBody() *DescribeResourceLifeCycleEventsResponseBody
}

type DescribeResourceLifeCycleEventsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLifeCycleEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLifeCycleEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLifeCycleEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLifeCycleEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceLifeCycleEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceLifeCycleEventsResponse) GetBody() *DescribeResourceLifeCycleEventsResponseBody {
	return s.Body
}

func (s *DescribeResourceLifeCycleEventsResponse) SetHeaders(v map[string]*string) *DescribeResourceLifeCycleEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLifeCycleEventsResponse) SetStatusCode(v int32) *DescribeResourceLifeCycleEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLifeCycleEventsResponse) SetBody(v *DescribeResourceLifeCycleEventsResponseBody) *DescribeResourceLifeCycleEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceLifeCycleEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

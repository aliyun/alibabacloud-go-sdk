// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientEventsResponseBody) *DescribeClientEventsResponse
	GetBody() *DescribeClientEventsResponseBody
}

type DescribeClientEventsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientEventsResponse) GetBody() *DescribeClientEventsResponseBody {
	return s.Body
}

func (s *DescribeClientEventsResponse) SetHeaders(v map[string]*string) *DescribeClientEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientEventsResponse) SetStatusCode(v int32) *DescribeClientEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientEventsResponse) SetBody(v *DescribeClientEventsResponseBody) *DescribeClientEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeClientEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventsListResponseBody) *DescribeEventsListResponse
	GetBody() *DescribeEventsListResponseBody
}

type DescribeEventsListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventsListResponse) GetBody() *DescribeEventsListResponseBody {
	return s.Body
}

func (s *DescribeEventsListResponse) SetHeaders(v map[string]*string) *DescribeEventsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsListResponse) SetStatusCode(v int32) *DescribeEventsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsListResponse) SetBody(v *DescribeEventsListResponseBody) *DescribeEventsListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventsListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

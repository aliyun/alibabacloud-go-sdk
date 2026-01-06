// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHistoryEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHistoryEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHistoryEventsResponseBody) *DescribeHistoryEventsResponse
	GetBody() *DescribeHistoryEventsResponseBody
}

type DescribeHistoryEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHistoryEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHistoryEventsResponse) GetBody() *DescribeHistoryEventsResponseBody {
	return s.Body
}

func (s *DescribeHistoryEventsResponse) SetHeaders(v map[string]*string) *DescribeHistoryEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryEventsResponse) SetStatusCode(v int32) *DescribeHistoryEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryEventsResponse) SetBody(v *DescribeHistoryEventsResponseBody) *DescribeHistoryEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeHistoryEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

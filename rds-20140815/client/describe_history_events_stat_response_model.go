// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHistoryEventsStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHistoryEventsStatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHistoryEventsStatResponseBody) *DescribeHistoryEventsStatResponse
	GetBody() *DescribeHistoryEventsStatResponseBody
}

type DescribeHistoryEventsStatResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryEventsStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryEventsStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHistoryEventsStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHistoryEventsStatResponse) GetBody() *DescribeHistoryEventsStatResponseBody {
	return s.Body
}

func (s *DescribeHistoryEventsStatResponse) SetHeaders(v map[string]*string) *DescribeHistoryEventsStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryEventsStatResponse) SetStatusCode(v int32) *DescribeHistoryEventsStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryEventsStatResponse) SetBody(v *DescribeHistoryEventsStatResponseBody) *DescribeHistoryEventsStatResponse {
	s.Body = v
	return s
}

func (s *DescribeHistoryEventsStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

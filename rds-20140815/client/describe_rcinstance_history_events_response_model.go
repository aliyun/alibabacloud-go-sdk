// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceHistoryEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceHistoryEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceHistoryEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceHistoryEventsResponseBody) *DescribeRCInstanceHistoryEventsResponse
	GetBody() *DescribeRCInstanceHistoryEventsResponseBody
}

type DescribeRCInstanceHistoryEventsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceHistoryEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceHistoryEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceHistoryEventsResponse) GetBody() *DescribeRCInstanceHistoryEventsResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceHistoryEventsResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceHistoryEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponse) SetStatusCode(v int32) *DescribeRCInstanceHistoryEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponse) SetBody(v *DescribeRCInstanceHistoryEventsResponseBody) *DescribeRCInstanceHistoryEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

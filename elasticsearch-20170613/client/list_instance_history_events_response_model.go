// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceHistoryEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceHistoryEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceHistoryEventsResponseBody) *ListInstanceHistoryEventsResponse
	GetBody() *ListInstanceHistoryEventsResponseBody
}

type ListInstanceHistoryEventsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceHistoryEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceHistoryEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryEventsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceHistoryEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceHistoryEventsResponse) GetBody() *ListInstanceHistoryEventsResponseBody {
	return s.Body
}

func (s *ListInstanceHistoryEventsResponse) SetHeaders(v map[string]*string) *ListInstanceHistoryEventsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceHistoryEventsResponse) SetStatusCode(v int32) *ListInstanceHistoryEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceHistoryEventsResponse) SetBody(v *ListInstanceHistoryEventsResponseBody) *ListInstanceHistoryEventsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceHistoryEventsResponse) Validate() error {
	return dara.Validate(s)
}

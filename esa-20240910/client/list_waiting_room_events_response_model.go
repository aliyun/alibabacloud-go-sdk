// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWaitingRoomEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWaitingRoomEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListWaitingRoomEventsResponseBody) *ListWaitingRoomEventsResponse
	GetBody() *ListWaitingRoomEventsResponseBody
}

type ListWaitingRoomEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingRoomEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingRoomEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomEventsResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWaitingRoomEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWaitingRoomEventsResponse) GetBody() *ListWaitingRoomEventsResponseBody {
	return s.Body
}

func (s *ListWaitingRoomEventsResponse) SetHeaders(v map[string]*string) *ListWaitingRoomEventsResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingRoomEventsResponse) SetStatusCode(v int32) *ListWaitingRoomEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingRoomEventsResponse) SetBody(v *ListWaitingRoomEventsResponseBody) *ListWaitingRoomEventsResponse {
	s.Body = v
	return s
}

func (s *ListWaitingRoomEventsResponse) Validate() error {
	return dara.Validate(s)
}

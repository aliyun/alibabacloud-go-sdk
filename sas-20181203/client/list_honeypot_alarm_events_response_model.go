// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAlarmEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotAlarmEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotAlarmEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotAlarmEventsResponseBody) *ListHoneypotAlarmEventsResponse
	GetBody() *ListHoneypotAlarmEventsResponseBody
}

type ListHoneypotAlarmEventsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotAlarmEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotAlarmEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAlarmEventsResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotAlarmEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotAlarmEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotAlarmEventsResponse) GetBody() *ListHoneypotAlarmEventsResponseBody {
	return s.Body
}

func (s *ListHoneypotAlarmEventsResponse) SetHeaders(v map[string]*string) *ListHoneypotAlarmEventsResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotAlarmEventsResponse) SetStatusCode(v int32) *ListHoneypotAlarmEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponse) SetBody(v *ListHoneypotAlarmEventsResponseBody) *ListHoneypotAlarmEventsResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotAlarmEventsResponse) Validate() error {
	return dara.Validate(s)
}

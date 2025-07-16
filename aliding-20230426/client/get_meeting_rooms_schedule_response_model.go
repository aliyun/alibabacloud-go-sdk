// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeetingRoomsScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMeetingRoomsScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMeetingRoomsScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetMeetingRoomsScheduleResponseBody) *GetMeetingRoomsScheduleResponse
	GetBody() *GetMeetingRoomsScheduleResponseBody
}

type GetMeetingRoomsScheduleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMeetingRoomsScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMeetingRoomsScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMeetingRoomsScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMeetingRoomsScheduleResponse) GetBody() *GetMeetingRoomsScheduleResponseBody {
	return s.Body
}

func (s *GetMeetingRoomsScheduleResponse) SetHeaders(v map[string]*string) *GetMeetingRoomsScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetMeetingRoomsScheduleResponse) SetStatusCode(v int32) *GetMeetingRoomsScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponse) SetBody(v *GetMeetingRoomsScheduleResponseBody) *GetMeetingRoomsScheduleResponse {
	s.Body = v
	return s
}

func (s *GetMeetingRoomsScheduleResponse) Validate() error {
	return dara.Validate(s)
}

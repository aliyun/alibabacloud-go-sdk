// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMeetingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMeetingRoomResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMeetingRoomResponseBody) *UpdateMeetingRoomResponse
	GetBody() *UpdateMeetingRoomResponseBody
}

type UpdateMeetingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMeetingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMeetingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMeetingRoomResponse) GetBody() *UpdateMeetingRoomResponseBody {
	return s.Body
}

func (s *UpdateMeetingRoomResponse) SetHeaders(v map[string]*string) *UpdateMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeetingRoomResponse) SetStatusCode(v int32) *UpdateMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeetingRoomResponse) SetBody(v *UpdateMeetingRoomResponseBody) *UpdateMeetingRoomResponse {
	s.Body = v
	return s
}

func (s *UpdateMeetingRoomResponse) Validate() error {
	return dara.Validate(s)
}

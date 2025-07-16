// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMeetingRoomGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMeetingRoomGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMeetingRoomGroupResponseBody) *UpdateMeetingRoomGroupResponse
	GetBody() *UpdateMeetingRoomGroupResponseBody
}

type UpdateMeetingRoomGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMeetingRoomGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMeetingRoomGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMeetingRoomGroupResponse) GetBody() *UpdateMeetingRoomGroupResponseBody {
	return s.Body
}

func (s *UpdateMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *UpdateMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeetingRoomGroupResponse) SetStatusCode(v int32) *UpdateMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeetingRoomGroupResponse) SetBody(v *UpdateMeetingRoomGroupResponseBody) *UpdateMeetingRoomGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateMeetingRoomGroupResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMeetingRoomGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMeetingRoomGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMeetingRoomGroupResponseBody) *DeleteMeetingRoomGroupResponse
	GetBody() *DeleteMeetingRoomGroupResponseBody
}

type DeleteMeetingRoomGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMeetingRoomGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMeetingRoomGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMeetingRoomGroupResponse) GetBody() *DeleteMeetingRoomGroupResponseBody {
	return s.Body
}

func (s *DeleteMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *DeleteMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeetingRoomGroupResponse) SetStatusCode(v int32) *DeleteMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeetingRoomGroupResponse) SetBody(v *DeleteMeetingRoomGroupResponseBody) *DeleteMeetingRoomGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteMeetingRoomGroupResponse) Validate() error {
	return dara.Validate(s)
}

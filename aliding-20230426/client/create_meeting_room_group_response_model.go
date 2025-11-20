// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMeetingRoomGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMeetingRoomGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMeetingRoomGroupResponseBody) *CreateMeetingRoomGroupResponse
	GetBody() *CreateMeetingRoomGroupResponseBody
}

type CreateMeetingRoomGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMeetingRoomGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMeetingRoomGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMeetingRoomGroupResponse) GetBody() *CreateMeetingRoomGroupResponseBody {
	return s.Body
}

func (s *CreateMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *CreateMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingRoomGroupResponse) SetStatusCode(v int32) *CreateMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingRoomGroupResponse) SetBody(v *CreateMeetingRoomGroupResponseBody) *CreateMeetingRoomGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMeetingRoomGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

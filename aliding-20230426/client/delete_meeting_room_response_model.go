// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMeetingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMeetingRoomResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMeetingRoomResponseBody) *DeleteMeetingRoomResponse
	GetBody() *DeleteMeetingRoomResponseBody
}

type DeleteMeetingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMeetingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMeetingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMeetingRoomResponse) GetBody() *DeleteMeetingRoomResponseBody {
	return s.Body
}

func (s *DeleteMeetingRoomResponse) SetHeaders(v map[string]*string) *DeleteMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeetingRoomResponse) SetStatusCode(v int32) *DeleteMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeetingRoomResponse) SetBody(v *DeleteMeetingRoomResponseBody) *DeleteMeetingRoomResponse {
	s.Body = v
	return s
}

func (s *DeleteMeetingRoomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

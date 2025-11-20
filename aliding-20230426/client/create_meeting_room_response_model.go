// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMeetingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMeetingRoomResponse
	GetStatusCode() *int32
	SetBody(v *CreateMeetingRoomResponseBody) *CreateMeetingRoomResponse
	GetBody() *CreateMeetingRoomResponseBody
}

type CreateMeetingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMeetingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMeetingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMeetingRoomResponse) GetBody() *CreateMeetingRoomResponseBody {
	return s.Body
}

func (s *CreateMeetingRoomResponse) SetHeaders(v map[string]*string) *CreateMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingRoomResponse) SetStatusCode(v int32) *CreateMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingRoomResponse) SetBody(v *CreateMeetingRoomResponseBody) *CreateMeetingRoomResponse {
	s.Body = v
	return s
}

func (s *CreateMeetingRoomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

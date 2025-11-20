// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMeetingRoomsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMeetingRoomsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMeetingRoomsResponse
	GetStatusCode() *int32
	SetBody(v *AddMeetingRoomsResponseBody) *AddMeetingRoomsResponse
	GetBody() *AddMeetingRoomsResponseBody
}

type AddMeetingRoomsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMeetingRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMeetingRoomsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsResponse) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMeetingRoomsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMeetingRoomsResponse) GetBody() *AddMeetingRoomsResponseBody {
	return s.Body
}

func (s *AddMeetingRoomsResponse) SetHeaders(v map[string]*string) *AddMeetingRoomsResponse {
	s.Headers = v
	return s
}

func (s *AddMeetingRoomsResponse) SetStatusCode(v int32) *AddMeetingRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMeetingRoomsResponse) SetBody(v *AddMeetingRoomsResponseBody) *AddMeetingRoomsResponse {
	s.Body = v
	return s
}

func (s *AddMeetingRoomsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

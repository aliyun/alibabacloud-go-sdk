// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MeetingFlashMinutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MeetingFlashMinutesResponse
	GetStatusCode() *int32
	SetBody(v *MeetingFlashMinutesResponseBody) *MeetingFlashMinutesResponse
	GetBody() *MeetingFlashMinutesResponseBody
}

type MeetingFlashMinutesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MeetingFlashMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MeetingFlashMinutesResponse) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesResponse) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MeetingFlashMinutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MeetingFlashMinutesResponse) GetBody() *MeetingFlashMinutesResponseBody {
	return s.Body
}

func (s *MeetingFlashMinutesResponse) SetHeaders(v map[string]*string) *MeetingFlashMinutesResponse {
	s.Headers = v
	return s
}

func (s *MeetingFlashMinutesResponse) SetStatusCode(v int32) *MeetingFlashMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *MeetingFlashMinutesResponse) SetBody(v *MeetingFlashMinutesResponseBody) *MeetingFlashMinutesResponse {
	s.Body = v
	return s
}

func (s *MeetingFlashMinutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

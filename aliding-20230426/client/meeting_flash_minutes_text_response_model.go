// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MeetingFlashMinutesTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MeetingFlashMinutesTextResponse
	GetStatusCode() *int32
	SetBody(v *MeetingFlashMinutesTextResponseBody) *MeetingFlashMinutesTextResponse
	GetBody() *MeetingFlashMinutesTextResponseBody
}

type MeetingFlashMinutesTextResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MeetingFlashMinutesTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MeetingFlashMinutesTextResponse) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextResponse) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MeetingFlashMinutesTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MeetingFlashMinutesTextResponse) GetBody() *MeetingFlashMinutesTextResponseBody {
	return s.Body
}

func (s *MeetingFlashMinutesTextResponse) SetHeaders(v map[string]*string) *MeetingFlashMinutesTextResponse {
	s.Headers = v
	return s
}

func (s *MeetingFlashMinutesTextResponse) SetStatusCode(v int32) *MeetingFlashMinutesTextResponse {
	s.StatusCode = &v
	return s
}

func (s *MeetingFlashMinutesTextResponse) SetBody(v *MeetingFlashMinutesTextResponseBody) *MeetingFlashMinutesTextResponse {
	s.Body = v
	return s
}

func (s *MeetingFlashMinutesTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

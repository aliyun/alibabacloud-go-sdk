// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduleConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduleConferenceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduleConferenceResponseBody) *UpdateScheduleConferenceResponse
	GetBody() *UpdateScheduleConferenceResponseBody
}

type UpdateScheduleConferenceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduleConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduleConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduleConferenceResponse) GetBody() *UpdateScheduleConferenceResponseBody {
	return s.Body
}

func (s *UpdateScheduleConferenceResponse) SetHeaders(v map[string]*string) *UpdateScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleConferenceResponse) SetStatusCode(v int32) *UpdateScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduleConferenceResponse) SetBody(v *UpdateScheduleConferenceResponseBody) *UpdateScheduleConferenceResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduleConferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

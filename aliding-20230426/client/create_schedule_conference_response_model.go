// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduleConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduleConferenceResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduleConferenceResponseBody) *CreateScheduleConferenceResponse
	GetBody() *CreateScheduleConferenceResponseBody
}

type CreateScheduleConferenceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduleConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduleConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduleConferenceResponse) GetBody() *CreateScheduleConferenceResponseBody {
	return s.Body
}

func (s *CreateScheduleConferenceResponse) SetHeaders(v map[string]*string) *CreateScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleConferenceResponse) SetStatusCode(v int32) *CreateScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleConferenceResponse) SetBody(v *CreateScheduleConferenceResponseBody) *CreateScheduleConferenceResponse {
	s.Body = v
	return s
}

func (s *CreateScheduleConferenceResponse) Validate() error {
	return dara.Validate(s)
}

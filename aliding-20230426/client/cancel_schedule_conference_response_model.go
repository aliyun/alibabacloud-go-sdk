// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelScheduleConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelScheduleConferenceResponse
	GetStatusCode() *int32
	SetBody(v *CancelScheduleConferenceResponseBody) *CancelScheduleConferenceResponse
	GetBody() *CancelScheduleConferenceResponseBody
}

type CancelScheduleConferenceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelScheduleConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelScheduleConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelScheduleConferenceResponse) GetBody() *CancelScheduleConferenceResponseBody {
	return s.Body
}

func (s *CancelScheduleConferenceResponse) SetHeaders(v map[string]*string) *CancelScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *CancelScheduleConferenceResponse) SetStatusCode(v int32) *CancelScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelScheduleConferenceResponse) SetBody(v *CancelScheduleConferenceResponseBody) *CancelScheduleConferenceResponse {
	s.Body = v
	return s
}

func (s *CancelScheduleConferenceResponse) Validate() error {
	return dara.Validate(s)
}

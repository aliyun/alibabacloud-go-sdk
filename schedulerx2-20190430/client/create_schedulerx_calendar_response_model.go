// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerxCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSchedulerxCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSchedulerxCalendarResponse
	GetStatusCode() *int32
	SetBody(v *CreateSchedulerxCalendarResponseBody) *CreateSchedulerxCalendarResponse
	GetBody() *CreateSchedulerxCalendarResponseBody
}

type CreateSchedulerxCalendarResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchedulerxCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchedulerxCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxCalendarResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSchedulerxCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSchedulerxCalendarResponse) GetBody() *CreateSchedulerxCalendarResponseBody {
	return s.Body
}

func (s *CreateSchedulerxCalendarResponse) SetHeaders(v map[string]*string) *CreateSchedulerxCalendarResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedulerxCalendarResponse) SetStatusCode(v int32) *CreateSchedulerxCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchedulerxCalendarResponse) SetBody(v *CreateSchedulerxCalendarResponseBody) *CreateSchedulerxCalendarResponse {
	s.Body = v
	return s
}

func (s *CreateSchedulerxCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

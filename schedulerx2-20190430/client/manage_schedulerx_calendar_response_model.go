// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageSchedulerxCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageSchedulerxCalendarResponse
	GetStatusCode() *int32
	SetBody(v *ManageSchedulerxCalendarResponseBody) *ManageSchedulerxCalendarResponse
	GetBody() *ManageSchedulerxCalendarResponseBody
}

type ManageSchedulerxCalendarResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageSchedulerxCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSchedulerxCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxCalendarResponse) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageSchedulerxCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageSchedulerxCalendarResponse) GetBody() *ManageSchedulerxCalendarResponseBody {
	return s.Body
}

func (s *ManageSchedulerxCalendarResponse) SetHeaders(v map[string]*string) *ManageSchedulerxCalendarResponse {
	s.Headers = v
	return s
}

func (s *ManageSchedulerxCalendarResponse) SetStatusCode(v int32) *ManageSchedulerxCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageSchedulerxCalendarResponse) SetBody(v *ManageSchedulerxCalendarResponseBody) *ManageSchedulerxCalendarResponse {
	s.Body = v
	return s
}

func (s *ManageSchedulerxCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

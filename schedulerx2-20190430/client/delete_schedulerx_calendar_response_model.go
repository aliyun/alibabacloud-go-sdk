// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerxCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchedulerxCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchedulerxCalendarResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchedulerxCalendarResponseBody) *DeleteSchedulerxCalendarResponse
	GetBody() *DeleteSchedulerxCalendarResponseBody
}

type DeleteSchedulerxCalendarResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchedulerxCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchedulerxCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxCalendarResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchedulerxCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchedulerxCalendarResponse) GetBody() *DeleteSchedulerxCalendarResponseBody {
	return s.Body
}

func (s *DeleteSchedulerxCalendarResponse) SetHeaders(v map[string]*string) *DeleteSchedulerxCalendarResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedulerxCalendarResponse) SetStatusCode(v int32) *DeleteSchedulerxCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponse) SetBody(v *DeleteSchedulerxCalendarResponseBody) *DeleteSchedulerxCalendarResponse {
	s.Body = v
	return s
}

func (s *DeleteSchedulerxCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

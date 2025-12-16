// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadSchedulerxCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadSchedulerxCalendarResponse
	GetStatusCode() *int32
	SetBody(v *ReadSchedulerxCalendarResponseBody) *ReadSchedulerxCalendarResponse
	GetBody() *ReadSchedulerxCalendarResponseBody
}

type ReadSchedulerxCalendarResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadSchedulerxCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadSchedulerxCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxCalendarResponse) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadSchedulerxCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadSchedulerxCalendarResponse) GetBody() *ReadSchedulerxCalendarResponseBody {
	return s.Body
}

func (s *ReadSchedulerxCalendarResponse) SetHeaders(v map[string]*string) *ReadSchedulerxCalendarResponse {
	s.Headers = v
	return s
}

func (s *ReadSchedulerxCalendarResponse) SetStatusCode(v int32) *ReadSchedulerxCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadSchedulerxCalendarResponse) SetBody(v *ReadSchedulerxCalendarResponseBody) *ReadSchedulerxCalendarResponse {
	s.Body = v
	return s
}

func (s *ReadSchedulerxCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

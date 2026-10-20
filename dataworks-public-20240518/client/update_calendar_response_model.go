// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCalendarResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCalendarResponseBody) *UpdateCalendarResponse
	GetBody() *UpdateCalendarResponseBody
}

type UpdateCalendarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarResponse) GoString() string {
	return s.String()
}

func (s *UpdateCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCalendarResponse) GetBody() *UpdateCalendarResponseBody {
	return s.Body
}

func (s *UpdateCalendarResponse) SetHeaders(v map[string]*string) *UpdateCalendarResponse {
	s.Headers = v
	return s
}

func (s *UpdateCalendarResponse) SetStatusCode(v int32) *UpdateCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCalendarResponse) SetBody(v *UpdateCalendarResponseBody) *UpdateCalendarResponse {
	s.Body = v
	return s
}

func (s *UpdateCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

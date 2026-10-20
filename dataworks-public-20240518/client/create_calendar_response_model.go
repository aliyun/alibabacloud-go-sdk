// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCalendarResponse
	GetStatusCode() *int32
	SetBody(v *CreateCalendarResponseBody) *CreateCalendarResponse
	GetBody() *CreateCalendarResponseBody
}

type CreateCalendarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarResponse) GoString() string {
	return s.String()
}

func (s *CreateCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCalendarResponse) GetBody() *CreateCalendarResponseBody {
	return s.Body
}

func (s *CreateCalendarResponse) SetHeaders(v map[string]*string) *CreateCalendarResponse {
	s.Headers = v
	return s
}

func (s *CreateCalendarResponse) SetStatusCode(v int32) *CreateCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCalendarResponse) SetBody(v *CreateCalendarResponseBody) *CreateCalendarResponse {
	s.Body = v
	return s
}

func (s *CreateCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

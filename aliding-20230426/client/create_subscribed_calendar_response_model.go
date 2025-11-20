// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscribedCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubscribedCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubscribedCalendarResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubscribedCalendarResponseBody) *CreateSubscribedCalendarResponse
	GetBody() *CreateSubscribedCalendarResponseBody
}

type CreateSubscribedCalendarResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubscribedCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubscribedCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubscribedCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubscribedCalendarResponse) GetBody() *CreateSubscribedCalendarResponseBody {
	return s.Body
}

func (s *CreateSubscribedCalendarResponse) SetHeaders(v map[string]*string) *CreateSubscribedCalendarResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscribedCalendarResponse) SetStatusCode(v int32) *CreateSubscribedCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscribedCalendarResponse) SetBody(v *CreateSubscribedCalendarResponseBody) *CreateSubscribedCalendarResponse {
	s.Body = v
	return s
}

func (s *CreateSubscribedCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

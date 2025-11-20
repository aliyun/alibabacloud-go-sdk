// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscribedCalendarsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSubscribedCalendarsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSubscribedCalendarsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSubscribedCalendarsResponseBody) *UpdateSubscribedCalendarsResponse
	GetBody() *UpdateSubscribedCalendarsResponseBody
}

type UpdateSubscribedCalendarsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubscribedCalendarsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubscribedCalendarsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSubscribedCalendarsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSubscribedCalendarsResponse) GetBody() *UpdateSubscribedCalendarsResponseBody {
	return s.Body
}

func (s *UpdateSubscribedCalendarsResponse) SetHeaders(v map[string]*string) *UpdateSubscribedCalendarsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscribedCalendarsResponse) SetStatusCode(v int32) *UpdateSubscribedCalendarsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubscribedCalendarsResponse) SetBody(v *UpdateSubscribedCalendarsResponseBody) *UpdateSubscribedCalendarsResponse {
	s.Body = v
	return s
}

func (s *UpdateSubscribedCalendarsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

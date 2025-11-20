// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCalendarsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCalendarsResponse
	GetStatusCode() *int32
	SetBody(v *ListCalendarsResponseBody) *ListCalendarsResponse
	GetBody() *ListCalendarsResponseBody
}

type ListCalendarsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalendarsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalendarsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCalendarsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCalendarsResponse) GetBody() *ListCalendarsResponseBody {
	return s.Body
}

func (s *ListCalendarsResponse) SetHeaders(v map[string]*string) *ListCalendarsResponse {
	s.Headers = v
	return s
}

func (s *ListCalendarsResponse) SetStatusCode(v int32) *ListCalendarsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalendarsResponse) SetBody(v *ListCalendarsResponseBody) *ListCalendarsResponse {
	s.Body = v
	return s
}

func (s *ListCalendarsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAttendeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAttendeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAttendeeResponse
	GetStatusCode() *int32
	SetBody(v *AddAttendeeResponseBody) *AddAttendeeResponse
	GetBody() *AddAttendeeResponseBody
}

type AddAttendeeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAttendeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAttendeeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeResponse) GoString() string {
	return s.String()
}

func (s *AddAttendeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAttendeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAttendeeResponse) GetBody() *AddAttendeeResponseBody {
	return s.Body
}

func (s *AddAttendeeResponse) SetHeaders(v map[string]*string) *AddAttendeeResponse {
	s.Headers = v
	return s
}

func (s *AddAttendeeResponse) SetStatusCode(v int32) *AddAttendeeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAttendeeResponse) SetBody(v *AddAttendeeResponseBody) *AddAttendeeResponse {
	s.Body = v
	return s
}

func (s *AddAttendeeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAttendeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAttendeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAttendeeResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAttendeeResponseBody) *RemoveAttendeeResponse
	GetBody() *RemoveAttendeeResponseBody
}

type RemoveAttendeeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAttendeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAttendeeResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeResponse) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAttendeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAttendeeResponse) GetBody() *RemoveAttendeeResponseBody {
	return s.Body
}

func (s *RemoveAttendeeResponse) SetHeaders(v map[string]*string) *RemoveAttendeeResponse {
	s.Headers = v
	return s
}

func (s *RemoveAttendeeResponse) SetStatusCode(v int32) *RemoveAttendeeResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAttendeeResponse) SetBody(v *RemoveAttendeeResponseBody) *RemoveAttendeeResponse {
	s.Body = v
	return s
}

func (s *RemoveAttendeeResponse) Validate() error {
	return dara.Validate(s)
}

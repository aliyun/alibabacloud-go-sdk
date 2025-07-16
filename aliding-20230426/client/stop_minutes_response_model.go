// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMinutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMinutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMinutesResponse
	GetStatusCode() *int32
	SetBody(v *StopMinutesResponseBody) *StopMinutesResponse
	GetBody() *StopMinutesResponseBody
}

type StopMinutesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMinutesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesResponse) GoString() string {
	return s.String()
}

func (s *StopMinutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMinutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMinutesResponse) GetBody() *StopMinutesResponseBody {
	return s.Body
}

func (s *StopMinutesResponse) SetHeaders(v map[string]*string) *StopMinutesResponse {
	s.Headers = v
	return s
}

func (s *StopMinutesResponse) SetStatusCode(v int32) *StopMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMinutesResponse) SetBody(v *StopMinutesResponseBody) *StopMinutesResponse {
	s.Body = v
	return s
}

func (s *StopMinutesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *StopHoneypotResponseBody) *StopHoneypotResponse
	GetBody() *StopHoneypotResponseBody
}

type StopHoneypotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s StopHoneypotResponse) GoString() string {
	return s.String()
}

func (s *StopHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopHoneypotResponse) GetBody() *StopHoneypotResponseBody {
	return s.Body
}

func (s *StopHoneypotResponse) SetHeaders(v map[string]*string) *StopHoneypotResponse {
	s.Headers = v
	return s
}

func (s *StopHoneypotResponse) SetStatusCode(v int32) *StopHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *StopHoneypotResponse) SetBody(v *StopHoneypotResponseBody) *StopHoneypotResponse {
	s.Body = v
	return s
}

func (s *StopHoneypotResponse) Validate() error {
	return dara.Validate(s)
}

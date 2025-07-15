// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCasterResponse
	GetStatusCode() *int32
	SetBody(v *StopCasterResponseBody) *StopCasterResponse
	GetBody() *StopCasterResponseBody
}

type StopCasterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCasterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCasterResponse) GoString() string {
	return s.String()
}

func (s *StopCasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCasterResponse) GetBody() *StopCasterResponseBody {
	return s.Body
}

func (s *StopCasterResponse) SetHeaders(v map[string]*string) *StopCasterResponse {
	s.Headers = v
	return s
}

func (s *StopCasterResponse) SetStatusCode(v int32) *StopCasterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCasterResponse) SetBody(v *StopCasterResponseBody) *StopCasterResponse {
	s.Body = v
	return s
}

func (s *StopCasterResponse) Validate() error {
	return dara.Validate(s)
}

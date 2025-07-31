// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOnlineTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopOnlineTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopOnlineTestResponse
	GetStatusCode() *int32
	SetBody(v *StopOnlineTestResponseBody) *StopOnlineTestResponse
	GetBody() *StopOnlineTestResponseBody
}

type StopOnlineTestResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopOnlineTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopOnlineTestResponse) String() string {
	return dara.Prettify(s)
}

func (s StopOnlineTestResponse) GoString() string {
	return s.String()
}

func (s *StopOnlineTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopOnlineTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopOnlineTestResponse) GetBody() *StopOnlineTestResponseBody {
	return s.Body
}

func (s *StopOnlineTestResponse) SetHeaders(v map[string]*string) *StopOnlineTestResponse {
	s.Headers = v
	return s
}

func (s *StopOnlineTestResponse) SetStatusCode(v int32) *StopOnlineTestResponse {
	s.StatusCode = &v
	return s
}

func (s *StopOnlineTestResponse) SetBody(v *StopOnlineTestResponseBody) *StopOnlineTestResponse {
	s.Body = v
	return s
}

func (s *StopOnlineTestResponse) Validate() error {
	return dara.Validate(s)
}

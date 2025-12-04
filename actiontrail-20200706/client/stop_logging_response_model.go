// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLoggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLoggingResponse
	GetStatusCode() *int32
	SetBody(v *StopLoggingResponseBody) *StopLoggingResponse
	GetBody() *StopLoggingResponseBody
}

type StopLoggingResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLoggingResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLoggingResponse) GoString() string {
	return s.String()
}

func (s *StopLoggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLoggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLoggingResponse) GetBody() *StopLoggingResponseBody {
	return s.Body
}

func (s *StopLoggingResponse) SetHeaders(v map[string]*string) *StopLoggingResponse {
	s.Headers = v
	return s
}

func (s *StopLoggingResponse) SetStatusCode(v int32) *StopLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLoggingResponse) SetBody(v *StopLoggingResponseBody) *StopLoggingResponse {
	s.Body = v
	return s
}

func (s *StopLoggingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDesktopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDesktopResponse
	GetStatusCode() *int32
	SetBody(v *StopDesktopResponseBody) *StopDesktopResponse
	GetBody() *StopDesktopResponseBody
}

type StopDesktopResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDesktopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDesktopResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDesktopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDesktopResponse) GetBody() *StopDesktopResponseBody {
	return s.Body
}

func (s *StopDesktopResponse) SetHeaders(v map[string]*string) *StopDesktopResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopResponse) SetStatusCode(v int32) *StopDesktopResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDesktopResponse) SetBody(v *StopDesktopResponseBody) *StopDesktopResponse {
	s.Body = v
	return s
}

func (s *StopDesktopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

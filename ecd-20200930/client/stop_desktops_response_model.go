// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse
	GetBody() *StopDesktopsResponseBody
}

type StopDesktopsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDesktopsResponse) GetBody() *StopDesktopsResponseBody {
	return s.Body
}

func (s *StopDesktopsResponse) SetHeaders(v map[string]*string) *StopDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopsResponse) SetStatusCode(v int32) *StopDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDesktopsResponse) SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse {
	s.Body = v
	return s
}

func (s *StopDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

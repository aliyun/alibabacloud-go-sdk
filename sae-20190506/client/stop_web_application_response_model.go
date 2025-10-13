// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWebApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopWebApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopWebApplicationResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationBody) *StopWebApplicationResponse
	GetBody() *WebApplicationBody
}

type StopWebApplicationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopWebApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopWebApplicationResponse) GoString() string {
	return s.String()
}

func (s *StopWebApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopWebApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopWebApplicationResponse) GetBody() *WebApplicationBody {
	return s.Body
}

func (s *StopWebApplicationResponse) SetHeaders(v map[string]*string) *StopWebApplicationResponse {
	s.Headers = v
	return s
}

func (s *StopWebApplicationResponse) SetStatusCode(v int32) *StopWebApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopWebApplicationResponse) SetBody(v *WebApplicationBody) *StopWebApplicationResponse {
	s.Body = v
	return s
}

func (s *StopWebApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

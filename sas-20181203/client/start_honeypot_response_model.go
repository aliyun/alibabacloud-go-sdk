// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *StartHoneypotResponseBody) *StartHoneypotResponse
	GetBody() *StartHoneypotResponseBody
}

type StartHoneypotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s StartHoneypotResponse) GoString() string {
	return s.String()
}

func (s *StartHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartHoneypotResponse) GetBody() *StartHoneypotResponseBody {
	return s.Body
}

func (s *StartHoneypotResponse) SetHeaders(v map[string]*string) *StartHoneypotResponse {
	s.Headers = v
	return s
}

func (s *StartHoneypotResponse) SetStatusCode(v int32) *StartHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *StartHoneypotResponse) SetBody(v *StartHoneypotResponseBody) *StartHoneypotResponse {
	s.Body = v
	return s
}

func (s *StartHoneypotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

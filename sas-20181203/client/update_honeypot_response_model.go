// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHoneypotResponseBody) *UpdateHoneypotResponse
	GetBody() *UpdateHoneypotResponseBody
}

type UpdateHoneypotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotResponse) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHoneypotResponse) GetBody() *UpdateHoneypotResponseBody {
	return s.Body
}

func (s *UpdateHoneypotResponse) SetHeaders(v map[string]*string) *UpdateHoneypotResponse {
	s.Headers = v
	return s
}

func (s *UpdateHoneypotResponse) SetStatusCode(v int32) *UpdateHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHoneypotResponse) SetBody(v *UpdateHoneypotResponseBody) *UpdateHoneypotResponse {
	s.Body = v
	return s
}

func (s *UpdateHoneypotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

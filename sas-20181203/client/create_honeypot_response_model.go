// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *CreateHoneypotResponseBody) *CreateHoneypotResponse
	GetBody() *CreateHoneypotResponseBody
}

type CreateHoneypotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotResponse) GoString() string {
	return s.String()
}

func (s *CreateHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHoneypotResponse) GetBody() *CreateHoneypotResponseBody {
	return s.Body
}

func (s *CreateHoneypotResponse) SetHeaders(v map[string]*string) *CreateHoneypotResponse {
	s.Headers = v
	return s
}

func (s *CreateHoneypotResponse) SetStatusCode(v int32) *CreateHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoneypotResponse) SetBody(v *CreateHoneypotResponseBody) *CreateHoneypotResponse {
	s.Body = v
	return s
}

func (s *CreateHoneypotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

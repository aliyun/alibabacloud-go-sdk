// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHoneypotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHoneypotNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateHoneypotNodeResponseBody) *CreateHoneypotNodeResponse
	GetBody() *CreateHoneypotNodeResponseBody
}

type CreateHoneypotNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoneypotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoneypotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateHoneypotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHoneypotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHoneypotNodeResponse) GetBody() *CreateHoneypotNodeResponseBody {
	return s.Body
}

func (s *CreateHoneypotNodeResponse) SetHeaders(v map[string]*string) *CreateHoneypotNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateHoneypotNodeResponse) SetStatusCode(v int32) *CreateHoneypotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoneypotNodeResponse) SetBody(v *CreateHoneypotNodeResponseBody) *CreateHoneypotNodeResponse {
	s.Body = v
	return s
}

func (s *CreateHoneypotNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

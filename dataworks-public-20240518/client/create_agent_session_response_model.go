// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentSessionResponseBody) *CreateAgentSessionResponse
	GetBody() *CreateAgentSessionResponseBody
}

type CreateAgentSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentSessionResponse) GetBody() *CreateAgentSessionResponseBody {
	return s.Body
}

func (s *CreateAgentSessionResponse) SetHeaders(v map[string]*string) *CreateAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentSessionResponse) SetStatusCode(v int32) *CreateAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentSessionResponse) SetBody(v *CreateAgentSessionResponseBody) *CreateAgentSessionResponse {
	s.Body = v
	return s
}

func (s *CreateAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

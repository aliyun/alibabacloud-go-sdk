// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomAgentResponseBody) *CreateCustomAgentResponse
	GetBody() *CreateCustomAgentResponseBody
}

type CreateCustomAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomAgentResponse) GetBody() *CreateCustomAgentResponseBody {
	return s.Body
}

func (s *CreateCustomAgentResponse) SetHeaders(v map[string]*string) *CreateCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomAgentResponse) SetStatusCode(v int32) *CreateCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomAgentResponse) SetBody(v *CreateCustomAgentResponseBody) *CreateCustomAgentResponse {
	s.Body = v
	return s
}

func (s *CreateCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

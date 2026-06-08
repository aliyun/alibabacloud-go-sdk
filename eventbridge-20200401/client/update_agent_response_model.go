// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentResponseBody) *UpdateAgentResponse
	GetBody() *UpdateAgentResponseBody
}

type UpdateAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentResponse) GetBody() *UpdateAgentResponseBody {
	return s.Body
}

func (s *UpdateAgentResponse) SetHeaders(v map[string]*string) *UpdateAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentResponse) SetStatusCode(v int32) *UpdateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentResponse) SetBody(v *UpdateAgentResponseBody) *UpdateAgentResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

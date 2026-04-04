// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAssistantAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppAssistantAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppAssistantAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppAssistantAgentResponseBody) *CreateAppAssistantAgentResponse
	GetBody() *CreateAppAssistantAgentResponseBody
}

type CreateAppAssistantAgentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppAssistantAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppAssistantAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppAssistantAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppAssistantAgentResponse) GetBody() *CreateAppAssistantAgentResponseBody {
	return s.Body
}

func (s *CreateAppAssistantAgentResponse) SetHeaders(v map[string]*string) *CreateAppAssistantAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAppAssistantAgentResponse) SetStatusCode(v int32) *CreateAppAssistantAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppAssistantAgentResponse) SetBody(v *CreateAppAssistantAgentResponseBody) *CreateAppAssistantAgentResponse {
	s.Body = v
	return s
}

func (s *CreateAppAssistantAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

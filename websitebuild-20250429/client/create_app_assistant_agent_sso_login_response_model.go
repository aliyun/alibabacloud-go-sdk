// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAssistantAgentSsoLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppAssistantAgentSsoLoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppAssistantAgentSsoLoginResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppAssistantAgentSsoLoginResponseBody) *CreateAppAssistantAgentSsoLoginResponse
	GetBody() *CreateAppAssistantAgentSsoLoginResponseBody
}

type CreateAppAssistantAgentSsoLoginResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppAssistantAgentSsoLoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppAssistantAgentSsoLoginResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentSsoLoginResponse) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentSsoLoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppAssistantAgentSsoLoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppAssistantAgentSsoLoginResponse) GetBody() *CreateAppAssistantAgentSsoLoginResponseBody {
	return s.Body
}

func (s *CreateAppAssistantAgentSsoLoginResponse) SetHeaders(v map[string]*string) *CreateAppAssistantAgentSsoLoginResponse {
	s.Headers = v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponse) SetStatusCode(v int32) *CreateAppAssistantAgentSsoLoginResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponse) SetBody(v *CreateAppAssistantAgentSsoLoginResponseBody) *CreateAppAssistantAgentSsoLoginResponse {
	s.Body = v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

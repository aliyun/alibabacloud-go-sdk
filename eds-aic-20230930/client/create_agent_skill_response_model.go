// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentSkillResponseBody) *CreateAgentSkillResponse
	GetBody() *CreateAgentSkillResponseBody
}

type CreateAgentSkillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSkillResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentSkillResponse) GetBody() *CreateAgentSkillResponseBody {
	return s.Body
}

func (s *CreateAgentSkillResponse) SetHeaders(v map[string]*string) *CreateAgentSkillResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentSkillResponse) SetStatusCode(v int32) *CreateAgentSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentSkillResponse) SetBody(v *CreateAgentSkillResponseBody) *CreateAgentSkillResponse {
	s.Body = v
	return s
}

func (s *CreateAgentSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

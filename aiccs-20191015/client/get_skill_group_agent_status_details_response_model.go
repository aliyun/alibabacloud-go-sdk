// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAgentStatusDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupAgentStatusDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupAgentStatusDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupAgentStatusDetailsResponseBody) *GetSkillGroupAgentStatusDetailsResponse
	GetBody() *GetSkillGroupAgentStatusDetailsResponseBody
}

type GetSkillGroupAgentStatusDetailsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupAgentStatusDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupAgentStatusDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupAgentStatusDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupAgentStatusDetailsResponse) GetBody() *GetSkillGroupAgentStatusDetailsResponseBody {
	return s.Body
}

func (s *GetSkillGroupAgentStatusDetailsResponse) SetHeaders(v map[string]*string) *GetSkillGroupAgentStatusDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponse) SetStatusCode(v int32) *GetSkillGroupAgentStatusDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponse) SetBody(v *GetSkillGroupAgentStatusDetailsResponseBody) *GetSkillGroupAgentStatusDetailsResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

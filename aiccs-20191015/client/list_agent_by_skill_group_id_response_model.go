// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentBySkillGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentBySkillGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentBySkillGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentBySkillGroupIdResponseBody) *ListAgentBySkillGroupIdResponse
	GetBody() *ListAgentBySkillGroupIdResponseBody
}

type ListAgentBySkillGroupIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentBySkillGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentBySkillGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentBySkillGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentBySkillGroupIdResponse) GetBody() *ListAgentBySkillGroupIdResponseBody {
	return s.Body
}

func (s *ListAgentBySkillGroupIdResponse) SetHeaders(v map[string]*string) *ListAgentBySkillGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) SetStatusCode(v int32) *ListAgentBySkillGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) SetBody(v *ListAgentBySkillGroupIdResponseBody) *ListAgentBySkillGroupIdResponse {
	s.Body = v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

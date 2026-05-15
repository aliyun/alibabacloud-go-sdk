// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAgentFromSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAgentFromSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAgentFromSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAgentFromSkillGroupResponseBody) *RemoveAgentFromSkillGroupResponse
	GetBody() *RemoveAgentFromSkillGroupResponseBody
}

type RemoveAgentFromSkillGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAgentFromSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAgentFromSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAgentFromSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveAgentFromSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAgentFromSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAgentFromSkillGroupResponse) GetBody() *RemoveAgentFromSkillGroupResponseBody {
	return s.Body
}

func (s *RemoveAgentFromSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveAgentFromSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveAgentFromSkillGroupResponse) SetStatusCode(v int32) *RemoveAgentFromSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAgentFromSkillGroupResponse) SetBody(v *RemoveAgentFromSkillGroupResponseBody) *RemoveAgentFromSkillGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveAgentFromSkillGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

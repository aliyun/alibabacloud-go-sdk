// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentSkillResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentSkillResponseBody) *DeleteAgentSkillResponse
	GetBody() *DeleteAgentSkillResponseBody
}

type DeleteAgentSkillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSkillResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentSkillResponse) GetBody() *DeleteAgentSkillResponseBody {
	return s.Body
}

func (s *DeleteAgentSkillResponse) SetHeaders(v map[string]*string) *DeleteAgentSkillResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentSkillResponse) SetStatusCode(v int32) *DeleteAgentSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentSkillResponse) SetBody(v *DeleteAgentSkillResponseBody) *DeleteAgentSkillResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

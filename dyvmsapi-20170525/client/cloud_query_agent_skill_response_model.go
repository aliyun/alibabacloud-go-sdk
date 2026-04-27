// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryAgentSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryAgentSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryAgentSkillResponseBody) *CloudQueryAgentSkillResponse
	GetBody() *CloudQueryAgentSkillResponseBody
}

type CloudQueryAgentSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryAgentSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryAgentSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryAgentSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryAgentSkillResponse) GetBody() *CloudQueryAgentSkillResponseBody {
	return s.Body
}

func (s *CloudQueryAgentSkillResponse) SetHeaders(v map[string]*string) *CloudQueryAgentSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryAgentSkillResponse) SetStatusCode(v int32) *CloudQueryAgentSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryAgentSkillResponse) SetBody(v *CloudQueryAgentSkillResponseBody) *CloudQueryAgentSkillResponse {
	s.Body = v
	return s
}

func (s *CloudQueryAgentSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentSkillMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentSkillMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentSkillMetaResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentSkillMetaResponseBody) *ListDataAgentSkillMetaResponse
	GetBody() *ListDataAgentSkillMetaResponseBody
}

type ListDataAgentSkillMetaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentSkillMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentSkillMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSkillMetaResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentSkillMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentSkillMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentSkillMetaResponse) GetBody() *ListDataAgentSkillMetaResponseBody {
	return s.Body
}

func (s *ListDataAgentSkillMetaResponse) SetHeaders(v map[string]*string) *ListDataAgentSkillMetaResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentSkillMetaResponse) SetStatusCode(v int32) *ListDataAgentSkillMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentSkillMetaResponse) SetBody(v *ListDataAgentSkillMetaResponseBody) *ListDataAgentSkillMetaResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentSkillMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

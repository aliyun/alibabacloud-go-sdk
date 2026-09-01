// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSkillMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentSkillMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentSkillMetaResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentSkillMetaResponseBody) *CreateDataAgentSkillMetaResponse
	GetBody() *CreateDataAgentSkillMetaResponseBody
}

type CreateDataAgentSkillMetaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentSkillMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentSkillMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSkillMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSkillMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentSkillMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentSkillMetaResponse) GetBody() *CreateDataAgentSkillMetaResponseBody {
	return s.Body
}

func (s *CreateDataAgentSkillMetaResponse) SetHeaders(v map[string]*string) *CreateDataAgentSkillMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentSkillMetaResponse) SetStatusCode(v int32) *CreateDataAgentSkillMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponse) SetBody(v *CreateDataAgentSkillMetaResponseBody) *CreateDataAgentSkillMetaResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentSkillMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

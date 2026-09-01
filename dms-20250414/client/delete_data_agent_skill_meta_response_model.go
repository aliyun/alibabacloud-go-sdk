// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentSkillMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentSkillMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentSkillMetaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentSkillMetaResponseBody) *DeleteDataAgentSkillMetaResponse
	GetBody() *DeleteDataAgentSkillMetaResponseBody
}

type DeleteDataAgentSkillMetaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentSkillMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentSkillMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentSkillMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentSkillMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentSkillMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentSkillMetaResponse) GetBody() *DeleteDataAgentSkillMetaResponseBody {
	return s.Body
}

func (s *DeleteDataAgentSkillMetaResponse) SetHeaders(v map[string]*string) *DeleteDataAgentSkillMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentSkillMetaResponse) SetStatusCode(v int32) *DeleteDataAgentSkillMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponse) SetBody(v *DeleteDataAgentSkillMetaResponseBody) *DeleteDataAgentSkillMetaResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentSkillMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

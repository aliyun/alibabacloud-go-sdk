// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolarClawAgentSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolarClawAgentSkillsResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolarClawAgentSkillsResponseBody) *UpdatePolarClawAgentSkillsResponse
	GetBody() *UpdatePolarClawAgentSkillsResponseBody
}

type UpdatePolarClawAgentSkillsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarClawAgentSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolarClawAgentSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentSkillsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolarClawAgentSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolarClawAgentSkillsResponse) GetBody() *UpdatePolarClawAgentSkillsResponseBody {
	return s.Body
}

func (s *UpdatePolarClawAgentSkillsResponse) SetHeaders(v map[string]*string) *UpdatePolarClawAgentSkillsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponse) SetStatusCode(v int32) *UpdatePolarClawAgentSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponse) SetBody(v *UpdatePolarClawAgentSkillsResponseBody) *UpdatePolarClawAgentSkillsResponse {
	s.Body = v
	return s
}

func (s *UpdatePolarClawAgentSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

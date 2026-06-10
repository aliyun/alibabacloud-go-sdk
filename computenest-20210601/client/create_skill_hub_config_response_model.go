// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillHubConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillHubConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillHubConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillHubConfigResponseBody) *CreateSkillHubConfigResponse
	GetBody() *CreateSkillHubConfigResponseBody
}

type CreateSkillHubConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillHubConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillHubConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillHubConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillHubConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillHubConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillHubConfigResponse) GetBody() *CreateSkillHubConfigResponseBody {
	return s.Body
}

func (s *CreateSkillHubConfigResponse) SetHeaders(v map[string]*string) *CreateSkillHubConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillHubConfigResponse) SetStatusCode(v int32) *CreateSkillHubConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillHubConfigResponse) SetBody(v *CreateSkillHubConfigResponseBody) *CreateSkillHubConfigResponse {
	s.Body = v
	return s
}

func (s *CreateSkillHubConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

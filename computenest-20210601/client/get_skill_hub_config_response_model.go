// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillHubConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillHubConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillHubConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillHubConfigResponseBody) *GetSkillHubConfigResponse
	GetBody() *GetSkillHubConfigResponseBody
}

type GetSkillHubConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillHubConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillHubConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillHubConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSkillHubConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillHubConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillHubConfigResponse) GetBody() *GetSkillHubConfigResponseBody {
	return s.Body
}

func (s *GetSkillHubConfigResponse) SetHeaders(v map[string]*string) *GetSkillHubConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSkillHubConfigResponse) SetStatusCode(v int32) *GetSkillHubConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillHubConfigResponse) SetBody(v *GetSkillHubConfigResponseBody) *GetSkillHubConfigResponse {
	s.Body = v
	return s
}

func (s *GetSkillHubConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

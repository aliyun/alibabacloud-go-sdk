// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillGroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillGroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillGroupConfigResponseBody) *CreateSkillGroupConfigResponse
	GetBody() *CreateSkillGroupConfigResponseBody
}

type CreateSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillGroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillGroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillGroupConfigResponse) GetBody() *CreateSkillGroupConfigResponseBody {
	return s.Body
}

func (s *CreateSkillGroupConfigResponse) SetHeaders(v map[string]*string) *CreateSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupConfigResponse) SetStatusCode(v int32) *CreateSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillGroupConfigResponse) SetBody(v *CreateSkillGroupConfigResponseBody) *CreateSkillGroupConfigResponse {
	s.Body = v
	return s
}

func (s *CreateSkillGroupConfigResponse) Validate() error {
	return dara.Validate(s)
}

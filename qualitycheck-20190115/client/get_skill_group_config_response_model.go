// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupConfigResponseBody) *GetSkillGroupConfigResponse
	GetBody() *GetSkillGroupConfigResponseBody
}

type GetSkillGroupConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupConfigResponse) GetBody() *GetSkillGroupConfigResponseBody {
	return s.Body
}

func (s *GetSkillGroupConfigResponse) SetHeaders(v map[string]*string) *GetSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupConfigResponse) SetStatusCode(v int32) *GetSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupConfigResponse) SetBody(v *GetSkillGroupConfigResponseBody) *GetSkillGroupConfigResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupConfigResponse) Validate() error {
	return dara.Validate(s)
}

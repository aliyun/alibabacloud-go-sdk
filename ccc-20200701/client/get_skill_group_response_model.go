// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupResponseBody) *GetSkillGroupResponse
	GetBody() *GetSkillGroupResponseBody
}

type GetSkillGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupResponse) GetBody() *GetSkillGroupResponseBody {
	return s.Body
}

func (s *GetSkillGroupResponse) SetHeaders(v map[string]*string) *GetSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupResponse) SetStatusCode(v int32) *GetSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupResponse) SetBody(v *GetSkillGroupResponseBody) *GetSkillGroupResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

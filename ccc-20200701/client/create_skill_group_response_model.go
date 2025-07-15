// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillGroupResponseBody) *CreateSkillGroupResponse
	GetBody() *CreateSkillGroupResponseBody
}

type CreateSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillGroupResponse) GetBody() *CreateSkillGroupResponseBody {
	return s.Body
}

func (s *CreateSkillGroupResponse) SetHeaders(v map[string]*string) *CreateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupResponse) SetStatusCode(v int32) *CreateSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillGroupResponse) SetBody(v *CreateSkillGroupResponseBody) *CreateSkillGroupResponse {
	s.Body = v
	return s
}

func (s *CreateSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

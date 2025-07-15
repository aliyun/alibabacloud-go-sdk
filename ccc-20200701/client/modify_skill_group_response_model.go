// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifySkillGroupResponseBody) *ModifySkillGroupResponse
	GetBody() *ModifySkillGroupResponseBody
}

type ModifySkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifySkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySkillGroupResponse) GetBody() *ModifySkillGroupResponseBody {
	return s.Body
}

func (s *ModifySkillGroupResponse) SetHeaders(v map[string]*string) *ModifySkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifySkillGroupResponse) SetStatusCode(v int32) *ModifySkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySkillGroupResponse) SetBody(v *ModifySkillGroupResponseBody) *ModifySkillGroupResponse {
	s.Body = v
	return s
}

func (s *ModifySkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

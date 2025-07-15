// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserLevelsOfSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserLevelsOfSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserLevelsOfSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserLevelsOfSkillGroupResponseBody) *ModifyUserLevelsOfSkillGroupResponse
	GetBody() *ModifyUserLevelsOfSkillGroupResponseBody
}

type ModifyUserLevelsOfSkillGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserLevelsOfSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserLevelsOfSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserLevelsOfSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserLevelsOfSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserLevelsOfSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserLevelsOfSkillGroupResponse) GetBody() *ModifyUserLevelsOfSkillGroupResponseBody {
	return s.Body
}

func (s *ModifyUserLevelsOfSkillGroupResponse) SetHeaders(v map[string]*string) *ModifyUserLevelsOfSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponse) SetStatusCode(v int32) *ModifyUserLevelsOfSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponse) SetBody(v *ModifyUserLevelsOfSkillGroupResponseBody) *ModifyUserLevelsOfSkillGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

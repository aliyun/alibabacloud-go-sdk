// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUsersToSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUsersToSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddUsersToSkillGroupResponseBody) *AddUsersToSkillGroupResponse
	GetBody() *AddUsersToSkillGroupResponseBody
}

type AddUsersToSkillGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUsersToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUsersToSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUsersToSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUsersToSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUsersToSkillGroupResponse) GetBody() *AddUsersToSkillGroupResponseBody {
	return s.Body
}

func (s *AddUsersToSkillGroupResponse) SetHeaders(v map[string]*string) *AddUsersToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUsersToSkillGroupResponse) SetStatusCode(v int32) *AddUsersToSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUsersToSkillGroupResponse) SetBody(v *AddUsersToSkillGroupResponseBody) *AddUsersToSkillGroupResponse {
	s.Body = v
	return s
}

func (s *AddUsersToSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

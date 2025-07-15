// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSkillGroupsToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSkillGroupsToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSkillGroupsToUserResponse
	GetStatusCode() *int32
	SetBody(v *AddSkillGroupsToUserResponseBody) *AddSkillGroupsToUserResponse
	GetBody() *AddSkillGroupsToUserResponseBody
}

type AddSkillGroupsToUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSkillGroupsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSkillGroupsToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSkillGroupsToUserResponse) GoString() string {
	return s.String()
}

func (s *AddSkillGroupsToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSkillGroupsToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSkillGroupsToUserResponse) GetBody() *AddSkillGroupsToUserResponseBody {
	return s.Body
}

func (s *AddSkillGroupsToUserResponse) SetHeaders(v map[string]*string) *AddSkillGroupsToUserResponse {
	s.Headers = v
	return s
}

func (s *AddSkillGroupsToUserResponse) SetStatusCode(v int32) *AddSkillGroupsToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSkillGroupsToUserResponse) SetBody(v *AddSkillGroupsToUserResponseBody) *AddSkillGroupsToUserResponse {
	s.Body = v
	return s
}

func (s *AddSkillGroupsToUserResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUsersFromSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUsersFromSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUsersFromSkillGroupResponseBody) *RemoveUsersFromSkillGroupResponse
	GetBody() *RemoveUsersFromSkillGroupResponseBody
}

type RemoveUsersFromSkillGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersFromSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUsersFromSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUsersFromSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUsersFromSkillGroupResponse) GetBody() *RemoveUsersFromSkillGroupResponseBody {
	return s.Body
}

func (s *RemoveUsersFromSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveUsersFromSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersFromSkillGroupResponse) SetStatusCode(v int32) *RemoveUsersFromSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponse) SetBody(v *RemoveUsersFromSkillGroupResponseBody) *RemoveUsersFromSkillGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveUsersFromSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

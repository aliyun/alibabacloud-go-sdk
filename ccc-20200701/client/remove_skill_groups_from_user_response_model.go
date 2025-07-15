// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSkillGroupsFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSkillGroupsFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSkillGroupsFromUserResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSkillGroupsFromUserResponseBody) *RemoveSkillGroupsFromUserResponse
	GetBody() *RemoveSkillGroupsFromUserResponseBody
}

type RemoveSkillGroupsFromUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSkillGroupsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSkillGroupsFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSkillGroupsFromUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupsFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSkillGroupsFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSkillGroupsFromUserResponse) GetBody() *RemoveSkillGroupsFromUserResponseBody {
	return s.Body
}

func (s *RemoveSkillGroupsFromUserResponse) SetHeaders(v map[string]*string) *RemoveSkillGroupsFromUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveSkillGroupsFromUserResponse) SetStatusCode(v int32) *RemoveSkillGroupsFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponse) SetBody(v *RemoveSkillGroupsFromUserResponseBody) *RemoveSkillGroupsFromUserResponse {
	s.Body = v
	return s
}

func (s *RemoveSkillGroupsFromUserResponse) Validate() error {
	return dara.Validate(s)
}

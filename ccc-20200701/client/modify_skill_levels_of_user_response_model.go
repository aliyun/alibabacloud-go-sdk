// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySkillLevelsOfUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySkillLevelsOfUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySkillLevelsOfUserResponse
	GetStatusCode() *int32
	SetBody(v *ModifySkillLevelsOfUserResponseBody) *ModifySkillLevelsOfUserResponse
	GetBody() *ModifySkillLevelsOfUserResponseBody
}

type ModifySkillLevelsOfUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySkillLevelsOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySkillLevelsOfUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySkillLevelsOfUserResponse) GoString() string {
	return s.String()
}

func (s *ModifySkillLevelsOfUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySkillLevelsOfUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySkillLevelsOfUserResponse) GetBody() *ModifySkillLevelsOfUserResponseBody {
	return s.Body
}

func (s *ModifySkillLevelsOfUserResponse) SetHeaders(v map[string]*string) *ModifySkillLevelsOfUserResponse {
	s.Headers = v
	return s
}

func (s *ModifySkillLevelsOfUserResponse) SetStatusCode(v int32) *ModifySkillLevelsOfUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponse) SetBody(v *ModifySkillLevelsOfUserResponseBody) *ModifySkillLevelsOfUserResponse {
	s.Body = v
	return s
}

func (s *ModifySkillLevelsOfUserResponse) Validate() error {
	return dara.Validate(s)
}

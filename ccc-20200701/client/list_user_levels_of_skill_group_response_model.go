// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserLevelsOfSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserLevelsOfSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserLevelsOfSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListUserLevelsOfSkillGroupResponseBody) *ListUserLevelsOfSkillGroupResponse
	GetBody() *ListUserLevelsOfSkillGroupResponseBody
}

type ListUserLevelsOfSkillGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserLevelsOfSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserLevelsOfSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserLevelsOfSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserLevelsOfSkillGroupResponse) GetBody() *ListUserLevelsOfSkillGroupResponseBody {
	return s.Body
}

func (s *ListUserLevelsOfSkillGroupResponse) SetHeaders(v map[string]*string) *ListUserLevelsOfSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponse) SetStatusCode(v int32) *ListUserLevelsOfSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponse) SetBody(v *ListUserLevelsOfSkillGroupResponseBody) *ListUserLevelsOfSkillGroupResponse {
	s.Body = v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

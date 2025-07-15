// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillLevelsOfUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillLevelsOfUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillLevelsOfUserResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillLevelsOfUserResponseBody) *ListSkillLevelsOfUserResponse
	GetBody() *ListSkillLevelsOfUserResponseBody
}

type ListSkillLevelsOfUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillLevelsOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillLevelsOfUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillLevelsOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillLevelsOfUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillLevelsOfUserResponse) GetBody() *ListSkillLevelsOfUserResponseBody {
	return s.Body
}

func (s *ListSkillLevelsOfUserResponse) SetHeaders(v map[string]*string) *ListSkillLevelsOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListSkillLevelsOfUserResponse) SetStatusCode(v int32) *ListSkillLevelsOfUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillLevelsOfUserResponse) SetBody(v *ListSkillLevelsOfUserResponseBody) *ListSkillLevelsOfUserResponse {
	s.Body = v
	return s
}

func (s *ListSkillLevelsOfUserResponse) Validate() error {
	return dara.Validate(s)
}

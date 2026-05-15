// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddSkillGroupResponseBody) *AddSkillGroupResponse
	GetBody() *AddSkillGroupResponseBody
}

type AddSkillGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *AddSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSkillGroupResponse) GetBody() *AddSkillGroupResponseBody {
	return s.Body
}

func (s *AddSkillGroupResponse) SetHeaders(v map[string]*string) *AddSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *AddSkillGroupResponse) SetStatusCode(v int32) *AddSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSkillGroupResponse) SetBody(v *AddSkillGroupResponseBody) *AddSkillGroupResponse {
	s.Body = v
	return s
}

func (s *AddSkillGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillGroupResponseBody) *UpdateSkillGroupResponse
	GetBody() *UpdateSkillGroupResponseBody
}

type UpdateSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillGroupResponse) GetBody() *UpdateSkillGroupResponseBody {
	return s.Body
}

func (s *UpdateSkillGroupResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupResponse) SetStatusCode(v int32) *UpdateSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillGroupResponse) SetBody(v *UpdateSkillGroupResponseBody) *UpdateSkillGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

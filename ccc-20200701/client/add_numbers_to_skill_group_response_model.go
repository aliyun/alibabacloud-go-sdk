// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddNumbersToSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddNumbersToSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddNumbersToSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddNumbersToSkillGroupResponseBody) *AddNumbersToSkillGroupResponse
	GetBody() *AddNumbersToSkillGroupResponseBody
}

type AddNumbersToSkillGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddNumbersToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddNumbersToSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddNumbersToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *AddNumbersToSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddNumbersToSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddNumbersToSkillGroupResponse) GetBody() *AddNumbersToSkillGroupResponseBody {
	return s.Body
}

func (s *AddNumbersToSkillGroupResponse) SetHeaders(v map[string]*string) *AddNumbersToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *AddNumbersToSkillGroupResponse) SetStatusCode(v int32) *AddNumbersToSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddNumbersToSkillGroupResponse) SetBody(v *AddNumbersToSkillGroupResponseBody) *AddNumbersToSkillGroupResponse {
	s.Body = v
	return s
}

func (s *AddNumbersToSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

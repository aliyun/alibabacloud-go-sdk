// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPhoneNumberToSkillGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPhoneNumberToSkillGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPhoneNumberToSkillGroupsResponse
	GetStatusCode() *int32
	SetBody(v *AddPhoneNumberToSkillGroupsResponseBody) *AddPhoneNumberToSkillGroupsResponse
	GetBody() *AddPhoneNumberToSkillGroupsResponseBody
}

type AddPhoneNumberToSkillGroupsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPhoneNumberToSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPhoneNumberToSkillGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPhoneNumberToSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *AddPhoneNumberToSkillGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPhoneNumberToSkillGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPhoneNumberToSkillGroupsResponse) GetBody() *AddPhoneNumberToSkillGroupsResponseBody {
	return s.Body
}

func (s *AddPhoneNumberToSkillGroupsResponse) SetHeaders(v map[string]*string) *AddPhoneNumberToSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponse) SetStatusCode(v int32) *AddPhoneNumberToSkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponse) SetBody(v *AddPhoneNumberToSkillGroupsResponseBody) *AddPhoneNumberToSkillGroupsResponse {
	s.Body = v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponse) Validate() error {
	return dara.Validate(s)
}

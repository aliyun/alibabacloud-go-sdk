// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneNumbersOfSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPhoneNumbersOfSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPhoneNumbersOfSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListPhoneNumbersOfSkillGroupResponseBody) *ListPhoneNumbersOfSkillGroupResponse
	GetBody() *ListPhoneNumbersOfSkillGroupResponseBody
}

type ListPhoneNumbersOfSkillGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPhoneNumbersOfSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPhoneNumbersOfSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPhoneNumbersOfSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPhoneNumbersOfSkillGroupResponse) GetBody() *ListPhoneNumbersOfSkillGroupResponseBody {
	return s.Body
}

func (s *ListPhoneNumbersOfSkillGroupResponse) SetHeaders(v map[string]*string) *ListPhoneNumbersOfSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponse) SetStatusCode(v int32) *ListPhoneNumbersOfSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponse) SetBody(v *ListPhoneNumbersOfSkillGroupResponseBody) *ListPhoneNumbersOfSkillGroupResponse {
	s.Body = v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

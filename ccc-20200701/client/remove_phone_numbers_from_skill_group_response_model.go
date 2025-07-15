// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumbersFromSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePhoneNumbersFromSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePhoneNumbersFromSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemovePhoneNumbersFromSkillGroupResponseBody) *RemovePhoneNumbersFromSkillGroupResponse
	GetBody() *RemovePhoneNumbersFromSkillGroupResponseBody
}

type RemovePhoneNumbersFromSkillGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePhoneNumbersFromSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePhoneNumbersFromSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumbersFromSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) GetBody() *RemovePhoneNumbersFromSkillGroupResponseBody {
	return s.Body
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) SetHeaders(v map[string]*string) *RemovePhoneNumbersFromSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) SetStatusCode(v int32) *RemovePhoneNumbersFromSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) SetBody(v *RemovePhoneNumbersFromSkillGroupResponseBody) *RemovePhoneNumbersFromSkillGroupResponse {
	s.Body = v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}

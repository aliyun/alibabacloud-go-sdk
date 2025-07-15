// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumberFromSkillGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePhoneNumberFromSkillGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePhoneNumberFromSkillGroupsResponse
	GetStatusCode() *int32
	SetBody(v *RemovePhoneNumberFromSkillGroupsResponseBody) *RemovePhoneNumberFromSkillGroupsResponse
	GetBody() *RemovePhoneNumberFromSkillGroupsResponseBody
}

type RemovePhoneNumberFromSkillGroupsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePhoneNumberFromSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePhoneNumberFromSkillGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumberFromSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) GetBody() *RemovePhoneNumberFromSkillGroupsResponseBody {
	return s.Body
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) SetHeaders(v map[string]*string) *RemovePhoneNumberFromSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) SetStatusCode(v int32) *RemovePhoneNumberFromSkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) SetBody(v *RemovePhoneNumberFromSkillGroupsResponseBody) *RemovePhoneNumberFromSkillGroupsResponse {
	s.Body = v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) Validate() error {
	return dara.Validate(s)
}

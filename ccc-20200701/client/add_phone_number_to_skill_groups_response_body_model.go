// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPhoneNumberToSkillGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddPhoneNumberToSkillGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddPhoneNumberToSkillGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddPhoneNumberToSkillGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPhoneNumberToSkillGroupsResponseBody
	GetRequestId() *string
}

type AddPhoneNumberToSkillGroupsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPhoneNumberToSkillGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPhoneNumberToSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetCode(v string) *AddPhoneNumberToSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetHttpStatusCode(v int32) *AddPhoneNumberToSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetMessage(v string) *AddPhoneNumberToSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetRequestId(v string) *AddPhoneNumberToSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

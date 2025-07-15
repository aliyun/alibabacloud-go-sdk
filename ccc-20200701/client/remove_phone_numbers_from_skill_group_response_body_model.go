// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumbersFromSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemovePhoneNumbersFromSkillGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemovePhoneNumbersFromSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemovePhoneNumbersFromSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemovePhoneNumbersFromSkillGroupResponseBody
	GetRequestId() *string
}

type RemovePhoneNumbersFromSkillGroupResponseBody struct {
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

func (s RemovePhoneNumbersFromSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumbersFromSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetCode(v string) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetHttpStatusCode(v int32) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetMessage(v string) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetRequestId(v string) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

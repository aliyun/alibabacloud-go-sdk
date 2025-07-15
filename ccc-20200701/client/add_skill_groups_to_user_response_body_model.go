// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSkillGroupsToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSkillGroupsToUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddSkillGroupsToUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddSkillGroupsToUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSkillGroupsToUserResponseBody
	GetRequestId() *string
}

type AddSkillGroupsToUserResponseBody struct {
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
	// BA7F9545-8312-4190-9BD0-63144B3F1ACC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSkillGroupsToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSkillGroupsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddSkillGroupsToUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSkillGroupsToUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddSkillGroupsToUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSkillGroupsToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSkillGroupsToUserResponseBody) SetCode(v string) *AddSkillGroupsToUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetHttpStatusCode(v int32) *AddSkillGroupsToUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetMessage(v string) *AddSkillGroupsToUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetRequestId(v string) *AddSkillGroupsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddUsersToSkillGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddUsersToSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddUsersToSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddUsersToSkillGroupResponseBody
	GetRequestId() *string
}

type AddUsersToSkillGroupResponseBody struct {
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

func (s AddUsersToSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersToSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddUsersToSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddUsersToSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddUsersToSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUsersToSkillGroupResponseBody) SetCode(v string) *AddUsersToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) SetHttpStatusCode(v int32) *AddUsersToSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) SetMessage(v string) *AddUsersToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) SetRequestId(v string) *AddUsersToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

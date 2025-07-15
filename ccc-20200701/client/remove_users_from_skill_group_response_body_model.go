// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveUsersFromSkillGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveUsersFromSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveUsersFromSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveUsersFromSkillGroupResponseBody
	GetRequestId() *string
}

type RemoveUsersFromSkillGroupResponseBody struct {
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

func (s RemoveUsersFromSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveUsersFromSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveUsersFromSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveUsersFromSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetCode(v string) *RemoveUsersFromSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetHttpStatusCode(v int32) *RemoveUsersFromSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetMessage(v string) *RemoveUsersFromSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetRequestId(v string) *RemoveUsersFromSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

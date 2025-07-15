// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSkillGroupsFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveSkillGroupsFromUserResponseBody
	GetCode() *string
	SetData(v string) *RemoveSkillGroupsFromUserResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *RemoveSkillGroupsFromUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveSkillGroupsFromUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveSkillGroupsFromUserResponseBody
	GetRequestId() *string
}

type RemoveSkillGroupsFromUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s RemoveSkillGroupsFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSkillGroupsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupsFromUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveSkillGroupsFromUserResponseBody) GetData() *string {
	return s.Data
}

func (s *RemoveSkillGroupsFromUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveSkillGroupsFromUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveSkillGroupsFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetCode(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetData(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetHttpStatusCode(v int32) *RemoveSkillGroupsFromUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetMessage(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetRequestId(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}

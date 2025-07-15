// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySkillGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifySkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifySkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySkillGroupResponseBody
	GetRequestId() *string
}

type ModifySkillGroupResponseBody struct {
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

func (s ModifySkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifySkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySkillGroupResponseBody) SetCode(v string) *ModifySkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySkillGroupResponseBody) SetHttpStatusCode(v int32) *ModifySkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySkillGroupResponseBody) SetMessage(v string) *ModifySkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySkillGroupResponseBody) SetRequestId(v string) *ModifySkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

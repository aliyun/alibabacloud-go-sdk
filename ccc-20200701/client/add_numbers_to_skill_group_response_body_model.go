// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddNumbersToSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddNumbersToSkillGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddNumbersToSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddNumbersToSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddNumbersToSkillGroupResponseBody
	GetRequestId() *string
}

type AddNumbersToSkillGroupResponseBody struct {
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

func (s AddNumbersToSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddNumbersToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddNumbersToSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddNumbersToSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddNumbersToSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddNumbersToSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddNumbersToSkillGroupResponseBody) SetCode(v string) *AddNumbersToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) SetHttpStatusCode(v int32) *AddNumbersToSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) SetMessage(v string) *AddNumbersToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) SetRequestId(v string) *AddNumbersToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

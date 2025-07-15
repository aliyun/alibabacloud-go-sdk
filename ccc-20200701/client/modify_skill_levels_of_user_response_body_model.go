// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySkillLevelsOfUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySkillLevelsOfUserResponseBody
	GetCode() *string
	SetData(v string) *ModifySkillLevelsOfUserResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifySkillLevelsOfUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifySkillLevelsOfUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySkillLevelsOfUserResponseBody
	GetRequestId() *string
}

type ModifySkillLevelsOfUserResponseBody struct {
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
	// E49D8B83-A3EC-44D4-A920-578BC3C698AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySkillLevelsOfUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySkillLevelsOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySkillLevelsOfUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySkillLevelsOfUserResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifySkillLevelsOfUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifySkillLevelsOfUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySkillLevelsOfUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySkillLevelsOfUserResponseBody) SetCode(v string) *ModifySkillLevelsOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetData(v string) *ModifySkillLevelsOfUserResponseBody {
	s.Data = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetHttpStatusCode(v int32) *ModifySkillLevelsOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetMessage(v string) *ModifySkillLevelsOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetRequestId(v string) *ModifySkillLevelsOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) Validate() error {
	return dara.Validate(s)
}

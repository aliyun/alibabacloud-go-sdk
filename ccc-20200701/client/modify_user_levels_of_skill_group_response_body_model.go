// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserLevelsOfSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyUserLevelsOfSkillGroupResponseBody
	GetCode() *string
	SetData(v string) *ModifyUserLevelsOfSkillGroupResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifyUserLevelsOfSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyUserLevelsOfSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyUserLevelsOfSkillGroupResponseBody
	GetRequestId() *string
}

type ModifyUserLevelsOfSkillGroupResponseBody struct {
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

func (s ModifyUserLevelsOfSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserLevelsOfSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetCode(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetData(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetHttpStatusCode(v int32) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetMessage(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetRequestId(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

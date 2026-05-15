// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSkillGroupResponseBody
	GetCode() *string
	SetData(v string) *AddSkillGroupResponseBody
	GetData() *string
	SetMessage(v string) *AddSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSkillGroupResponseBody
	GetSuccess() *bool
}

type AddSkillGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSkillGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *AddSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSkillGroupResponseBody) SetCode(v string) *AddSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetData(v string) *AddSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetMessage(v string) *AddSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetRequestId(v string) *AddSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSkillGroupResponseBody) SetSuccess(v bool) *AddSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

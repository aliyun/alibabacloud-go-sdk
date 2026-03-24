// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDigitalEmployeeSkillResponseBody
	GetRequestId() *string
	SetSkillName(v string) *CreateDigitalEmployeeSkillResponseBody
	GetSkillName() *string
}

type CreateDigitalEmployeeSkillResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// skill
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
}

func (s CreateDigitalEmployeeSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDigitalEmployeeSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateDigitalEmployeeSkillResponseBody) SetRequestId(v string) *CreateDigitalEmployeeSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDigitalEmployeeSkillResponseBody) SetSkillName(v string) *CreateDigitalEmployeeSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *CreateDigitalEmployeeSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

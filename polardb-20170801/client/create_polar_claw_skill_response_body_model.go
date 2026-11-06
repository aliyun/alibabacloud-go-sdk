// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsUpdate(v string) *CreatePolarClawSkillResponseBody
	GetIsUpdate() *string
	SetRequestId(v string) *CreatePolarClawSkillResponseBody
	GetRequestId() *string
	SetSkillName(v string) *CreatePolarClawSkillResponseBody
	GetSkillName() *string
}

type CreatePolarClawSkillResponseBody struct {
	// Used for system scheduled task checks. Modifications are not supported for the current scheduled task.
	//
	// example:
	//
	// true
	IsUpdate *string `json:"IsUpdate,omitempty" xml:"IsUpdate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s CreatePolarClawSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarClawSkillResponseBody) GetIsUpdate() *string {
	return s.IsUpdate
}

func (s *CreatePolarClawSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarClawSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *CreatePolarClawSkillResponseBody) SetIsUpdate(v string) *CreatePolarClawSkillResponseBody {
	s.IsUpdate = &v
	return s
}

func (s *CreatePolarClawSkillResponseBody) SetRequestId(v string) *CreatePolarClawSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarClawSkillResponseBody) SetSkillName(v string) *CreatePolarClawSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *CreatePolarClawSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

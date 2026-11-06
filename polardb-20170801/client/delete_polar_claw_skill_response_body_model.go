// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolarClawSkillResponseBody
	GetRequestId() *string
	SetSkillName(v string) *DeletePolarClawSkillResponseBody
	GetSkillName() *string
}

type DeletePolarClawSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s DeletePolarClawSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarClawSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarClawSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *DeletePolarClawSkillResponseBody) SetRequestId(v string) *DeletePolarClawSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarClawSkillResponseBody) SetSkillName(v string) *DeletePolarClawSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *DeletePolarClawSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

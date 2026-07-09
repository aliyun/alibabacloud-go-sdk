// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluatorSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEvaluatorSkillResponseBody
	GetRequestId() *string
	SetSkillName(v string) *CreateEvaluatorSkillResponseBody
	GetSkillName() *string
}

type CreateEvaluatorSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The name of the created skill.
	//
	// example:
	//
	// trace_context_loader
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
}

func (s CreateEvaluatorSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluatorSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEvaluatorSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEvaluatorSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateEvaluatorSkillResponseBody) SetRequestId(v string) *CreateEvaluatorSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEvaluatorSkillResponseBody) SetSkillName(v string) *CreateEvaluatorSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *CreateEvaluatorSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

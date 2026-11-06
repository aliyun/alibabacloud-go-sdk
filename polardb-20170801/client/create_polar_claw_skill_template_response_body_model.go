// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawSkillTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePolarClawSkillTemplateResponseBody
	GetRequestId() *string
	SetSkillTemplateId(v string) *CreatePolarClawSkillTemplateResponseBody
	GetSkillTemplateId() *string
}

type CreatePolarClawSkillTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill template ID.
	//
	// example:
	//
	// pcpg-skill-**************
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
}

func (s CreatePolarClawSkillTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawSkillTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarClawSkillTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarClawSkillTemplateResponseBody) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *CreatePolarClawSkillTemplateResponseBody) SetRequestId(v string) *CreatePolarClawSkillTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarClawSkillTemplateResponseBody) SetSkillTemplateId(v string) *CreatePolarClawSkillTemplateResponseBody {
	s.SkillTemplateId = &v
	return s
}

func (s *CreatePolarClawSkillTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

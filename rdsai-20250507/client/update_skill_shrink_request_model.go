// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *UpdateSkillShrinkRequest
	GetContentShrink() *string
	SetDbtypesShrink(v string) *UpdateSkillShrinkRequest
	GetDbtypesShrink() *string
	SetDescription(v string) *UpdateSkillShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdateSkillShrinkRequest
	GetName() *string
	SetSkillId(v string) *UpdateSkillShrinkRequest
	GetSkillId() *string
}

type UpdateSkillShrinkRequest struct {
	// The content of the skill.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database engines.
	DbtypesShrink *string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty"`
	// The description of the Skill. It can be up to 1000 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the skill, which can contain only lowercase letters, numbers, and hyphens.
	//
	// example:
	//
	// sql-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8f6a2111-3828-4a9f-a3ce-51ce73c6****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s UpdateSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *UpdateSkillShrinkRequest) GetDbtypesShrink() *string {
	return s.DbtypesShrink
}

func (s *UpdateSkillShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSkillShrinkRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateSkillShrinkRequest) SetContentShrink(v string) *UpdateSkillShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetDbtypesShrink(v string) *UpdateSkillShrinkRequest {
	s.DbtypesShrink = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetDescription(v string) *UpdateSkillShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetName(v string) *UpdateSkillShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetSkillId(v string) *UpdateSkillShrinkRequest {
	s.SkillId = &v
	return s
}

func (s *UpdateSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

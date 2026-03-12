// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *CreateSkillShrinkRequest
	GetContentShrink() *string
	SetDbtypesShrink(v string) *CreateSkillShrinkRequest
	GetDbtypesShrink() *string
	SetDescription(v string) *CreateSkillShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateSkillShrinkRequest
	GetName() *string
}

type CreateSkillShrinkRequest struct {
	// The content of the skill.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database engines.
	//
	// This parameter is required.
	DbtypesShrink *string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty"`
	// The description of the skill. It can be up to 1000 characters in length.
	//
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the skill, which can contain only lowercase letters, numbers, and hyphens.
	//
	// This parameter is required.
	//
	// example:
	//
	// query-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *CreateSkillShrinkRequest) GetDbtypesShrink() *string {
	return s.DbtypesShrink
}

func (s *CreateSkillShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSkillShrinkRequest) SetContentShrink(v string) *CreateSkillShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetDbtypesShrink(v string) *CreateSkillShrinkRequest {
	s.DbtypesShrink = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetDescription(v string) *CreateSkillShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetName(v string) *CreateSkillShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

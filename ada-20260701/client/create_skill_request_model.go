// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSkillRequest
	GetDescription() *string
	SetMetadata(v interface{}) *CreateSkillRequest
	GetMetadata() interface{}
	SetName(v string) *CreateSkillRequest
	GetName() *string
	SetVisibility(v string) *CreateSkillRequest
	GetVisibility() *string
}

type CreateSkillRequest struct {
	// The description of the Skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// A Skill for reviewing code quality, security risks, and coding standards.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Skill metadata in JSON object format. Exactly one content source must be provided. For more information about the fields, see "Request parameter description".
	//
	// This parameter is required.
	//
	// example:
	//
	// {"skillMd":"# Code Review\\nCheck code quality, security risks, and coding standards."}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The unique identifier of the Skill. Only letters, digits, underscores, and hyphens are supported. The value can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The visibility of the Skill. Valid values:
	//
	// - user
	//
	// - tenant
	//
	// Default value: user.
	//
	// example:
	//
	// user
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillRequest) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateSkillRequest) GetName() *string {
	return s.Name
}

func (s *CreateSkillRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateSkillRequest) SetDescription(v string) *CreateSkillRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillRequest) SetMetadata(v interface{}) *CreateSkillRequest {
	s.Metadata = v
	return s
}

func (s *CreateSkillRequest) SetName(v string) *CreateSkillRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillRequest) SetVisibility(v string) *CreateSkillRequest {
	s.Visibility = &v
	return s
}

func (s *CreateSkillRequest) Validate() error {
	return dara.Validate(s)
}

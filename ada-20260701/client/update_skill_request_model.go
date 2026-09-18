// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSkillRequest
	GetDescription() *string
	SetExpectedVersion(v int64) *UpdateSkillRequest
	GetExpectedVersion() *int64
	SetMetadata(v interface{}) *UpdateSkillRequest
	GetMetadata() interface{}
	SetName(v string) *UpdateSkillRequest
	GetName() *string
	SetVisibility(v string) *UpdateSkillRequest
	GetVisibility() *string
}

type UpdateSkillRequest struct {
	// The updated description of the Skill.
	//
	// example:
	//
	// A Skill for performing code reviews, security checks, and risk alerts
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected version number.
	//
	// example:
	//
	// 2
	ExpectedVersion *int64 `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// The updated Skill metadata. The JSON object is replaced as a whole. The content supports exactly one of Transit ID, bundleUrl, or skillMd.
	//
	// example:
	//
	// {"transitId":"transit_example456"}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the Skill to update. This parameter is used only to locate the Skill and cannot be used to modify the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The updated visibility. Valid values: `user` and `tenant`.
	//
	// example:
	//
	// tenant
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateSkillRequest) GetMetadata() interface{} {
	return s.Metadata
}

func (s *UpdateSkillRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSkillRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateSkillRequest) SetDescription(v string) *UpdateSkillRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillRequest) SetExpectedVersion(v int64) *UpdateSkillRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateSkillRequest) SetMetadata(v interface{}) *UpdateSkillRequest {
	s.Metadata = v
	return s
}

func (s *UpdateSkillRequest) SetName(v string) *UpdateSkillRequest {
	s.Name = &v
	return s
}

func (s *UpdateSkillRequest) SetVisibility(v string) *UpdateSkillRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateSkillRequest) Validate() error {
	return dara.Validate(s)
}

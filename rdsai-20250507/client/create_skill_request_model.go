// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *CreateSkillRequest
	GetContent() map[string]interface{}
	SetDbtypes(v []*string) *CreateSkillRequest
	GetDbtypes() []*string
	SetDescription(v string) *CreateSkillRequest
	GetDescription() *string
	SetName(v string) *CreateSkillRequest
	GetName() *string
}

type CreateSkillRequest struct {
	// The content of the skill.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database engines.
	//
	// This parameter is required.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
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

func (s CreateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *CreateSkillRequest) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *CreateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillRequest) GetName() *string {
	return s.Name
}

func (s *CreateSkillRequest) SetContent(v map[string]interface{}) *CreateSkillRequest {
	s.Content = v
	return s
}

func (s *CreateSkillRequest) SetDbtypes(v []*string) *CreateSkillRequest {
	s.Dbtypes = v
	return s
}

func (s *CreateSkillRequest) SetDescription(v string) *CreateSkillRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillRequest) SetName(v string) *CreateSkillRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillRequest) Validate() error {
	return dara.Validate(s)
}

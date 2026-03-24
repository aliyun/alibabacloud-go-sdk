// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDigitalEmployeeSkillRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateDigitalEmployeeSkillRequest
	GetDisplayName() *string
	SetEnable(v bool) *CreateDigitalEmployeeSkillRequest
	GetEnable() *bool
	SetFiles(v []*CreateDigitalEmployeeSkillRequestFiles) *CreateDigitalEmployeeSkillRequest
	GetFiles() []*CreateDigitalEmployeeSkillRequestFiles
	SetRemark(v string) *CreateDigitalEmployeeSkillRequest
	GetRemark() *string
	SetSkillName(v string) *CreateDigitalEmployeeSkillRequest
	GetSkillName() *string
}

type CreateDigitalEmployeeSkillRequest struct {
	// The description of the skill.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the skill.
	//
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether to enable the skill.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of skill files.
	//
	// This parameter is required.
	Files []*CreateDigitalEmployeeSkillRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The remarks.
	//
	// example:
	//
	// remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The name of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// skill
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
}

func (s CreateDigitalEmployeeSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDigitalEmployeeSkillRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateDigitalEmployeeSkillRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateDigitalEmployeeSkillRequest) GetFiles() []*CreateDigitalEmployeeSkillRequestFiles {
	return s.Files
}

func (s *CreateDigitalEmployeeSkillRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateDigitalEmployeeSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateDigitalEmployeeSkillRequest) SetDescription(v string) *CreateDigitalEmployeeSkillRequest {
	s.Description = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequest) SetDisplayName(v string) *CreateDigitalEmployeeSkillRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequest) SetEnable(v bool) *CreateDigitalEmployeeSkillRequest {
	s.Enable = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequest) SetFiles(v []*CreateDigitalEmployeeSkillRequestFiles) *CreateDigitalEmployeeSkillRequest {
	s.Files = v
	return s
}

func (s *CreateDigitalEmployeeSkillRequest) SetRemark(v string) *CreateDigitalEmployeeSkillRequest {
	s.Remark = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequest) SetSkillName(v string) *CreateDigitalEmployeeSkillRequest {
	s.SkillName = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequest) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDigitalEmployeeSkillRequestFiles struct {
	// The content of the file.
	//
	// example:
	//
	// ---
	//
	// name: skill
	//
	// description: description
	//
	// ---
	//
	// # skill
	//
	// skill test
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// SKILL.md
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateDigitalEmployeeSkillRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeSkillRequestFiles) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeSkillRequestFiles) GetContent() *string {
	return s.Content
}

func (s *CreateDigitalEmployeeSkillRequestFiles) GetName() *string {
	return s.Name
}

func (s *CreateDigitalEmployeeSkillRequestFiles) SetContent(v string) *CreateDigitalEmployeeSkillRequestFiles {
	s.Content = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequestFiles) SetName(v string) *CreateDigitalEmployeeSkillRequestFiles {
	s.Name = &v
	return s
}

func (s *CreateDigitalEmployeeSkillRequestFiles) Validate() error {
	return dara.Validate(s)
}

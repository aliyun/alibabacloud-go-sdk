// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDigitalEmployeeSkillRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateDigitalEmployeeSkillRequest
	GetDisplayName() *string
	SetEnable(v bool) *UpdateDigitalEmployeeSkillRequest
	GetEnable() *bool
	SetFiles(v []*UpdateDigitalEmployeeSkillRequestFiles) *UpdateDigitalEmployeeSkillRequest
	GetFiles() []*UpdateDigitalEmployeeSkillRequestFiles
	SetRemark(v string) *UpdateDigitalEmployeeSkillRequest
	GetRemark() *string
}

type UpdateDigitalEmployeeSkillRequest struct {
	// Description
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Display name
	//
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Is enabled
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// List of skill files
	//
	// This parameter is required.
	Files []*UpdateDigitalEmployeeSkillRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// Remark
	//
	// example:
	//
	// remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateDigitalEmployeeSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDigitalEmployeeSkillRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateDigitalEmployeeSkillRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateDigitalEmployeeSkillRequest) GetFiles() []*UpdateDigitalEmployeeSkillRequestFiles {
	return s.Files
}

func (s *UpdateDigitalEmployeeSkillRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDigitalEmployeeSkillRequest) SetDescription(v string) *UpdateDigitalEmployeeSkillRequest {
	s.Description = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequest) SetDisplayName(v string) *UpdateDigitalEmployeeSkillRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequest) SetEnable(v bool) *UpdateDigitalEmployeeSkillRequest {
	s.Enable = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequest) SetFiles(v []*UpdateDigitalEmployeeSkillRequestFiles) *UpdateDigitalEmployeeSkillRequest {
	s.Files = v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequest) SetRemark(v string) *UpdateDigitalEmployeeSkillRequest {
	s.Remark = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequest) Validate() error {
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

type UpdateDigitalEmployeeSkillRequestFiles struct {
	// File content
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
	// File name
	//
	// example:
	//
	// SKILL.md
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateDigitalEmployeeSkillRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeSkillRequestFiles) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeSkillRequestFiles) GetContent() *string {
	return s.Content
}

func (s *UpdateDigitalEmployeeSkillRequestFiles) GetName() *string {
	return s.Name
}

func (s *UpdateDigitalEmployeeSkillRequestFiles) SetContent(v string) *UpdateDigitalEmployeeSkillRequestFiles {
	s.Content = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequestFiles) SetName(v string) *UpdateDigitalEmployeeSkillRequestFiles {
	s.Name = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillRequestFiles) Validate() error {
	return dara.Validate(s)
}

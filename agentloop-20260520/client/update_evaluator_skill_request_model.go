// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluatorSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *UpdateEvaluatorSkillRequest
	GetAgentSpace() *string
	SetDescription(v string) *UpdateEvaluatorSkillRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateEvaluatorSkillRequest
	GetDisplayName() *string
	SetEnable(v bool) *UpdateEvaluatorSkillRequest
	GetEnable() *bool
	SetFiles(v []*UpdateEvaluatorSkillRequestFiles) *UpdateEvaluatorSkillRequest
	GetFiles() []*UpdateEvaluatorSkillRequestFiles
	SetClientToken(v string) *UpdateEvaluatorSkillRequest
	GetClientToken() *string
}

type UpdateEvaluatorSkillRequest struct {
	// The name of the AgentSpace.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The description of the skill.
	//
	// example:
	//
	// Read trace context to assist evaluation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the skill.
	//
	// example:
	//
	// Read Trace Context
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether the skill is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of skill files. If specified, the content of the skill files is updated.
	//
	// example:
	//
	// [{"name":"SKILL.md","content":"# Trace Context Loader","remark":"Primary skill description"}]
	Files []*UpdateEvaluatorSkillRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The idempotency token. CloudSpec declares this query parameter, but the backend does not currently perform idempotency checks.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateEvaluatorSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorSkillRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *UpdateEvaluatorSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEvaluatorSkillRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateEvaluatorSkillRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateEvaluatorSkillRequest) GetFiles() []*UpdateEvaluatorSkillRequestFiles {
	return s.Files
}

func (s *UpdateEvaluatorSkillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEvaluatorSkillRequest) SetAgentSpace(v string) *UpdateEvaluatorSkillRequest {
	s.AgentSpace = &v
	return s
}

func (s *UpdateEvaluatorSkillRequest) SetDescription(v string) *UpdateEvaluatorSkillRequest {
	s.Description = &v
	return s
}

func (s *UpdateEvaluatorSkillRequest) SetDisplayName(v string) *UpdateEvaluatorSkillRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateEvaluatorSkillRequest) SetEnable(v bool) *UpdateEvaluatorSkillRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEvaluatorSkillRequest) SetFiles(v []*UpdateEvaluatorSkillRequestFiles) *UpdateEvaluatorSkillRequest {
	s.Files = v
	return s
}

func (s *UpdateEvaluatorSkillRequest) SetClientToken(v string) *UpdateEvaluatorSkillRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEvaluatorSkillRequest) Validate() error {
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

type UpdateEvaluatorSkillRequestFiles struct {
	// The content of the skill file.
	//
	// This parameter is required.
	//
	// example:
	//
	// # Trace Context Loader
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the skill file.
	//
	// This parameter is required.
	//
	// example:
	//
	// SKILL.md
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The remark for the file.
	//
	// example:
	//
	// Primary skill description
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateEvaluatorSkillRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorSkillRequestFiles) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorSkillRequestFiles) GetContent() *string {
	return s.Content
}

func (s *UpdateEvaluatorSkillRequestFiles) GetName() *string {
	return s.Name
}

func (s *UpdateEvaluatorSkillRequestFiles) GetRemark() *string {
	return s.Remark
}

func (s *UpdateEvaluatorSkillRequestFiles) SetContent(v string) *UpdateEvaluatorSkillRequestFiles {
	s.Content = &v
	return s
}

func (s *UpdateEvaluatorSkillRequestFiles) SetName(v string) *UpdateEvaluatorSkillRequestFiles {
	s.Name = &v
	return s
}

func (s *UpdateEvaluatorSkillRequestFiles) SetRemark(v string) *UpdateEvaluatorSkillRequestFiles {
	s.Remark = &v
	return s
}

func (s *UpdateEvaluatorSkillRequestFiles) Validate() error {
	return dara.Validate(s)
}

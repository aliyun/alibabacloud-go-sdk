// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *UpdateSkillRequest
	GetContent() map[string]interface{}
	SetDbtypes(v []*string) *UpdateSkillRequest
	GetDbtypes() []*string
	SetDescription(v string) *UpdateSkillRequest
	GetDescription() *string
	SetName(v string) *UpdateSkillRequest
	GetName() *string
	SetSkillId(v string) *UpdateSkillRequest
	GetSkillId() *string
	SetUpdateNote(v string) *UpdateSkillRequest
	GetUpdateNote() *string
	SetUploadId(v string) *UpdateSkillRequest
	GetUploadId() *string
	SetUploadToken(v string) *UpdateSkillRequest
	GetUploadToken() *string
	SetWorkspaceId(v string) *UpdateSkillRequest
	GetWorkspaceId() *string
}

type UpdateSkillRequest struct {
	// The content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database types.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The description of the Skill. The description can be up to 1000 characters in length.
	//
	// example:
	//
	// SQL Review Expert: Comprehensively reviews SQL for security, performance, and compliance, identifies risks, and provides optimization suggestions. Activated immediately when a user submits SQL or asks "SQL review", "SQL Review", "any risks?", or "how to optimize"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the Skill. The name can contain only lowercase letters, digits, and hyphens.
	//
	// example:
	//
	// sql-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the Skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8f6a2111-3828-4a9f-a3ce-51ce73c6****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The version update note for the ContextDB Skill. Specify this parameter when confirming storage after the check passes. The note can be up to 1000 characters in length.
	//
	// example:
	//
	// Update Skill documentation
	UpdateNote *string `json:"UpdateNote,omitempty" xml:"UpdateNote,omitempty"`
	// The upload session ID returned by GetUploadPolicy for ContextDB full package updates. This parameter is not required for original RDS field updates.
	//
	// example:
	//
	// upload-example
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// The upload session token paired with UploadId, used to verify ownership. This parameter is not required for original RDS field updates.
	//
	// example:
	//
	// token-example
	UploadToken *string `json:"UploadToken,omitempty" xml:"UploadToken,omitempty"`
	// The ContextDB workspace ID. This parameter is not required for regular Skill calls.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateSkillRequest) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *UpdateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateSkillRequest) GetUpdateNote() *string {
	return s.UpdateNote
}

func (s *UpdateSkillRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UpdateSkillRequest) GetUploadToken() *string {
	return s.UploadToken
}

func (s *UpdateSkillRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateSkillRequest) SetContent(v map[string]interface{}) *UpdateSkillRequest {
	s.Content = v
	return s
}

func (s *UpdateSkillRequest) SetDbtypes(v []*string) *UpdateSkillRequest {
	s.Dbtypes = v
	return s
}

func (s *UpdateSkillRequest) SetDescription(v string) *UpdateSkillRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillRequest) SetName(v string) *UpdateSkillRequest {
	s.Name = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillId(v string) *UpdateSkillRequest {
	s.SkillId = &v
	return s
}

func (s *UpdateSkillRequest) SetUpdateNote(v string) *UpdateSkillRequest {
	s.UpdateNote = &v
	return s
}

func (s *UpdateSkillRequest) SetUploadId(v string) *UpdateSkillRequest {
	s.UploadId = &v
	return s
}

func (s *UpdateSkillRequest) SetUploadToken(v string) *UpdateSkillRequest {
	s.UploadToken = &v
	return s
}

func (s *UpdateSkillRequest) SetWorkspaceId(v string) *UpdateSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateSkillRequest) Validate() error {
	return dara.Validate(s)
}

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
	SetUpdateNote(v string) *UpdateSkillShrinkRequest
	GetUpdateNote() *string
	SetUploadId(v string) *UpdateSkillShrinkRequest
	GetUploadId() *string
	SetUploadToken(v string) *UpdateSkillShrinkRequest
	GetUploadToken() *string
	SetWorkspaceId(v string) *UpdateSkillShrinkRequest
	GetWorkspaceId() *string
}

type UpdateSkillShrinkRequest struct {
	// The content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database types.
	DbtypesShrink *string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty"`
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

func (s *UpdateSkillShrinkRequest) GetUpdateNote() *string {
	return s.UpdateNote
}

func (s *UpdateSkillShrinkRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UpdateSkillShrinkRequest) GetUploadToken() *string {
	return s.UploadToken
}

func (s *UpdateSkillShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *UpdateSkillShrinkRequest) SetUpdateNote(v string) *UpdateSkillShrinkRequest {
	s.UpdateNote = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetUploadId(v string) *UpdateSkillShrinkRequest {
	s.UploadId = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetUploadToken(v string) *UpdateSkillShrinkRequest {
	s.UploadToken = &v
	return s
}

func (s *UpdateSkillShrinkRequest) SetWorkspaceId(v string) *UpdateSkillShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

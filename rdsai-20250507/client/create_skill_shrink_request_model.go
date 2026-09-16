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
	SetUploadId(v string) *CreateSkillShrinkRequest
	GetUploadId() *string
	SetUploadToken(v string) *CreateSkillShrinkRequest
	GetUploadToken() *string
	SetWorkspaceId(v string) *CreateSkillShrinkRequest
	GetWorkspaceId() *string
}

type CreateSkillShrinkRequest struct {
	// The content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database types.
	DbtypesShrink *string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty"`
	// The Skill description. The description can be up to 1000 characters in length.
	//
	// example:
	//
	// SQL query optimization skill
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Skill name. The name can contain only lowercase letters, digits, and hyphens.
	//
	// example:
	//
	// query-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Skill upload session ID.
	//
	// example:
	//
	// upload-example
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// The Skill upload session token.
	//
	// example:
	//
	// token-example
	UploadToken *string `json:"UploadToken,omitempty" xml:"UploadToken,omitempty"`
	// The ContextDB workspace ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateSkillShrinkRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateSkillShrinkRequest) GetUploadToken() *string {
	return s.UploadToken
}

func (s *CreateSkillShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *CreateSkillShrinkRequest) SetUploadId(v string) *CreateSkillShrinkRequest {
	s.UploadId = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetUploadToken(v string) *CreateSkillShrinkRequest {
	s.UploadToken = &v
	return s
}

func (s *CreateSkillShrinkRequest) SetWorkspaceId(v string) *CreateSkillShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}

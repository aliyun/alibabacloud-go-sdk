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
	SetUploadId(v string) *CreateSkillRequest
	GetUploadId() *string
	SetUploadToken(v string) *CreateSkillRequest
	GetUploadToken() *string
	SetWorkspaceId(v string) *CreateSkillRequest
	GetWorkspaceId() *string
}

type CreateSkillRequest struct {
	// The content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database types.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
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

func (s *CreateSkillRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateSkillRequest) GetUploadToken() *string {
	return s.UploadToken
}

func (s *CreateSkillRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *CreateSkillRequest) SetUploadId(v string) *CreateSkillRequest {
	s.UploadId = &v
	return s
}

func (s *CreateSkillRequest) SetUploadToken(v string) *CreateSkillRequest {
	s.UploadToken = &v
	return s
}

func (s *CreateSkillRequest) SetWorkspaceId(v string) *CreateSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateSkillRequest) Validate() error {
	return dara.Validate(s)
}

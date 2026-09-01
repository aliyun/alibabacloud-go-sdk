// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSkillMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDataAgentSkillMetaRequest
	GetDescription() *string
	SetSkillName(v string) *CreateDataAgentSkillMetaRequest
	GetSkillName() *string
	SetUploadLocation(v string) *CreateDataAgentSkillMetaRequest
	GetUploadLocation() *string
	SetWorkspaceId(v string) *CreateDataAgentSkillMetaRequest
	GetWorkspaceId() *string
}

type CreateDataAgentSkillMetaRequest struct {
	// The skill description.
	//
	// - By default, this parameter is optional. The backend parses the ZIP package specified by UploadLocation to obtain the skill description.
	//
	// example:
	//
	// This is a demo skill description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The skill name.
	//
	// - By default, this parameter is optional. The backend parses the ZIP package specified by UploadLocation to obtain the skill name.
	//
	// example:
	//
	// data-query-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The full path for uploading the skill ZIP file.
	//
	// - Format: The UploadDir field returned by the DescribeSkillFileUploadSignature operation concatenated with the file name.
	//
	// - Example: ${UploadDir}/${Filename}
	UploadLocation *string `json:"UploadLocation,omitempty" xml:"UploadLocation,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// tmbbtfv8***********zuqko6
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentSkillMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSkillMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSkillMetaRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataAgentSkillMetaRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateDataAgentSkillMetaRequest) GetUploadLocation() *string {
	return s.UploadLocation
}

func (s *CreateDataAgentSkillMetaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentSkillMetaRequest) SetDescription(v string) *CreateDataAgentSkillMetaRequest {
	s.Description = &v
	return s
}

func (s *CreateDataAgentSkillMetaRequest) SetSkillName(v string) *CreateDataAgentSkillMetaRequest {
	s.SkillName = &v
	return s
}

func (s *CreateDataAgentSkillMetaRequest) SetUploadLocation(v string) *CreateDataAgentSkillMetaRequest {
	s.UploadLocation = &v
	return s
}

func (s *CreateDataAgentSkillMetaRequest) SetWorkspaceId(v string) *CreateDataAgentSkillMetaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentSkillMetaRequest) Validate() error {
	return dara.Validate(s)
}

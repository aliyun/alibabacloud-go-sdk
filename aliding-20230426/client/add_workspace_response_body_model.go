// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspace(v *AddWorkspaceResponseBodyWorkspace) *AddWorkspaceResponseBody
	GetWorkspace() *AddWorkspaceResponseBodyWorkspace
}

type AddWorkspaceResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Workspace *AddWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s AddWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkspaceResponseBody) GetWorkspace() *AddWorkspaceResponseBodyWorkspace {
	return s.Workspace
}

func (s *AddWorkspaceResponseBody) SetRequestId(v string) *AddWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkspaceResponseBody) SetWorkspace(v *AddWorkspaceResponseBodyWorkspace) *AddWorkspaceResponseBody {
	s.Workspace = v
	return s
}

func (s *AddWorkspaceResponseBody) Validate() error {
	if s.Workspace != nil {
		if err := s.Workspace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWorkspaceResponseBodyWorkspace struct {
	// example:
	//
	// ding16b241fd05********288
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/XXXX.jpg
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// example:
	//
	// 2023-05-15T11:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 01472825524039877041
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// workspace_description
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *AddWorkspaceResponseBodyWorkspaceIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	// example:
	//
	// 2023-05-15T11:29Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 01472825524039877041
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// workspace_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// READER
	PermissionRole *string `json:"PermissionRole,omitempty" xml:"PermissionRole,omitempty"`
	// example:
	//
	// root_node_uuid
	RootNodeId *string `json:"RootNodeId,omitempty" xml:"RootNodeId,omitempty"`
	// example:
	//
	// lHiicjNFM2iSFYSdz2iPuI8ZwiEiE
	TeamId *string `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
	// example:
	//
	// TEAM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// workspace_url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// workspace_id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceResponseBodyWorkspace) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponseBodyWorkspace) GetCorpId() *string {
	return s.CorpId
}

func (s *AddWorkspaceResponseBodyWorkspace) GetCover() *string {
	return s.Cover
}

func (s *AddWorkspaceResponseBodyWorkspace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddWorkspaceResponseBodyWorkspace) GetCreatorId() *string {
	return s.CreatorId
}

func (s *AddWorkspaceResponseBodyWorkspace) GetDescription() *string {
	return s.Description
}

func (s *AddWorkspaceResponseBodyWorkspace) GetIcon() *AddWorkspaceResponseBodyWorkspaceIcon {
	return s.Icon
}

func (s *AddWorkspaceResponseBodyWorkspace) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *AddWorkspaceResponseBodyWorkspace) GetModifierId() *string {
	return s.ModifierId
}

func (s *AddWorkspaceResponseBodyWorkspace) GetName() *string {
	return s.Name
}

func (s *AddWorkspaceResponseBodyWorkspace) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *AddWorkspaceResponseBodyWorkspace) GetRootNodeId() *string {
	return s.RootNodeId
}

func (s *AddWorkspaceResponseBodyWorkspace) GetTeamId() *string {
	return s.TeamId
}

func (s *AddWorkspaceResponseBodyWorkspace) GetType() *string {
	return s.Type
}

func (s *AddWorkspaceResponseBodyWorkspace) GetUrl() *string {
	return s.Url
}

func (s *AddWorkspaceResponseBodyWorkspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCorpId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCover(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *AddWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCreatorId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetDescription(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetIcon(v *AddWorkspaceResponseBodyWorkspaceIcon) *AddWorkspaceResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetModifiedTime(v string) *AddWorkspaceResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetModifierId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetName(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetPermissionRole(v string) *AddWorkspaceResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetRootNodeId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetTeamId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetType(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetUrl(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWorkspaceResponseBodyWorkspaceIcon struct {
	// example:
	//
	// TEAM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/avatar3.jpg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddWorkspaceResponseBodyWorkspaceIcon) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) GetType() *string {
	return s.Type
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) GetValue() *string {
	return s.Value
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) SetType(v string) *AddWorkspaceResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) SetValue(v string) *AddWorkspaceResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) Validate() error {
	return dara.Validate(s)
}

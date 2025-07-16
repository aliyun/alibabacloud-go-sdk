// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkspacesResponseBody
	GetRequestId() *string
	SetWorkspace(v []*GetWorkspacesResponseBodyWorkspace) *GetWorkspacesResponseBody
	GetWorkspace() []*GetWorkspacesResponseBodyWorkspace
}

type GetWorkspacesResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Workspace []*GetWorkspacesResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Repeated"`
}

func (s GetWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspacesResponseBody) GetWorkspace() []*GetWorkspacesResponseBodyWorkspace {
	return s.Workspace
}

func (s *GetWorkspacesResponseBody) SetRequestId(v string) *GetWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspacesResponseBody) SetWorkspace(v []*GetWorkspacesResponseBodyWorkspace) *GetWorkspacesResponseBody {
	s.Workspace = v
	return s
}

func (s *GetWorkspacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWorkspacesResponseBodyWorkspace struct {
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
	Description *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *GetWorkspacesResponseBodyWorkspaceIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
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
	// ZgpG2NdyVXXjrKKzIBqMp5zkVMwvDqPk
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
	// By8jQS1ZYjGn5b0M
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspacesResponseBodyWorkspace) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponseBodyWorkspace) GetCorpId() *string {
	return s.CorpId
}

func (s *GetWorkspacesResponseBodyWorkspace) GetCover() *string {
	return s.Cover
}

func (s *GetWorkspacesResponseBodyWorkspace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspacesResponseBodyWorkspace) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetWorkspacesResponseBodyWorkspace) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspacesResponseBodyWorkspace) GetIcon() *GetWorkspacesResponseBodyWorkspaceIcon {
	return s.Icon
}

func (s *GetWorkspacesResponseBodyWorkspace) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetWorkspacesResponseBodyWorkspace) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetWorkspacesResponseBodyWorkspace) GetName() *string {
	return s.Name
}

func (s *GetWorkspacesResponseBodyWorkspace) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *GetWorkspacesResponseBodyWorkspace) GetRootNodeId() *string {
	return s.RootNodeId
}

func (s *GetWorkspacesResponseBodyWorkspace) GetTeamId() *string {
	return s.TeamId
}

func (s *GetWorkspacesResponseBodyWorkspace) GetType() *string {
	return s.Type
}

func (s *GetWorkspacesResponseBodyWorkspace) GetUrl() *string {
	return s.Url
}

func (s *GetWorkspacesResponseBodyWorkspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspacesResponseBodyWorkspace) SetCorpId(v string) *GetWorkspacesResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetCover(v string) *GetWorkspacesResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetCreateTime(v string) *GetWorkspacesResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetCreatorId(v string) *GetWorkspacesResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetDescription(v string) *GetWorkspacesResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetIcon(v *GetWorkspacesResponseBodyWorkspaceIcon) *GetWorkspacesResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetModifiedTime(v string) *GetWorkspacesResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetModifierId(v string) *GetWorkspacesResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetName(v string) *GetWorkspacesResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetPermissionRole(v string) *GetWorkspacesResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetRootNodeId(v string) *GetWorkspacesResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetTeamId(v string) *GetWorkspacesResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetType(v string) *GetWorkspacesResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetUrl(v string) *GetWorkspacesResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) SetWorkspaceId(v string) *GetWorkspacesResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspace) Validate() error {
	return dara.Validate(s)
}

type GetWorkspacesResponseBodyWorkspaceIcon struct {
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/avatar3.jpg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetWorkspacesResponseBodyWorkspaceIcon) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponseBodyWorkspaceIcon) GetType() *string {
	return s.Type
}

func (s *GetWorkspacesResponseBodyWorkspaceIcon) GetValue() *string {
	return s.Value
}

func (s *GetWorkspacesResponseBodyWorkspaceIcon) SetType(v string) *GetWorkspacesResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaceIcon) SetValue(v string) *GetWorkspacesResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaceIcon) Validate() error {
	return dara.Validate(s)
}
